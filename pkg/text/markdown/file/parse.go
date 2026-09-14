package file

import (
	"github.com/funtimecoding/soil/pkg/text/markdown/file/flat"
	"github.com/yuin/goldmark/v2/parser"
)

func (f *File) Parse() *flat.Flat {
	o := parser.New().Parse(*f.source)
	l := flat.New()

	if false {
		Walk(f.source, o, l)
	}

	WalkTree(f.source, o, l)

	return l
}
