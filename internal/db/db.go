package db

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func Init() {
	loadEnv()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_SSLMODE"),
	)

	fmt.Println(dsn)

	var err error
	Instance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := runMigrations(); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
}

func loadEnv() {
	// viper.SetConfigFile(".env")
	viper.AutomaticEnv() // Load environment variables

	// Try to read .env file, but don't fail if it doesn't exist
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// Only log fatal if error is not about missing file
			log.Fatal("Error reading config file:", err)
		}
		// If file is not found, just continue with OS environment variables
		log.Println("No .env file found, using OS environment variables")
	}

	// Set default values if neither .env nor OS env vars are set
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_USER", "user")
	viper.SetDefault("DB_PASSWORD", "password")
	viper.SetDefault("DB_NAME", "bookdb")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_SSLMODE", "disable")
}

func runMigrations() error {
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_SSLMODE"),
	)

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
		if err := m.Force(int(version)); err != nil {
			return fmt.Errorf("failed to force version: %v", err)
		}
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %v", err)
	}

	fmt.Println("Migrations went good")
	return nil
}

func RunMigrationsDown() error {
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_SSLMODE"),
	)

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
