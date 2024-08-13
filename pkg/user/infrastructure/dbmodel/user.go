package dbmodel

import (
	"time"

	"github.com/namhq1989/go-utilities/timezone"
	"github.com/namhq1989/versionary-server/internal/database"
	apperrors "github.com/namhq1989/versionary-server/internal/utils/error"
	"github.com/namhq1989/versionary-server/pkg/user/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Timezone  string             `bson:"timezone"`
	Providers []UserProvider     `bson:"providers"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

type UserProvider struct {
	Source string `bson:"source"`
	UID    string `bson:"uid"`
}

func (m User) ToDomain() domain.User {
	providers := make([]domain.UserProvider, 0)
	for _, provider := range m.Providers {
		providers = append(providers, domain.UserProvider{
			Source: provider.Source,
			UID:    provider.UID,
		})
	}

	dTimezone, _ := timezone.GetTimezoneData(m.Timezone)

	return domain.User{
		ID:        m.ID.Hex(),
		Name:      m.Name,
		Email:     m.Email,
		Timezone:  *dTimezone,
		Providers: providers,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (User) FromDomain(user domain.User) (*User, error) {
	id, err := database.ObjectIDFromString(user.ID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	providers := make([]UserProvider, 0)
	for _, provider := range user.Providers {
		providers = append(providers, UserProvider{
			Source: provider.Source,
			UID:    provider.UID,
		})
	}

	return &User{
		ID:        id,
		Name:      user.Name,
		Email:     user.Email,
		Timezone:  user.Timezone.Identifier,
		Providers: providers,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
