package request

type PublishSource struct {
	Name      string `json:"Name"`
	Component string `json:"Component,omitempty"`
}
