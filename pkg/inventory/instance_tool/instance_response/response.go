package instance_response

type Response struct {
	Name   string `json:"name"`
	Host   string `json:"host"`
	Active bool   `json:"active"`
}
