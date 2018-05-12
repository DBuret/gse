// GSE - Go Show Env - see https://github.com/DBuret/gse
package main

import (
	"bytes"
	"fmt"
	//"html"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	//	"strings"

	"github.com/DBuret/pathandport"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
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

func handler(w http.ResponseWriter, r *http.Request) {
	o := output{
		Method:       r.Method,
		Host:         r.Host,
		Proto:        r.Proto,
		Url:          fmt.Sprint(r.URL),
		HeaderOutput: []string{},
		BodyOutput:   "",
		EnvOutput:    []string{}}

	log.Print(",", o.Method, " ,", o.Host, " ,", o.Url)

	//headers
	header := r.Header
	sortedHeaders := make([]string, 0, len(header))
	for k := range header {
		sortedHeaders = append(sortedHeaders, k)
	}
	sort.Strings(sortedHeaders)
	for k := range sortedHeaders {
		//fmt.Fprintf(w, " %s = %s\n", k, r.Header[k])
		o.HeaderOutput = append(o.HeaderOutput,
			fmt.Sprintf("%s=%s", sortedHeaders[k], header[sortedHeaders[k]]))
	}

	//var b strings.Builder

	//body
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	o.BodyOutput = buf.String() // Does a complete copy of the bytes in the buffer.

	// ENV
	for _, e := range os.Environ() {
		//pair := strings.Split(e, "=")
		o.EnvOutput = append(o.EnvOutput, e)
	}
	sort.Strings(o.EnvOutput)

	t, err := template.ParseFiles("template.html")
	check(err)
	t.Execute(w, o)
}

func main() {
	var programVersion = "0.4"
	var programName = "gse"

	var uri string
	var port string

	uri, port = pathandport.Parse(programName, "28657")

	mux := http.NewServeMux()
	mux.HandleFunc(uri, handler)
	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}

	log.Printf("Starting %s (%s) on port %s with basepath %s ...\n", programName, programVersion, port, uri)
	log.Fatal(s.ListenAndServe())
}

/*type Request struct {
        // Method specifies the HTTP method (GET, POST, PUT, etc.).
        Method string

        // URL specifies either the URI being requested (for server
        // requests)
        //
        // the URL is parsed from the URI
        // supplied on the Request-Line as stored in RequestURI.  For
        // most requests, fields other than Path and RawQuery will be
        // empty. (See RFC 2616, Section 5.1.2)
        URL *url.URL

        // The protocol version for incoming server requests.
        //
        // See the docs on Transport for details.
        Proto      string // "HTTP/1.0"
        ProtoMajor int    // 1
        ProtoMinor int    // 0

        // Header contains the request header fields either received
        // by the server or to be sent by the client.
        //
        // If a server received a request with header lines,
        //
        //	Host: example.com
        //	accept-encoding: gzip, deflate
        //	Accept-Language: en-us
        //	fOO: Bar
        //	foo: two
        //
        // then
        //
        //	Header = map[string][]string{
        //		"Accept-Encoding": {"gzip, deflate"},
        //		"Accept-Language": {"en-us"},
        //		"Foo": {"Bar", "two"},
        //	}
        //
        // For incoming requests, the Host header is promoted to the
        // Request.Host field and removed from the Header map.
        //
        // HTTP defines that header names are case-insensitive. The
        // request parser implements this by using CanonicalHeaderKey,
        // making the first character and any characters following a
        // hyphen uppercase and the rest lowercase.
        //
        Header Header

        // Body is the request's body.
        // For server requests the Request Body is always non-nil
        // but will return EOF immediately when no body is present.
        // The Server will close the request body. The ServeHTTP
        // Handler does not need to.
        Body io.ReadCloser




        // ContentLength records the length of the associated content.
        // The value -1 indicates that the length is unknown.
        // Values >= 0 indicate that the given number of bytes may
        // be read from Body.
        // For client requests, a value of 0 with a non-nil Body is
        // also treated as unknown.
        ContentLength int64

        // TransferEncoding lists the transfer encodings from outermost to
        // innermost. An empty list denotes the "identity" encoding.
        // TransferEncoding can usually be ignored; chunked encoding is
        // automatically added and removed as necessary when sending and
        // receiving requests.
        TransferEncoding []string

        // Close indicates whether to close the connection after
        // replying to this request (for servers) or after sending this
        // request and reading its response (for clients).
        //
        // For server requests, the HTTP server handles this automatically
        // and this field is not needed by Handlers.
        Close bool

        // For server requests Host specifies the host on which the
        // URL is sought. Per RFC 2616, this is either the value of
        // the "Host" header or the host name given in the URL itself.
        // It may be of the form "host:port". For international domain
        // names, Host may be in Punycode or Unicode form. Use
        // golang.org/x/net/idna to convert it to either format if
        // needed.
        Host string

        // Form contains the parsed form data, including both the URL
        // field's query parameters and the POST or PUT form data.
        // This field is only available after ParseForm is called.
        // The HTTP client ignores Form and uses Body instead.
        Form url.Values

        // PostForm contains the parsed form data from POST, PATCH,
        // or PUT body parameters.
        //
        // This field is only available after ParseForm is called.
        PostForm url.Values

        // MultipartForm is the parsed multipart form, including file uploads.
        // This field is only available after ParseMultipartForm is called.
        MultipartForm *multipart.Form

        // Trailer specifies additional headers that are sent after the request
        // body.
        //
        // For server requests the Trailer map initially contains only the
        // trailer keys, with nil values. (The client declares which trailers it
        // will later send.)  While the handler is reading from Body, it must
        // not reference Trailer. After reading from Body returns EOF, Trailer
        // can be read again and will contain non-nil values, if they were sent
        // by the client.
        //
        // For client requests Trailer must be initialized to a map containing
        // the trailer keys to later send. The values may be nil or their final
        // values. The ContentLength must be 0 or -1, to send a chunked request.
        // After the HTTP request is sent the map values can be updated while
        // the request body is read. Once the body returns EOF, the caller must
        // not mutate Trailer.
        //
        // Few HTTP clients, servers, or proxies support HTTP trailers.
        Trailer Header

        // RemoteAddr allows HTTP servers and other software to record
        // the network address that sent the request, usually for
        // logging. This field is not filled in by ReadRequest and
        // has no defined format. The HTTP server in this package
        // sets RemoteAddr to an "IP:port" address before invoking a
        // handler.
        RemoteAddr string

        // RequestURI is the unmodified Request-URI of the
        // Request-Line (RFC 2616, Section 5.1) as sent by the client
        // to a server. Usually the URL field should be used instead.
        RequestURI string

        // TLS allows HTTP servers and other software to record
        // information about the TLS connection on which the request
        // was received. This field is not filled in by ReadRequest.
        // The HTTP server in this package sets the field for
        // TLS-enabled connections before invoking a handler;
        // otherwise it leaves the field nil.
        // This field is ignored by the HTTP client.
        TLS *tls.ConnectionState

        // Cancel is an optional channel whose closure indicates that the client
        // request should be regarded as canceled. Not all implementations of
        // RoundTripper may support Cancel.
        //
        // For server requests, this field is not applicable.
        //
        // Deprecated: Use the Context and WithContext methods
        // instead. If a Request's Cancel field and context are both
        // set, it is undefined whether Cancel is respected.
        Cancel <-chan struct{}

        // Response is the redirect response which caused this request
        // to be created. This field is only populated during client
        // redirects.
        Response *Response
        // contains filtered or unexported fields
}


*/
