package handler

import (
	"net/http"

	user "github.com/aleks20905/testWeb_templ/view/userView"

	"github.com/labstack/echo/v4"
)

func HandlerUserShow(c echo.Context) error {
	return render(c, user.Show())
}

func HandleRedir(c echo.Context) error {

	return c.Redirect(http.StatusSeeOther, "/user")
}
