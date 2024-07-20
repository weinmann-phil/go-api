package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/weinmann-phil/go-api/internals/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnvironment() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatalf("Error loading .env file")
	}

}

func SetupDatabase() (*gorm.DB, error) {
	dbHost := os.Getenv(DbHost)
	dbPort := os.Getenv(DbPort)
	dbUser := os.Getenv(DbUser)
	dbPass := os.Getenv(DbPass)
	dbName := os.Getenv(DbName)
	s := []string{"postgres://", dbUser, ":", dbPass, "@", dbHost, ":", dbPort, "/", dbName}

	dsn := strings.Join(s, "")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// logrus.Fatalf("Could not connect to database: %v", err)
		return nil, fmt.Errorf("Could not connect to database: %v", err)
	}

	if err := MigrateDatabase(db); err != nil {
		return nil, fmt.Errorf("Error runnig migration %v", err)
	}

	return db, nil
}

func MigrateDatabase(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&model.User{},
	); err != nil {
		errorMessage := fmt.Sprintf("Error migrating database %v", err)
		logrus.Error(errorMessage)
		return errors.New(errorMessage)
	}

	return nil
}
