package handler

import (
	"log"

	"github.com/aleks20905/testWeb_templ/db"
	"github.com/aleks20905/testWeb_templ/db/model"
	user "github.com/aleks20905/testWeb_templ/view/userView"

	"github.com/labstack/echo/v4"
)

func HandlerUserShow(c echo.Context) error {
	return render(c, user.Show(getdata()))
}
func getdata() []model.SensorData {

	data, err := db.GetDataByDevice("Device 2")
	if err != nil {
		log.Fatalf("Error : %v\n", err)

	}

	return data
}
