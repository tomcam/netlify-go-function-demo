package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/apex/gateway"
)

func main() {
	usehttp := flag.Bool("http", false, "use http rather than AWS Lambda")
	flag.Parse()
	listener := gateway.ListenAndServe
	if *usehttp {
		listener = http.ListenAndServe
	}

	http.HandleFunc("/", hello)

	log.Fatal(listener(":3000", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Cache-Control", "public, max-age=60")
	b, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Printf("could not dump request: %v", err)
		return
	}
	w.Write(b)
}