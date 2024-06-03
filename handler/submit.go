package handler

import (
	"log"
	"net/http"
	"slices"
	"strconv"

	components "github.com/aleks20905/testWeb_templ/view/components/question"

	"github.com/labstack/echo/v4"
)

var DEBUG bool = true

func HandleSubmitQuestion(c echo.Context) error {
	userAnswer := c.FormValue("userAnswer")
	nQuestion, err := strconv.Atoi(c.FormValue("Nquestion"))

	if err != nil {
		log.Println(" Error Invalid question number:")
		return c.String(http.StatusInternalServerError, "Invalid question number") // not working ????
	}
	log.Println("User selected:", userAnswer)
	log.Println("Nquestion :", nQuestion)

	var result string
	if slices.Contains(components.GetQuestionAnswer(nQuestion), userAnswer) {
		log.Println(" Corect Answer:")
		result = "<lable class=\"result corect\"> Corect Answer </lable>"
	} else {
		log.Println(" Wrong Answer ")
		result = "<lable class=\"result wrong\"> Wrong Answer </lable>"
	}

	// Return the result to be swapped in by htmx
	return c.HTML(http.StatusOK, result)
}

func HandleSubmitOpenQuestion(c echo.Context) error {
	userAnswer := c.FormValue("userAnswer")
	nQuestion, err := strconv.Atoi(c.FormValue("Nquestion"))

	if err != nil {
		log.Println(" Error Invalid question number:")                             // cuz returning error doest work
		return c.String(http.StatusInternalServerError, "Invalid question number") // not working ????
	}

	if DEBUG {
		log.Println("User selected:", userAnswer)
		log.Println("Nquestion :", nQuestion)
	}
	var result string
	if slices.Contains(components.GetQuestionAnswer(nQuestion), userAnswer) {
		if DEBUG {
			log.Println(" Corect Answer:")
		}
		result = "<lable class=\"result corect\"> Corect Answer </lable>"
	} else {
		if DEBUG {
			log.Println(" Wrong Answer ")
		}
		result = "<lable class=\"result wrong\"> Wrong Answer </lable>"
	}

	// Return the result to be swapped in by htmx
	return c.HTML(http.StatusOK, result)
}
