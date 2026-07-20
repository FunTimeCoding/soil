package result

type FileReferences struct {
	File    string        `json:"file"`
	Symbols []*References `json:"symbols"`
}
