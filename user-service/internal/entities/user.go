package entities

import (
	"context"
	"reflect"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Last_name *string   `json:"last_name"`
	Email     *string   `json:"email"`
}
type UserResquest struct {
	Username  *string `json:"username"`
	Password  *string `json:"password"`
	Name      *string `json:"name"`
	Last_name *string `json:"last_name"`
	Email     *string `json:"email"`
}

func (p UserResquest) Valid(ctx context.Context) (problems map[string]string) {
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
