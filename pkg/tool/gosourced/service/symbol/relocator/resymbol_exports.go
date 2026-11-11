package relocator

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/identifier"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
	"go/token"
	"go/types"
	"path/filepath"
)

func resymbolExports(
	out *sink.Sink,
	set *token.FileSet,
	flipped []types.Object,
) {
	for _, o := range flipped {
		directory := filepath.Dir(set.Position(o.Pos()).Filename)
		out.Resymbol(
			directory,
			o.Name(),
			directory,
			identifier.FlipName(o.Name()),
		)
	}
}
