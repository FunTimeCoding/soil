package argument

type PushDecision struct {
	Session       string   `json:"session"`
	Question      string   `json:"question"`
	DefaultAction string   `json:"default_action"`
	Options       []string `json:"options"`
	Frames        []string `json:"frames"`
}
