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
	"regexp"
	"sort"
	"strings"

	"github.com/namsral/flag"
)

var (
	programVersion = "4.1"
	programName    = "gse"

	// Trace logger: debug
	Trace *log.Logger
	// Info logger: standard info logs
	Info *log.Logger
	// Warning logger: non fatal errors
	Warning *log.Logger
	// Error logger: panic
	Error *log.Logger

	uri, mark             string
	port                  int
	healthcheck, loggerEP bool
)

// Init : setup loggers
func Init(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func check(err error) {
	if err != nil {
		Error.Printf("%s", err)
	}
}

func sanitize(i string) string {
	re := regexp.MustCompile(`[^0-9A-Za-z\{\}\:\,\[\]]`)
	return fmt.Sprintf("%q\n", re.ReplaceAllString(i, ""))
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
	o.BodyOutput = buf.String()

	// ENV
	for _, e := range os.Environ() {
		o.EnvOutput = append(o.EnvOutput, e)
	}
	sort.Strings(o.EnvOutput)

	t, err := template.ParseFiles("template.html")
	check(err)
	t.Execute(w, o)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	Trace.Printf("%s %s%s", r.Method, r.Host, r.URL)
	w.WriteHeader(200)
	fmt.Fprintln(w, programVersion+mark)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(200)
		fmt.Fprintln(w, fmt.Sprintf("healthcheck has been switched to %t", healthcheck))
		Trace.Printf("healthcheck has been switched to %t", healthcheck)
	}
}

func loggerHandler(w http.ResponseWriter, r *http.Request) {
	if loggerEP {
		switch r.Method {
		case "GET":
			w.WriteHeader(400)
			fmt.Fprintln(w, "logger endpoint uncorrectly called with GET method")
			Warning.Printf("logger endpoint uncorrectly called with GET method")
		case "POST", "PUT":
			bdy := new(bytes.Buffer)
			bdy.ReadFrom(r.Body)

			w.WriteHeader(200)
			fmt.Fprintln(w, "data ingested.")
			Trace.Printf(fmt.Sprintf("%s %s%s, LOGGER: %s",
				r.Method,
				r.Host,
				r.URL,
				sanitize(bdy.String())))
		}
	} else {
		w.WriteHeader(400)
		fmt.Fprintln(w, "logger endpoint called but not activited in configuration")
		Warning.Printf("logger endpoint called but not activited in configuration")
	}
}

func main() {

	Init(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

	// env parsing
	confManager := flag.NewFlagSetWithEnvPrefix(os.Args[0], "GSE", 0)

	confManager.StringVar(&uri, "basepath", "/gse", "base path in the url")
	confManager.IntVar(&port, "port", 28657, "default listening port")
	confManager.BoolVar(&healthcheck, "healthcheck", true, "enable/disable healthckeck endpoint")
	confManager.BoolVar(&loggerEP, "logger", false, "enable/disable logger endpoint")
	confManager.StringVar(&mark, "stamp", "", "specify a stamp to be added to version endpoint answer")

	confManager.Parse(os.Args[1:])

	// get an http server
	mux := http.NewServeMux()

	// handlers
	//	 /uri
	mux.HandleFunc(uri, showEnvHandler)
	mux.HandleFunc(uri+"/", showEnvHandler)

	//	/uri/version
	mux.HandleFunc(uri+"/version", versionHandler)
	mux.HandleFunc(uri+"/version/", versionHandler)

	//	/uri/health
	mux.HandleFunc(uri+"/health", healthHandler)
	mux.HandleFunc(uri+"/health/", healthHandler)

	//	/uri/logger
	mux.HandleFunc(uri+"/logger", loggerHandler)
	mux.HandleFunc(uri+"/logger/", loggerHandler)

	//	default handler for /
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
	Info.Printf("Starting %s (%s) on port %d with basepath %s, healthtcheck=%t, and logger_endpoint=%t...\n",
		programName,
		programVersion+mark,
		port,
		uri,
		healthcheck,
		loggerEP)

	log.Fatal(s.ListenAndServe())
}
