package seed

import (
	"fmt"
	"encoding/csv"
	"os"

	"gorm.io/gorm"
	"go-restapi-template/db/model"
)

func readCSV(filePath string) [][]string {
	csvFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return record	
}

func SeedUser(db *gorm.DB) {
	record := readCSV("db/seed/test_data/users.csv")

	for i:=1; i < len(record); i++ {
		user := model.User{FirstName: record[i][0], LastName: record[i][1]}
		db.Create(&user)
	}
}