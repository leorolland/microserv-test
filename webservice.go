package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func OpenWebservice() {

	http.HandleFunc("/", root)
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/ready", healthz)

	port := GetEnvOrPanic("WS_PORT")

	fmt.Printf("Starting webserver on port %s\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}

func root(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Welcome to a basic Go webserver, here is some useful information :\n")

	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "hostname: %s", name)

}

func healthz(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "OK")

}
