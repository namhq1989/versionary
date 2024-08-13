package dto

import "github.com/namhq1989/versionary-server/pkg/user/domain"

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (User) FromDomain(user domain.User) User {
	return User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
