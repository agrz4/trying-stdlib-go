package main

import (
	"fmt"
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("AuthMiddleware")
		next.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	logging := AuthMiddleware(mux)
	log.Fatal(http.ListenAndServe(":8080", logging))
}
