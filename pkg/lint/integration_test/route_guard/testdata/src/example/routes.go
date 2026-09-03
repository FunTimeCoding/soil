package example

import "net/http"

func routes() {
	m := http.NewServeMux()
	m.HandleFunc("GET /alfa", nil)
	m.Handle("/bravo", http.NewServeMux())
	http.HandleFunc("/charlie", nil)
}
