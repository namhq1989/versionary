package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/go-utilities/appcontext"
	appjwt "github.com/namhq1989/versionary-server/internal/utils/jwt"
	"github.com/namhq1989/versionary-server/pkg/user/application"
)

type server struct {
	app  application.App
	echo *echo.Echo
	jwt  appjwt.Operations
}

func RegisterServer(_ *appcontext.AppContext, app application.App, e *echo.Echo, jwt *appjwt.JWT) error {
	var s = server{
		app:  app,
		echo: e,
		jwt:  jwt,
	}

	s.registerMeRoutes()

	return nil
}
