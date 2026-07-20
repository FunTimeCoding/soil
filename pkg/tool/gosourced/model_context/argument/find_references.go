package argument

type FindReferences struct {
	PackagePath string  `json:"package_path"`
	Symbol      string  `json:"symbol"`
	Receiver    string  `json:"receiver"`
	File        string  `json:"file"`
	Limit       float64 `json:"limit"`
	Offset      float64 `json:"offset"`
}
