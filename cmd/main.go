package main

import (
	"transaction/config"
	"transaction/repo"
	"transaction/rest"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// config init
	cfg := config.New("0204", "postgres", "student", "localhost", 5432)

	db, err := sqlx.Connect("postgres", cfg.GetPgURL())
	if err != nil {
		panic(err)
	}

	repository := repo.New(db)

	h := rest.New(repository)

	h.Run()
}
