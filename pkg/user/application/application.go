package application

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/versionary-server/pkg/user/application/command"
	"github.com/namhq1989/versionary-server/pkg/user/application/query"
	"github.com/namhq1989/versionary-server/pkg/user/domain"
	"github.com/namhq1989/versionary-server/pkg/user/dto"
)

type (
	Commands interface {
		UpdateMe(ctx *appcontext.AppContext, performerID string, req dto.UpdateMeRequest) (*dto.UpdateMeResponse, error)
	}
	Queries interface {
		GetMe(ctx *appcontext.AppContext, performerID string, _ dto.GetMeRequest) (*dto.GetMeResponse, error)
	}
	App interface {
		Commands
		Queries
	}

	appCommandHandlers struct {
		command.UpdateMeHandler
	}
	appQueryHandler struct {
		query.GetMeHandler
	}
	Application struct {
		appCommandHandlers
		appQueryHandler
	}
)

var _ App = (*Application)(nil)

func New(
	userRepository domain.UserRepository,
) *Application {
	return &Application{
		appCommandHandlers: appCommandHandlers{
			UpdateMeHandler: command.NewUpdateMeHandler(userRepository),
		},
		appQueryHandler: appQueryHandler{
			GetMeHandler: query.NewGetMeHandler(userRepository),
		},
	}
}
