package app

import (
	"amexProject/app/interfaces"
	"amexProject/app/service"
	"amexProject/repository"
	"github.com/jmoiron/sqlx"
)

type App struct {
	ProductSrv  interfaces.ProductSrv
}

func New(tx *sqlx.Tx) *App {
	store := repository.New(tx)
	return &App{
		ProductSrv: service.New(store),
	}
}

