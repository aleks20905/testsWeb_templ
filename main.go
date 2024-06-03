package main

import (
	"fmt"

	"github.com/aleks20905/testWeb_templ/handler"
	"github.com/labstack/echo/v4"
)

func main() {

	app := echo.New()

	app.Static("/css", "css")
	app.Static("/assets", "assets")

	app.GET("/user", handler.HandlerUserShow)
	app.POST("/submit/question", handler.HandleSubmitQuestion)
	app.POST("/submit/Open_question", handler.HandleSubmitOpenQuestion)

	fmt.Print("app starting")
	app.Start(":3000")

}
