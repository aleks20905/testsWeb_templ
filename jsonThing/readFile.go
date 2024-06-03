package jsonthing

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadQuestions() ([]Question, error) {
	file, err := os.Open("assets/questions.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var questions []Question
	err = json.Unmarshal(byteValue, &questions)
	if err != nil {
		return nil, err
	}

	// Add auto-generating IDs
	for i := range questions {
		questions[i].ID = i
	}

	return questions, nil
}

func ReadOpenQuestions() ([]OpenQuestion, error) {
	file, err := os.Open("assets/open_questions.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var questions []OpenQuestion
	err = json.Unmarshal(byteValue, &questions)
	if err != nil {
		return nil, err
	}

	// Add auto-generating IDs
	for i := range questions {
		questions[i].ID = i
	}

	return questions, nil
}
