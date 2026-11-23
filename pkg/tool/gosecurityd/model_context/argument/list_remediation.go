package argument

type ListRemediation struct {
	Selection
	Score *float64 `json:"score"`
}
