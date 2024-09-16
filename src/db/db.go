package db

import (
	"fmt"
	"go-backend/src/config"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDB(cfg config.Config) *sqlx.DB {
	db, err := sqlx.Open("postgres", cfg.DBAdress)
	if err != nil {
		fmt.Printf("%s\n", err)
		log.Fatal("Can't connect to databases")
	}

	log.Println("Database connected successfuly")
	return db
}
