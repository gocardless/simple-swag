package main

import (
	_ "embed"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

//go:embed index.html
var indexPage []byte

func main() {
	swaggerPath := flag.String("filename", "~/swagger.json", "a path to the swagger/openapi spec")
	port := flag.Int("port", 9000, "port to serve http over")
	host := flag.String("host", "127.0.0.1", "host ip to serve using")

	flag.Parse()

	if _, err := os.Stat(*swaggerPath); os.IsNotExist(err) {
		log.Fatalf("swagger file not found for path %s", *swaggerPath)
		return
	}

	swaggerBytes, err := ioutil.ReadFile(*swaggerPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, err = w.Write(indexPage)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf(fmt.Sprintf("GET - / 200 UserAgent: %s", r.UserAgent()))
		}
	})
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, err = w.Write(swaggerBytes)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf(fmt.Sprintf("GET - /swagger.json 200 UserAgent: %s", r.UserAgent()))
		}
	})

	mux.HandleFunc("/health/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		log.Printf(fmt.Sprintf("GET - /health 200 UserAgent: %s", r.UserAgent()))
	})

	log.Printf("Listening on http://%s:%d ...", *host, *port)
	srv := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", *host, *port),
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		Handler:           mux,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
