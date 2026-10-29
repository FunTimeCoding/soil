package failure

type Failure struct {
	Error           string `json:"error"`
	EventIdentifier string `json:"event_identifier,omitempty"`
}
