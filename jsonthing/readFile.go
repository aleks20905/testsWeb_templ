package jsonthing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var SubjectQuestionsMap = make(map[string]*SubjectQuestions)

func LoadAllSubjects() {

	var subjectsDir string = "assets/subjects/"

	// Dynamically discover subjects in the "subjects" directory
	availableSubjects, err := listSubjects(subjectsDir)
	if err != nil || len(availableSubjects) == 0 {
		log.Fatalf("Error discovering subjects: %v", err)
	}

	// Load all subjects into memory and cache them in the map
	for _, subject := range availableSubjects {
		fmt.Printf("Loading questions for subject: %s\n", subject)

		// Load questions for the current subject
		subjectQuestions, err := loadSubjectQuestions(subjectsDir, subject)
		if err != nil {
			log.Printf("Failed to load questions for subject %s: %v", subject, err)
			continue
		}

		SubjectQuestionsMap[subject] = subjectQuestions
	}

}

func loadSubjectQuestions(subjectsDir, subject string) (*SubjectQuestions, error) {
	questionsFile := filepath.Join(subjectsDir, subject, "questions.json")
	openQuestionsFile := filepath.Join(subjectsDir, subject, "open_questions.json")

	var multipleChoice []Question
	var openEnded []OpenQuestion

	// Load multiple-choice questions
	err := loadJSONFile(questionsFile, &multipleChoice)
	if err != nil {
		return nil, fmt.Errorf("error loading multiple-choice questions: %w", err)
	}

	// Auto-generate IDs for multiple-choice questions
	for i := range multipleChoice {
		multipleChoice[i].ID = i + 1 // Simple auto-increment logic
	}

	// Load open-ended questions
	err = loadJSONFile(openQuestionsFile, &openEnded)
	if err != nil {
		return nil, fmt.Errorf("error loading open-ended questions: %w", err)
	}

	// Auto-generate IDs for open-ended questions
	for i := range openEnded {
		openEnded[i].ID = i + 1 // Simple auto-increment logic
	}

	return &SubjectQuestions{
		MultipleChoice: multipleChoice,
		OpenEnded:      openEnded,
	}, nil
}

// Helper function to load a JSON file and unmarshal it into a Go data structure
func loadJSONFile(filePath string, target interface{}) error {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	return json.Unmarshal(byteValue, target)
}

// Function to list all subject directories in the subjects/ folder
func listSubjects(subjectsDir string) ([]string, error) {
	var subjects []string

	files, err := ioutil.ReadDir(subjectsDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read subjects directory: %w", err)
	}

	for _, file := range files {
		if file.IsDir() {
			subjects = append(subjects, file.Name()) // Add subject (folder name) to the list
		}
	}

	return subjects, nil
}

func ReadQuestions() ([]Question, error) {
	file, err := os.Open("assets/subjects/kop_arhitekturi/questions.json")
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
	file, err := os.Open("assets/subjects/kop_arhitekturi/open_questions.json")
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
