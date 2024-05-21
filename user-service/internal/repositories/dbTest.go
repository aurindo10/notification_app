package repositories

import (
	"fmt"
	"log"

	"github.com/aurindo10/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewTestDb(dbName string) *gorm.DB {
	config.LoadEnv()

	dbUser := config.GetEnv("POSTGRES_USER", "defaultUser")
	dbPassword := config.GetEnv("POSTGRES_PASSWORD", "defaultPassword")
	dbHost := config.GetEnv("POSTGRES_HOST", "localhost")
	dbPort := config.GetEnv("POSTGRES_PORT", "5432")
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort)

	db, err := gorm.Open(postgres.Open(connStr+" dbname=postgres"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados padrão: %v", err)
	}
	if err := db.Exec("DROP DATABASE " + dbName).Error; err != nil {
		log.Fatalf("Erro ao limpar ao banco de dados padrão: %v", err)
	}
	err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName)).Error
	if err != nil {
		log.Fatalf("Erro ao criar o banco de dados de teste: %v", err)
	}

	testConnStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		dbUser, dbName, dbPassword, dbHost, dbPort)

	testDB, err := gorm.Open(postgres.Open(testConnStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão com o banco de dados de teste: %v", err)
	}
	testDB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	err = testDB.AutoMigrate(&UserDB{})
	if err != nil {
		log.Fatalf("Erro ao realizar a migração do banco de dados de teste: %v", err)
	}
	fmt.Println("Conexão bem-sucedida com o banco de dados de teste!")
	return testDB
}

func ClearDb(dbName string) error {
	config.LoadEnv()
	dbUser := config.GetEnv("POSTGRES_USER", "defaultUser")
	dbPassword := config.GetEnv("POSTGRES_PASSWORD", "defaultPassword")
	dbHost := config.GetEnv("POSTGRES_HOST", "localhost")
	dbPort := config.GetEnv("POSTGRES_PORT", "5432")
	testConnStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		dbUser, dbName, dbPassword, dbHost, dbPort)
	db, err := gorm.Open(postgres.Open(testConnStr), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	return db.Exec("DROP DATABASE " + dbName).Error
}
