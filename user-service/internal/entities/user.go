package entities

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID
	Username  string
	Password  string
	Name      string
	Last_name *string
	Email     *string
}
