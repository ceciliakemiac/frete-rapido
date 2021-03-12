package database

import (
	"fmt"
	"log"
	"os"

	"github.com/ceciliakemiac/frete-rapido/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	PG *gorm.DB
}

func SetupDatabase() (pg *Database, err error) {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DATABASE_NAME")

	credentials := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(credentials), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		log.Println("Error getting Postgres Connection: ", err)
		return nil, err
	}

	err = db.AutoMigrate(&model.Quote{}, &model.Freight{})
	if err != nil {
		log.Println("Error migrating datagabe: ", err)
		return nil, err
	}

	pg = &Database{
		PG: db,
	}

	return pg, nil
}
