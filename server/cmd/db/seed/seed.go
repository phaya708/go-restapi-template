package main

import (
	"go-restapi-template/db/seed"
	"go-restapi-template/app/infrastructure/db"
)

func main() {
	db := db.NewDB()
	seed.SeedUser(db)
}