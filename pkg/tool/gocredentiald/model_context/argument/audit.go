package argument

type Audit struct {
	StaleYears float64 `json:"stale_years"`
	Bucket     string  `json:"bucket"`
	Limit      float64 `json:"limit"`
	Offset     float64 `json:"offset"`
}
