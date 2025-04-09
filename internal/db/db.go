package db

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func Init() {
	dsn := fmt.Sprintf("host=localhost user=user password=password dbname=bookdb port=5432 sslmode=disable")
	var err error
	Instance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run migrations
	if err := runMigrations(); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// if err := RunMigrationsDown(); err != nil {
	// 	log.Fatal("Failed to run migrations:", err)
	// }
}

func runMigrations() error {
	databaseURL := "postgres://user:password@localhost:5432/bookdb?sslmode=disable"
	m, err := migrate.New(
		"file://migrations",
		databaseURL,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %v", err)
	}
	defer m.Close()

	// Force version to clean dirty state
	version, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		return fmt.Errorf("failed to get migration version: %v", err)
	}

	if dirty {
		// Force the version to clean state
		if err := m.Force(int(version)); err != nil {
			return fmt.Errorf("failed to force version: %v", err)
		}
	}

	// Run migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %v", err)
	}

	fmt.Println("Migrations went good")
	return nil
}
func RunMigrationsDown() error {
	databaseURL := "postgres://user:password@localhost:5432/bookdb?sslmode=disable"
	m, err := migrate.New(
		"file://migrations",
		databaseURL,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %v", err)
	}
	defer m.Close()

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations down: %v", err)
	}

	return nil
}
