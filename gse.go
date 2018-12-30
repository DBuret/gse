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
	"github.com/DBuret/pathandport"
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
		EnvOutput:    []string{}}

	Info.Printf(",", o.Method, " ,", o.Host, " ,", o.Url)

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
	var programVersion = "0.6"
	var programName = "gse"
	
	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	uri, port, info, err := pathAndPort.Parse(programName, "28657")
	
	check(err)

	mux := http.NewServeMux()
	
	mux.HandleFunc(uri, showEnvHandler)
	
	mux.HandleFunc(uri + "/version", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprintln(w, "Oops, you requested an unknown location.\n FYI, my base path is " + uri)
	})
	
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintln(w, "Oops, you requested an unknown location.\n FYI, my base path is " + uri)
	})
	
	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}

	Info.Printf("%s\nStarting %s (%s) on port %s with basepath %s ...\n", info,programName, programVersion, port, uri)
	log.Fatal(s.ListenAndServe())
}


