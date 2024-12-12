package minichatgo

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	fmt.Printf("%s\n", connStr)
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate() {
	m, err := migrate.New("file://./schema", "postgres://postgres:qwerty@db:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("Could not initializing migrations %v", err.Error())
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migarations failed %v", err)
	}

	fmt.Print("Database migrated successufully")
}