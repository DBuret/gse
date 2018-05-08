// GSE - Go Show Env - see https://github.com/DBuret/gse
package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"github.com/DBuret/pathandport"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// direct html, 90's nostalgia
	fmt.Fprint(w, "<html>\n<head>\n")

	//css
	fmt.Fprint(w, "<style type=\"text/css\">\n")
	fmt.Fprintf(w, "h1 {font-family: Calibri, Sans-Serif;}\n")
	fmt.Fprintf(w, "h2 {\n   font-family: Calibri, Sans-Serif;\n   color: darkblue;\n}\n")
	fmt.Fprintf(w, "code {background: #dddddd;display: block}\n")
	fmt.Fprint(w, "</style>\n")

	//title
	fmt.Fprint(w, "<title>GSE</title></head>\n\n<body>\n")

	//lets start
	fmt.Fprintf(w, "<h1>HTTP request</h1>\n")

	// main infos
	fmt.Fprintf(w, "<h2>Method, Host, URL and Protocol</h2>\n<code>\n")
	fmt.Fprintf(w, "Method=%s<br>\n", r.Method)
	fmt.Fprintf(w, "Host=%s<br>\n", r.Host)
	fmt.Fprint(w, "URL=", r.URL, "<br>\n")
	fmt.Fprintf(w, "Proto=%s<br>\n", r.Proto)

	fmt.Fprint(w, "</code>\n\n")

	// http header
	fmt.Fprintf(w, "<h2>HTTP Headers received</h2>\n")
	fmt.Fprintf(w, "<small><i>brackets are not part of the headers</i></small>\n<code>\n")
	sortedHeaders := make([]string, 0, len(r.Header))
	for k := range r.Header {
		sortedHeaders = append(sortedHeaders, k)
	}
	sort.Strings(sortedHeaders)

	for k := range sortedHeaders {
		//fmt.Fprintf(w, " %s = %s\n", k, r.Header[k])
		fmt.Fprintf(w, "%s=", sortedHeaders[k])
		fmt.Fprintf(w, "%s", r.Header[sortedHeaders[k]])
		fmt.Fprint(w, "<br>\n")
	}
	fmt.Fprint(w, "</code>\n\n")

	// body of the http request (for PUT,POST, ...)
	fmt.Fprintf(w, "<h2>Body of the Request (for methods PUT, POST, ...)</h2>\n<code>")

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	s := buf.String() // Does a complete copy of the bytes in the buffer.
	fmt.Fprint(w, s)
	fmt.Fprint(w, "\n</code>\n\n")

	// ENV
	fmt.Fprint(w, "<h1>ENV</h1>\n")
	fmt.Fprint(w, "<code>\n")
	for _, e := range os.Environ() {
		//pair := strings.Split(e, "=")
		fmt.Fprintln(w, e, "<br>")
	}
	fmt.Fprint(w, "</code>\n\n")
	fmt.Fprint(w, "\n</body>\n</html>")
}

func main() {
	var programVersion = "0.3"
	var programName = "gse"

	var uri string
	var port string

	uri, port = pathandport.Parse(programName, "28657")

	log.Printf("Started %s (%s) on port %s with basepath %s ...\n", programName, programVersion, port, uri)

	http.HandleFunc(uri, handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
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
