package testutil
import (
	"fmt"
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

func NewTestDB() *gorm.DB {
	dbHost := "db" // docker-compose.ymlで定義したコンテナ名
	dbUser := "postgres"
	dbPassword := "password"
	dbName := "test"
	dbPort := "5432"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
		return nil
	}

	return db
}



