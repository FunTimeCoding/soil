package relocator

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/relocation"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
)

func Finish(
	r *output.Results,
	plan *relocation.Plan,
	out *sink.Sink,
) (*output.Results, error) {
	reportExports(r, plan.Set, plan.Flipped)
	resymbolExports(out, plan.Set, plan.Flipped)

	return executeMove(r, plan, out)
}
