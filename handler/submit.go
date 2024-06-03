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
		return c.String(http.StatusInternalServerError, "Invalid question number")
	}
	log.Println("User selected:", userAnswer)

	// Determine if the answer is correct
	var result string
	if slices.Contains(components.GetQuestionAnswer(nQuestion), userAnswer) {
		result = "<lable class=\"result corect\"> Corect Answer </lable>"
	} else {
		result = "<lable class=\"result wrong\"> Wrong Answer </lable>"
	}

	// Return the result to be swapped in by htmx
	return c.HTML(http.StatusOK, result)
}
