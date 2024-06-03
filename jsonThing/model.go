package jsonthing

type Question struct {
	ID       int
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   []string `json:"answer"`
}
type OpenQuestion struct {
	ID       int
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
