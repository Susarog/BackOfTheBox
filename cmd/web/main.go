package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	var app application

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	app.logger = logger

	logger.Info("starting server", "addr", *addr)
	err := http.ListenAndServe(*addr, routes())
	logger.Error(err.Error())
}
