package repository

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

func InitialiseDB() *sqlx.Tx {
	if db, err := sqlx.Connect("pgx", "postgresql://root:secret@localhost:5432/amex_example?sslmode=disable"); err != nil {
		log.Fatal("could not connect to the database, error: ", err)
	} else {
		return db.MustBegin()
	}
	return nil
}

