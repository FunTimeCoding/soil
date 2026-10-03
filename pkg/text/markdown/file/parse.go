package file

import (
	"github.com/funtimecoding/soil/pkg/text/markdown/file/flat"
	"github.com/yuin/goldmark/v2/parser"
)

func (f *File) Parse() *flat.Flat {
	return WalkTree(f.source, parser.New().Parse(*f.source), flat.New())
}
