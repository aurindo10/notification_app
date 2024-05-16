package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv carrega as variáveis de ambiente do arquivo .env
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}
}

func GetEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
