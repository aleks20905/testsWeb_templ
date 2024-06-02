package handler

import (
	"log"
	"net/http"
	"slices"
	"strconv"

	components "github.com/aleks20905/testWeb_templ/view/components/question"

	"github.com/labstack/echo/v4"
)

func HandleSubmitQuestion(c echo.Context) error {
	userAnswer := c.FormValue("userAnswer")
	nQuestion, err := strconv.Atoi(c.FormValue("Nquestion"))
	if err != nil {
		panic(err)
	}
	log.Println("User selected:", userAnswer)
	// Process the user's answer, for example, log it or store it
	if slices.Contains(components.GetQuestionAnswer(nQuestion), userAnswer) {
		log.Println("corect Answer")
	}

	// Redirect back to the question page or display a result page
	return c.Redirect(http.StatusSeeOther, "/user")
}
