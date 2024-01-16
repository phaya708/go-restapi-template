package main

import (
	"go-restapi-template/db/reset"
	"go-restapi-template/app/infrastructure/db"
)

func main() {
	db := db.NewDB()
	reset.Reset(db)
}