package file

import (
	"github.com/funtimecoding/soil/pkg/text/markdown/file/flat"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/text"
)

func (f *File) Parse() *flat.Flat {
	o := goldmark.DefaultParser().Parse(text.NewReader(*f.source))
	l := flat.New()

	if false {
		Walk(f.source, o, l)
	}

	WalkTree(f.source, o, l)

	return l
}
