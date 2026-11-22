package callback_report

type Task struct {
	Hosts map[string]*Outcome `json:"hosts"`
}
