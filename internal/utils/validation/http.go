package validation

import (
	"net/http"

	"github.com/gookit/validate"
	"github.com/labstack/echo/v4"
)

func ValidateHTTPPayload[T any](next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			req T
		)

		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"data":    nil,
				"code":    "bad_request",
				"message": "Bad request",
			})
		}

		if v := validate.Struct(req); !v.Validate() {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"data":    nil,
				"code":    "bad_request",
				"message": v.Errors.OneError().Error(),
			})
		}

		// assign to context
		c.Set("req", req)
		return next(c)
	}
}
