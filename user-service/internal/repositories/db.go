package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/aurindo10/config"
	_ "github.com/lib/pq"
)

func NewDb() *sql.DB {
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
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão com o banco de dados: %v", err)
	}
	defer db.Close()
	// Verifica se a conexão está funcionando
	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	fmt.Println("Conexão bem-sucedida com o banco de dados!")
	return db
}
