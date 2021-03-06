package main

import (
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/lucksolutions/go-microservice/healthz"
)

var version = "1.0.0"

func main() {
	log.Println("Starting app...")

	/*
	   Configure the application here. This could include reading ENV variables, loading
	   a configuration files, etc.
	*/
	httpAddr := os.Getenv("HTTP_ADDR")

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	/*
		Setup the /healthz service handler
	*/

	hc := &healthz.Config{
		Hostname: hostname,
	}

	healthzHander, err := healthz.Handler(hc)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/healthz", healthzHander)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, html, hostname, version)
	})

	/*
		Create any additional service handlers here
		...
	*/

	/*
		Start the HTTP Server
	*/
	log.Printf("HTTP Service listening on %s", httpAddr)
	httpErr := http.ListenAndServe(httpAddr, nil)
	if httpErr != nil {
		log.Fatal(httpErr)
	}

}
