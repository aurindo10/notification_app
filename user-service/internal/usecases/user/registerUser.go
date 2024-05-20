package usecases

import (
	"github.com/aurindo10/internal/entities"
	"github.com/aurindo10/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserUserCase struct {
	repository repositories.Repository
}

func (c *RegisterUserUserCase) Execute(p *entities.UserResquest) (*repositories.UserResponseRepository, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*p.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	encryptedPassword := string(bytes)

	newUser := &entities.User{
		Username:  *p.Username,
		Password:  encryptedPassword,
		Name:      *p.Name,
		Last_name: p.Last_name,
		Email:     p.Email,
	}
	res, err := c.repository.RegisterUser(newUser)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewRegisterUserUserCase(repository repositories.Repository) *RegisterUserUserCase {
	return &RegisterUserUserCase{
		repository: repository,
	}
}
