package repositories

import (
	"fmt"
	"log"

	"github.com/aurindo10/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewTestDb(dbName string) *gorm.DB {
	// Carrega as variáveis de ambiente
	config.LoadEnv()

	// Obtém as variáveis de ambiente para o banco de dados de teste
	dbUser := config.GetEnv("POSTGRES_USER", "defaultUser")
	dbPassword := config.GetEnv("POSTGRES_PASSWORD", "defaultPassword")
	dbHost := config.GetEnv("POSTGRES_HOST", "localhost")
	dbPort := config.GetEnv("POSTGRES_PORT", "5432")
	// Monta a string de conexão ao banco de dados padrão 'postgres'
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort)

	// Conecta ao banco de dados padrão 'postgres' para criar o banco de dados de teste
	db, err := gorm.Open(postgres.Open(connStr+" dbname=postgres"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados padrão: %v", err)
	}
	if err := ClearDb(db, dbName); err != nil {
		log.Fatalf("Erro ao limpar ao banco de dados padrão: %v", err)
	}
	// Cria o banco de dados de teste
	err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName)).Error
	if err != nil {
		log.Fatalf("Erro ao criar o banco de dados de teste: %v", err)
	}

	// Monta a string de conexão para o banco de dados de teste
	testConnStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		dbUser, dbName, dbPassword, dbHost, dbPort)

	// Abre a conexão com o banco de dados de teste
	testDB, err := gorm.Open(postgres.Open(testConnStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão com o banco de dados de teste: %v", err)
	}
	// Realiza a migração automática do banco de dados
	testDB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	err = testDB.AutoMigrate(&UserDB{})
	if err != nil {
		log.Fatalf("Erro ao realizar a migração do banco de dados de teste: %v", err)
	}
	fmt.Println("Conexão bem-sucedida com o banco de dados de teste!")
	return testDB
}

func ClearDb(db *gorm.DB, dbName string) error {
	err := db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s;", dbName)).Error
	if err != nil {
		log.Fatalf("Erro ao remover o banco de dados de teste: %v", err)
		return err
	}
	fmt.Println("Banco de dados de teste removido com sucesso!")
	return nil
}
