package database

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostgresConnection() (db *gorm.DB, err error) {
	addr := strings.Split(os.Getenv("POSTGRES_ADDR"), ":")
	user := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DATABASE_NAME")

	credentials := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", addr[0], user, password, dbname, addr[1])

	if db, err = gorm.Open(postgres.Open(credentials), &gorm.Config{
		PrepareStmt: true,
	}); err != nil {
		log.Println("Error getting Postgres Connection: ", err)
		return
	}

	return db, nil
}
