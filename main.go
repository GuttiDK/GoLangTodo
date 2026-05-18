package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	mux := http.NewServeMux()

	// Serve frontend: check for .next (Next.js) or dist (Vite) or fall back to static
	var frontendDir string
	if _, err := os.Stat("./frontend/.next/standalone"); err == nil {
		// Next.js production build
		// Serve public files and .next
		publicFS := http.FileServer(http.Dir("./frontend/public"))
		nextFS := http.FileServer(http.Dir("./frontend/.next/standalone/frontend"))

		mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Try public first
			if _, err := os.Stat(filepath.Join("./frontend/public", r.URL.Path)); err == nil {
				publicFS.ServeHTTP(w, r)
			} else {
				// Fall through to Next.js
				nextFS.ServeHTTP(w, r)
			}
		}))
	} else if _, err := os.Stat("./dist"); err == nil {
		// Vite build
		frontendDir = "./dist"
		mux.Handle("/", http.FileServer(http.Dir(frontendDir)))
	} else {
		// Development fallback
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
