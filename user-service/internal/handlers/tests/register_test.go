package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/aurindo10/internal/repositories"
)

func TestRegisterUser(t *testing.T) {
	clientData := map[string]string{
		"username":  "user1s23",
		"password":  "passwords123",
		"name":      "Johna",
		"last_name": "Dsoe",
		"email":     "johndaosea@example.com",
	}
	clientDataJson, error := json.Marshal(clientData)
	if error != nil {
		t.Fatalf("Erro ao converter dados do cliente para JSON: %v", error)
	}
	resp, err := http.Post("http://localhost:3000/api/v1/registeruser", "application/json", bytes.NewBuffer(clientDataJson))
	if err != nil {
		t.Fatalf("Erro ao fazer solicitação POST: %v", err)
	}
	defer resp.Body.Close()
	var decodedResponse repositories.UserResponseRepository
	if err := json.NewDecoder(resp.Body).Decode(&decodedResponse); err != nil {
		t.Fatalf("Erro ao decodificar: %v", err)
	}
}
func TestProtectionAgainstEmailDuplicated(t *testing.T) {
	clientData := map[string]string{
		"username":  "user1s23",
		"password":  "passwords123",
		"name":      "Johna",
		"last_name": "Dsoe",
		"email":     "johndaosea@example.com",
	}
	clientDataJson, error := json.Marshal(clientData)
	if error != nil {
		t.Fatalf("Erro ao converter dados do cliente para JSON: %v", error)
	}
	resp, err := http.Post("http://localhost:3000/api/v1/registeruser", "application/json", bytes.NewBuffer(clientDataJson))
	if err != nil {
		t.Fatalf("Erro ao fazer solicitação POST: %v", err)
	}
	_, err = http.Post("http://localhost:3000/api/v1/registeruser", "application/json", bytes.NewBuffer(clientDataJson))
	if err != nil {
		t.Fatalf("Erro ao fazer solicitação POST: %v", err)
	}
	defer resp.Body.Close()
	var decodedResponse string
	if err := json.NewDecoder(resp.Body).Decode(&decodedResponse); err != nil {
		t.Fatalf("Erro ao decodificar: %v", err)
	}
	if decodedResponse != "email já existe" {
		t.Fatalf("Teste falhou")
	}
}
func TestEmpitFieldRegister(t *testing.T) {
	clientData := map[string]string{
		"username":  "user1s23",
		"password":  "passwords123",
		"name":      "Johna",
		"last_name": "Dsoe",
	}
	clientDataJson, error := json.Marshal(clientData)
	if error != nil {
		t.Fatalf("Erro ao converter dados do cliente para JSON: %v", error)
	}
	resp, err := http.Post("http://localhost:3000/api/v1/registeruser", "application/json", bytes.NewBuffer(clientDataJson))
	if err != nil {
		t.Fatalf("Erro ao fazer solicitação POST: %v", err)
	}
	defer resp.Body.Close()
	if resp.Status != "400 Bad Request" {
		t.Fatalf("Teste falhou")
	}
}
