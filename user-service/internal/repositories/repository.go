package repositories

import (
	"context"
	"reflect"

	"github.com/aurindo10/internal/entities"
	"github.com/google/uuid"
)

type UserResponseRepository struct {
	Sucess bool      `json:"sucess"`
	Id     uuid.UUID `json:"id"`
}
type Repository interface {
	RegisterUser(p *entities.User) (*UserResponseRepository, error)
	IsEmailAlreadyExists(p *string) (*UserDB, error)
	GetUser(p *LoginParams) (*UserDB, error)
}

type UserDB struct {
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	Name      string    `gorm:"column:name"`
	Last_name *string   `gorm:"column:last_name"`
	Email     *string   `gorm:"column:email"`
}
type LoginParams struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}
type ResponseParamsLogin struct {
	Token *string `json:"token"`
}

func (p LoginParams) Valid(ctx context.Context) (problems map[string]string) {
	problems = make(map[string]string)

	v := reflect.ValueOf(p)
	t := reflect.TypeOf(p)

	// Verifica se o dado passado é uma estrutura
	if v.Kind() != reflect.Struct {
		problems["error"] = "O tipo de dados não é uma estrutura"
		return problems
	}

	// Itera sobre os campos da estrutura
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := t.Field(i).Name

		// Verifica se o campo é nulo
		if field.IsNil() {
			problems[fieldName] = "Campo está nulo"
		}
	}

	return problems
}
