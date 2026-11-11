package relocator

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/concern_line"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/identifier"
	"go/token"
	"go/types"
)

func reportExports(
	r *output.Results,
	set *token.FileSet,
	flipped []types.Object,
) {
	for _, o := range flipped {
		concern_line.Add(
			r,
			set,
			o.Pos(),
			constant.ConcernExported,
			true,
			"%s → %s",
			o.Name(),
			identifier.FlipName(o.Name()),
		)
	}
}
