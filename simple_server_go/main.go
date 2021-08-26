package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	port = "8080"
)

func main() {
	if err := startWebServer(); err != nil {
		log.Println(err)
	}
}

func startWebServer() error {
	http.HandleFunc("/", hello)
	http.HandleFunc("/metrics", metrics)
	http.HandleFunc("/health", health)

	log.Printf("Starting server at port %s \n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		return err
	}

	return nil
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello, world!\n")
	fmt.Fprintf(w, "try /metrics\n")
	fmt.Fprintf(w, "try /health\n")
}

func health(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "UP\n")
}

func metrics(w http.ResponseWriter, req *http.Request) {
	promhttp.Handler()
}
