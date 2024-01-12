package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	err := http.ListenAndServe(*addr, routes())
	if err != nil {
		log.Println(err.Error())
	}
}
