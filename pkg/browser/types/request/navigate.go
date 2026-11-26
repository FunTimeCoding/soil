package request

type Navigate struct {
	Context string `json:"context"`
	Locator string `json:"url"`
	Wait    string `json:"wait"`
}
