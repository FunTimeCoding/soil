package result

import "github.com/funtimecoding/soil/pkg/tool/gosourced/types/result/references"

func NewFileReferences(
	file string,
	symbols []*references.References,
) *FileReferences {
	return &FileReferences{File: file, Symbols: symbols}
}
