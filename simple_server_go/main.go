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
	http.HandleFunc("/500", err500)
	http.HandleFunc("/404", err404)
	http.HandleFunc("/metrics", metrics)
	http.HandleFunc("/health", health)

	log.Printf("Starting server at port %s \n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		return err
	}

	return nil
}

func err500(w http.ResponseWriter, req *http.Request) {
	errCode := http.StatusInternalServerError
	errMsg := fmt.Sprintf("There was an error: %v", errCode)

	requestLog(req.RequestURI, errCode)
	requestError(w, req, errMsg, errCode)
}

func err404(w http.ResponseWriter, req *http.Request) {
	errCode := http.StatusNotFound
	errMsg := fmt.Sprintf("There was an error: %v", errCode)

	requestLog(req.RequestURI, errCode)
	requestError(w, req, errMsg, errCode)
}

func hello(w http.ResponseWriter, req *http.Request) {
	requestLog(req.RequestURI, http.StatusOK)

	fmt.Fprintf(w, "hello, world!\n")
	fmt.Fprintf(w, "try /metrics\n")
	fmt.Fprintf(w, "try /health\n")
}

func health(w http.ResponseWriter, req *http.Request) {
	requestLog(req.RequestURI, http.StatusOK)

	fmt.Fprintf(w, "UP\n")
}

func metrics(w http.ResponseWriter, req *http.Request) {
	requestLog(req.RequestURI, http.StatusOK)

	promhttp.Handler()
}

func requestLog(URL string, errCode int) {
	log.Printf("%s %v", URL, errCode)
}

func requestError(w http.ResponseWriter, req *http.Request, errMsg string, errCode int) {
	requestLog(req.RequestURI, errCode)
	http.Error(w, errMsg, errCode)
}
