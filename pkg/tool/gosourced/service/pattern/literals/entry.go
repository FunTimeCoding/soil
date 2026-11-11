package literals

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/literal"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/pattern_site"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/result/location"
	"go/token"
	"os"
)

func Entry(
	directory string,
	set *token.FileSet,
	contents map[string][]byte,
	site *literal.Site,
	shape string,
) (*pattern_site.Entry, error) {
	position := set.Position(site.Target.Pos())
	content, okay := contents[position.Filename]

	if !okay {
		read, e := os.ReadFile(position.Filename)

		if e != nil {
			return nil, e
		}

		content = read
		contents[position.Filename] = read
	}

	return pattern_site.NewEntry(
		shape,
		collapsedSource(content, set, site.Outer),
		location.New(
			system.RelativePath(directory, position.Filename),
			position.Line,
			site.Package.PkgPath,
		),
	), nil
}
