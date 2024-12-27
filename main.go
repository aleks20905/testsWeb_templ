package main

import (
	"fmt"

	"github.com/aleks20905/testWeb_templ/handler"
	"github.com/aleks20905/testWeb_templ/jsonthing"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load()

	jsonthing.LoadAllSubjects()

	app := echo.New()

	app.Static("/css", "css")
	app.Static("/assets", "assets")

	app.GET("/", handler.HandleRedir) // risky very BAD BIG BUGG prob Митака ме накара
	app.GET("/user", handler.HandlerUserShow)
	app.POST("/submit/question", handler.HandleSubmitQuestion)
	app.POST("/submit/Open_question", handler.HandleSubmitOpenQuestion)

	app.GET("/admin", handler.HandlerAdminShow)

	fmt.Println("app starting")
	// fmt.Println("http://localhost:3000/user")

	addr := "localhost:3000"
	app.Logger.Fatal(app.Start(addr))

}
