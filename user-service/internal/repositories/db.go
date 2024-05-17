package repositories

import (
	"fmt"
	"log"

	"github.com/aurindo10/config"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {
	// Carrega as variáveis de ambiente
	config.LoadEnv()

	// Obtém as variáveis de ambiente
	dbUser := config.GetEnv("POSTGRES_USER", "defaultUser")
	dbName := config.GetEnv("POSTGRES_DB", "defaultDB")
	dbPassword := config.GetEnv("POSTGRES_PASSWORD", "defaultPassword")
	dbHost := config.GetEnv("POSTGRES_HOST", "localhost")
	dbPort := config.GetEnv("POSTGRES_PORT", "5432")

	// Monta a string de conexão
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		dbUser, dbName, dbPassword, dbHost, dbPort)

	// Abre a conexão com o banco de dados
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão com o banco de dados: %v", err)
	}

	// Verifica se a conexão está funcionando
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	fmt.Println("Conexão bem-sucedida com o banco de dados!")
	return db
}
