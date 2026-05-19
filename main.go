package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	mux := http.NewServeMux()

	var frontendDir string
	if _, err := os.Stat("./frontend/.next/standalone"); err == nil {
		publicFS := http.FileServer(http.Dir("./frontend/public"))
		nextFS := http.FileServer(http.Dir("./frontend/.next/standalone/frontend"))

		mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if _, err := os.Stat(filepath.Join("./frontend/public", r.URL.Path)); err == nil {
				publicFS.ServeHTTP(w, r)
			} else {
				nextFS.ServeHTTP(w, r)
			}
		}))
	} else if _, err := os.Stat("./dist"); err == nil {
		frontendDir = "./dist"
		mux.Handle("/", http.FileServer(http.Dir(frontendDir)))
	} else {
		frontendDir = "./static"
		mux.Handle("/", http.FileServer(http.Dir(frontendDir)))
	}

	mux.HandleFunc("/api/todos", todosHandler)
	mux.HandleFunc("/api/todos/", todoHandler)

	addr := ":8080"
	log.Printf("Starting server at http://localhost%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
