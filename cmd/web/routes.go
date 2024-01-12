package main

import (
	"net/http"
)

func routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping)

	return mux
}
