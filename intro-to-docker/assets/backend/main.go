package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

const (
	addr string = ":8081"
)

func main() {
	name, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request received at %s\n", r.URL.Path)
		fmt.Fprintf(w, "%s", name)
	})

	handler := cors.Default().Handler(mux)
	log.Printf("Starting to listen on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, handler))

}
