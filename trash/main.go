package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintln(w, r.Method, r.URL, r.Proto); err != nil {
		log.Print(err)
	}
	for k, v := range r.Header {
		if _, err := fmt.Fprintf(w, "Header[%q] = %q\n", k, v); err != nil {
			log.Print(err)
		}
	}
	if _, err := fmt.Fprintf(w, "Host = %q\n", r.Host); err != nil {
		log.Print(err)
	}
	if _, err := fmt.Fprintf(w, "ReamoteAddr = %q\n", r.RemoteAddr); err != nil {
		log.Print(err)
	}
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		if _, err := fmt.Fprintf(w, "Form[%q] = %q\n", k, v); err != nil {
			log.Print(err)
		}
	}
}
func main() {
	http.HandleFunc("/", handler)
	//http.HandleFunc("/count")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
