package response

type LogFile struct {
	Version        int            `json:"Version"`
	LogsByFilename map[string]any `json:"LogsByFilename"`
}
