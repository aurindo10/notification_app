package entities

import (
	"context"

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

func (c User) Valid(ctx context.Context) (problems map[string]string) {
	return nil
}
