package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	log.Printf("Hostname: %s, clientIp: %s", hostname, r.RemoteAddr)
	fmt.Fprintf(w, "Hello World")
}

func main() {
	port := flag.String("port", "8080", "service port")
	flag.Parse()
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":"+*port, nil)
}
