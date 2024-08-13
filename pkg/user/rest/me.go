package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/versionary-server/internal/utils/httprespond"
	"github.com/namhq1989/versionary-server/internal/utils/validation"
	"github.com/namhq1989/versionary-server/pkg/user/dto"
)

func (s server) registerMeRoutes() {
	g := s.echo.Group("/api/user")

	g.GET("/me", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.GetMeRequest)
			performerID = ctx.GetUserID()
		)

		resp, err := s.app.GetMe(ctx, performerID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.GetMeRequest](next)
	})

	g.PUT("/me", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.UpdateMeRequest)
			performerID = ctx.GetUserID()
		)

		resp, err := s.app.UpdateMe(ctx, performerID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.UpdateMeRequest](next)
	})
}
