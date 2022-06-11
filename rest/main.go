package main

import (
	"net/http"
	"rest/handler"
)

func main() {
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/search/", handler.SearchHandler)
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/playlist/", handler.PlaylistHandler)

	http.ListenAndServe(":8080", nil)
}
