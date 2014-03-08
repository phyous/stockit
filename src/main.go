package main

import (
	"fmt"
	"net/http"
	"net/url"
	"flag"
	"log"
	"os"
	"io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	q, _ := url.ParseQuery(r.URL.RawQuery)
	response, err := http.Get("http://dev.markitondemand.com/Api/v2/Quote/json?symbol=" + q["symbol"][0])
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		defer response.Body.Close()
		contents, _ := ioutil.ReadAll(response.Body)
		fmt.Fprintf(w, "%s", string(contents))
	}
}

var logger *log.Logger
var portVar string
const logFlags int = log.Ldate | log.Ltime | log.Lmicroseconds
func init() {
	flag.StringVar(&portVar, "port", "8080", "port to start server on")
	flag.Parse()

	logger = log.New(os.Stderr, "[stockit] ", logFlags)
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	portStrFmt := ":" + portVar
	logger.Println("Started server on port " + portVar)
	http.HandleFunc("/quote/", handler)
	http.ListenAndServe(portStrFmt, Log(http.DefaultServeMux))
}
