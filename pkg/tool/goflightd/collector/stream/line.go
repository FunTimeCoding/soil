package stream

type line struct {
	Timestamp        string `json:"timestamp"`
	Subsystem        string `json:"subsystem"`
	Category         string `json:"category"`
	EventMessage     string `json:"eventMessage"`
	ProcessImagePath string `json:"processImagePath"`
}
