package handler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"slices"
	"strconv"

	components "github.com/aleks20905/testWeb_templ/view/components/question"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"

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

func HandleSubmitOpenQuestion(c echo.Context) error {
	userAnswer := c.FormValue("userAnswer")
	question := c.FormValue("question")

	if DEBUG {
		log.Println("userAnswer:", userAnswer)
		log.Println("question :", question)
	}

	var result string
	// Wrap CheckOpenAnswer call in a try block
	check, err := CheckOpenAnswer(question, userAnswer)
	if err != nil {
		log.Println("Error checking answer:", err)
		return err
	}

	if check { // check
		if DEBUG {
			log.Println(" Corect Answer")
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

func CheckOpenAnswer(question, answer string) (bool, error) {

	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return false, err
	}
	defer client.Close()

	// question := " Какво измерват компютърните benchmarks"
	// answer := "производителността на машината, обикновенно чрез отчинане на времето за изичсления "

	// The Gemini 1.5 models are versatile and work with multi-turn conversations (like chat)
	model := client.GenerativeModel("gemini-1.5-flash")
	// Initialize the chat
	cs := model.StartChat()
	cs.History = []*genai.Content{
		&genai.Content{
			Parts: []genai.Part{
				genai.Text(fmt.Sprintf("вярно ли е отговори само с 1 за вярно или 0 за грешно : %s ?", question)),
			},
			Role: "user",
		},
	}

	// fmt.Sprint(resp.Candidates[0].Content.Parts)[1:5] short version

	resp, err := cs.SendMessage(ctx, genai.Text(answer))
	if err != nil {
		return false, err
	}
	if resp != nil {
		caindidates := resp.Candidates
		if caindidates != nil {
			for _, can := range caindidates {
				//fmt.Println("can: ", can)
				content := can.Content
				if content != nil {

					fmt.Println("resp: ", content.Parts[0]) //[True. \n]
					boolValue, err := strconv.ParseBool(fmt.Sprint(content.Parts)[1:2])
					if err != nil {
						return false, nil
					}
					return boolValue, nil
				}
			}
		}
	}
	return false, errors.New("unexpected error in CheckOpenAnswer")
}
