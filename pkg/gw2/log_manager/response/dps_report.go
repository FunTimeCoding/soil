package response

type DpsReport struct {
	Locator         any `json:"Url"`
	ProcessingError any `json:"ProcessingError"`
	UploadTime      any `json:"UploadTime"`
}
