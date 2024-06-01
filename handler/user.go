package handler

import (
	user "github.com/aleks20905/testWeb_templ/view/userView"

	"github.com/labstack/echo/v4"
)

func HandlerUserShow(c echo.Context) error {
	return render(c, user.Show())
}
