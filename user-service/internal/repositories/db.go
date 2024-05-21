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
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	// Optional: Update existing rows to have a UUID
	// db.Exec("UPDATE user_dbs SET id = uuid_generate_v4() WHERE id IS NULL;")

	// Set default value for the id column to use uuid_generate_v4()
	// db.Exec("ALTER TABLE user_dbs ALTER COLUMN id SET DEFAULT uuid_generate_v4();")
	err = db.AutoMigrate(&UserDB{})
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
