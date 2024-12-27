package jsonthing

// Question represents a multiple-choice question structure
type Question struct {
	ID       int      `json:"id"` // Auto-generated ID
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   []string `json:"answer"` // Multiple correct answers
}

// OpenQuestion represents an open-ended question structure
type OpenQuestion struct {
	ID       int    `json:"id"` // Auto-generated ID
	Question string `json:"question"`
	Answer   string `json:"answer"` // Single correct answer
}

// SubjectQuestions stores both multiple-choice and open-ended questions for a subject
type SubjectQuestions struct {
	MultipleChoice []Question
	OpenEnded      []OpenQuestion
}
