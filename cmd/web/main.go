package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "admin:mysecretpassword@localhost:5432/postgres", "Postgres data source name")
	flag.Parse()

	var app application

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	app.logger = logger

	db, err := sql.Open("postgres", *dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	logger.Info("starting server", "addr", *addr)
	err = http.ListenAndServe(*addr, routes())
	logger.Error(err.Error())
}
