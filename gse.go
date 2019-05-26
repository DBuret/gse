// GSE - Go Show Env - see https://github.com/DBuret/gse
package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"

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
		Error.Printf("%s", err)
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
		EnvOutput:    []string{}}

	Trace.Printf("%s %s%s", o.Method, o.Host, o.Url)

	//headers
	header := r.Header
	sortedHeaders := make([]string, 0, len(header))
	for k := range header {
		sortedHeaders = append(sortedHeaders, k)
	}
	sort.Strings(sortedHeaders)
	for k := range sortedHeaders {
		o.HeaderOutput = append(o.HeaderOutput,
			fmt.Sprintf("%s=%s", sortedHeaders[k], strings.Join(header[sortedHeaders[k]], ", ")))
	}

	//body
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	o.BodyOutput = buf.String() // copy buffer.

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
	var uri, mark string
	var port int
	var healthcheck bool

	Init(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

	// env parsing
	confManager := flag.NewFlagSetWithEnvPrefix(os.Args[0], "GSE", 0)

	confManager.StringVar(&uri, "basepath", "/gse", "base path in the url")
	confManager.IntVar(&port, "port", 28657, "default listening port")
	confManager.BoolVar(&healthcheck, "healthcheck", true, "enable/disable healthckeck endpoint")
	confManager.StringVar(&mark, "stamp", "", "specify a stamp to be added to version endpoint answer")

	confManager.Parse(os.Args[1:])

	// get an http server
	mux := http.NewServeMux()

	// handlers
	//	 /uri
	mux.HandleFunc(uri, showEnvHandler)

	//	/uri/version
	mux.HandleFunc(uri+"/version", func(w http.ResponseWriter, r *http.Request) {
		Trace.Printf("%s %s%s", r.Method, r.Host, r.URL)
		w.WriteHeader(200)
		fmt.Fprintln(w, programVersion+mark)
	})

	//	/uri/health
	mux.HandleFunc(uri+"/health", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// health status
		case "GET":
			if healthcheck {
				w.WriteHeader(200)
				fmt.Fprintln(w, "I'm alive")
			} else {
				w.WriteHeader(503)
				fmt.Fprintln(w, "I´m sick")
			}
		// flip/flop health state
		case "POST", "PUT":
			// this is dirty - no mutex...
			healthcheck = !healthcheck
			Trace.Printf("healthcheck has been switched to %t", healthcheck)
		}
	})

	//	/
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintln(w, "Oops, you requested an unknown location.\n FYI, my base path is "+uri)
	})

	// start http server
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	// log config used
	Info.Printf("Starting %s (%s) on port %d with basepath %s and healthtcheck=%t...\n",
		programName,
		programVersion+mark,
		port,
		uri,
		healthcheck)

	log.Fatal(s.ListenAndServe())
}
