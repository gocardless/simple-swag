package main

import (
	_ "embed"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//go:embed index.html
var indexPage []byte

func main() {
	swaggerPath := flag.String("filename", "~/swagger.json", "a path to the swagger/openapi spec")
	port := flag.Int("port", 9000, "port to serve http over")

	flag.Parse()

	swaggerBytes, err := ioutil.ReadFile(*swaggerPath)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, err = w.Write(indexPage)
		if err != nil {
			log.Fatal(err)
		}
	})
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, err = w.Write(swaggerBytes)
		if err != nil {
			log.Fatal(err)
		}
	})

	log.Printf("Listening on :%d...", *port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", *port), mux)
	if err != nil {
		log.Fatal(err)
	}
}
