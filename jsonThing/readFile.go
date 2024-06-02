package jsonthing

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadQuestionsFromFile() ([]Question, error) {
	file, err := os.Open("assets/idk.json")
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

	return questions, nil
}
