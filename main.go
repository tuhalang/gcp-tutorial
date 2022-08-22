package main

import (
	"flag"
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	port := flag.String("port", "8080", "service port")
	flag.Parse()
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":"+*port, nil)
}
