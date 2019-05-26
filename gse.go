// GSE - Go Show Env - see https://github.com/DBuret/gse
package main

import (
	"bytes"
	"strings"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"io"
	"io/ioutil"
	"github.com/namsral/flag"
)

var (
    Trace   *log.Logger
    Info    *log.Logger
    Warning *log.Logger
    Error   *log.Logger
)

func Init(
    traceHandle io.Writer,
    infoHandle io.Writer,
    warningHandle io.Writer,
    errorHandle io.Writer) {

    Trace = log.New(traceHandle,
        "TRACE: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    Info = log.New(infoHandle,
        "INFO: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    Warning = log.New(warningHandle,
        "WARNING: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    Error = log.New(errorHandle,
        "ERROR: ",
        log.Ldate|log.Ltime|log.Lshortfile)
}

func check(err error) {
	if err != nil {
		Error.Printf(err)
	}
}

type output struct {
	Method       string
	Host         string
	Url          string
	Proto        string
	HeaderOutput []string
	BodyOutput   string
	EnvOutput    []string
}

func showEnvHandler(w http.ResponseWriter, r *http.Request) {
	o := output{
		Method:       r.Method,
		Host:         r.Host,
		Proto:        r.Proto,
		Url:          fmt.Sprint(r.URL),
		HeaderOutput: []string{},
		BodyOutput:   "",
		EnvOutput:    []string{}
	}

	Trace.Printf(",", o.Method, " ,", o.Host, " ,", o.Url)

	//headers
	header := r.Header
	sortedHeaders := make([]string, 0, len(header))
	for k := range header {
		sortedHeaders = append(sortedHeaders, k)
	}
	sort.Strings(sortedHeaders)
	for k := range sortedHeaders {
		o.HeaderOutput = append(o.HeaderOutput,
			fmt.Sprintf("%s=%s", sortedHeaders[k], strings.Join(header[sortedHeaders[k]],", ")))
	}
	
	//body
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	o.BodyOutput = buf.String() // complete copy of the bytes in the buffer.

	// ENV
	for _, e := range os.Environ() {
		o.EnvOutput = append(o.EnvOutput, e)
	}
	sort.Strings(o.EnvOutput)

	t, err := template.ParseFiles("template.html")
	check(err)
	t.Execute(w, o)
}

func main() {
	var programVersion = "3.0"
	var programName = "gse"
	
	Init(ioutil.Stdout, os.Stdout, os.Stdout, os.Stderr)
    
    // env parsing
    flag.StringVar(&uri,"basepath", "/gse", "base path in the url")
    flag.IntVar(&port,"port", 28657, "default listening port")
    flag.BoolVar(&healthcheck,"healthcheck", true, "enable/disable healthckeck endpoint")

    flag.Parse()

	// get an http server
	mux := http.NewServeMux()
	
	// handlers
	//	 /uri
	mux.HandleFunc(uri, showEnvHandler)
	
	//	/uri/version
	mux.HandleFunc(uri + "/version", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprintln(w, programVersion)
	})
	
	//	/uri/health
	mux.HandleFunc(uri + "/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method = "GET" {
			if healthcheck {
				w.WriteHeader(200)
				fmt.Fprintln(w, "I'm alive")
			} else {
				w.WriteHeader(503)
				fmt.Fprintln(w, "I´m sick")
			}
		} else if r.Method = "POST" {
			healthcheck = !healthcheck
		}
	})
	
	//	/
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintln(w, "Oops, you requested an unknown location.\n FYI, my base path is " + uri)
	})
	
	
	// start http server
	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}

	// log config used
	Info.Printf("Starting %s (%s) on port %s with basepath %s and healthtcheck=%s...\n",
		programName,
		programVersion, 
		port, 
		uri, 
		healthcheck )
	log.Fatal(s.ListenAndServe())
}