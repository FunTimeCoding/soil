package result

import "github.com/funtimecoding/soil/pkg/tool/gosourced/service/result/references"

type FileReferences struct {
	File    string                   `json:"file"`
	Symbols []*references.References `json:"symbols"`
}
