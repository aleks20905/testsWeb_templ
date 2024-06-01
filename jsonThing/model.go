package model

type Question struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   []string `json:"answer"`
}
