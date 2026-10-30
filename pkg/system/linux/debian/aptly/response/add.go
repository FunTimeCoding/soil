package response

type Add struct {
	FailedFiles []string `json:"FailedFiles"`
	Report      Report   `json:"Report"`
}
