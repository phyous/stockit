/**
 * Where to get stock info:
 * http://dev.markitondemand.com/
 * Sample call:
 * http://dev.markitondemand.com/Api/v2/Quote/json?symbol=AAPL
 */
package main

import (
	"fmt"
	"net/http"
	"net/url"
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

func main() {
	http.HandleFunc("/quote/", handler)
	http.ListenAndServe(":8080", nil)
}
