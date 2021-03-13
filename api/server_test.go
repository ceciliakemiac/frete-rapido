package api

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ceciliakemiac/frete-rapido/database"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetMock() (*gorm.DB, sqlmock.Sqlmock) {
	var db *sql.DB
	var pg *gorm.DB
	var mock sqlmock.Sqlmock

	db, mock, _ = sqlmock.New()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	pg, _ = gorm.Open(dialector, &gorm.Config{})

	return pg, mock
}

func InitServer() (*Server, sqlmock.Sqlmock) {
	addr := os.Getenv("API_ADDR")
	pg, mock := GetMock()

	db := &database.Database{
		PG: pg,
	}

	server, _ := NewServer(addr, db)
	return server, mock
}

func TestNewServer(t *testing.T) {
	w := httptest.NewRecorder()
	server, _ := InitServer()

	server.router.ServeHTTP(w, httptest.NewRequest("GET", "/api", nil))

	if w.Code != http.StatusOK {
		t.Error("Did not get expected HTTP status code, got", w.Code)
	}
	if w.Body.String() != `"Hello from Frete Rápido Desafio backend server!"` {
		t.Error("Did not get expected greeting, got ", w.Body.String())
	}
}
