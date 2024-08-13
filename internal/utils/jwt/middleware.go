package appjwt

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/namhq1989/go-utilities/appcontext"
)

func (j JWT) RequireLoggedIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			ctx   = c.Get("ctx").(*appcontext.AppContext)
			token = strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")
		)

		claims, err := j.ParseAccessToken(ctx, token)
		if claims == nil || err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"data":    nil,
				"code":    "unauthorized",
				"message": "Unauthorized",
			})
		}

		ctx.SetUserID(claims.UserID)
		ctx.SetTimezone(claims.Timezone)
		return next(c)
	}
}
