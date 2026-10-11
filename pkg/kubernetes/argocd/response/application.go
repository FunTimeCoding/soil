package response

type Application struct {
	Metadata Metadata `json:"metadata"`
	Status   Status   `json:"status"`
}
