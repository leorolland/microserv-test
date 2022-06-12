package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

var db *MicroservDB
var rootPageGetCount = 0

func OpenWebservice(dbConnection *MicroservDB) {

	db = dbConnection

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
	t := time.Now().Nanosecond()

	if r.URL.Path != "/" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Welcome to a basic Go webserver, here is some useful information :\n")

	fmt.Fprintf(w, "\n[RAM] GET / count: %d\n", rootPageGetCount)
	rootPageGetCount += 1

	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "\npath: %s\n", exe)
	ppid := os.Getppid()
	fmt.Fprintf(w, "pid: %d\n", ppid)

	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "\nhostname: %s\n", name)

	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	inames := []string{}
	for _, i := range interfaces {
		inames = append(inames, i.Name)
	}
	fmt.Fprintf(w, "interfaces: %s\n", strings.Join(inames, ", "))

	if db != nil {
		fmt.Fprintf(w, "\npostgresql connection is open")
		tables, ns := db.ListTables()
		tablesJoined := strings.Join(tables, ", ")
		fmt.Fprintf(w, "\npostgresql tables: %s (%f ms)", tablesJoined, float64(ns)/1e6)
	}

	fmt.Fprintf(w, "\n\nrendered in %f ms", float64(time.Now().Nanosecond()-t)/1e6)
}

func healthz(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "OK")

}
