package repositories

import (
	"Leadgen/internal/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func New(cfg config.DB) *sqlx.DB {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s  sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.DB, cfg.Password))
	if err != nil {
		fmt.Println("connect failed")
		log.Fatalln(err)
	}
	if err := db.Ping(); err != nil {
		fmt.Println("Ping failed:", err)
		log.Fatalln(err)
	}
	return db
}
