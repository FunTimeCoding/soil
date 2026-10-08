package instance_response

type Response struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Active   bool   `json:"active"`
}
