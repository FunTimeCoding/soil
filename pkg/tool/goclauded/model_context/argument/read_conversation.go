package argument

type ReadConversation struct {
	Session string  `json:"session"`
	Around  string  `json:"around"`
	Count   float64 `json:"count"`
}
