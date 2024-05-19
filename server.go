package main

import (
	"log"
	"net/http"
)

func runServer(port string) {
	mux := http.NewServeMux()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello from /test")
	})

	s := http.Server{
		Addr:    port,
		Handler: mux,
	}

	log.Fatalln(s.ListenAndServe())
}
