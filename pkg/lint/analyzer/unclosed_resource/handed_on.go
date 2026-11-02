package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"github.com/funtimecoding/soil/pkg/lint/types/call_sink"
	"go/ast"
	"go/types"
)

func (s *Summaries) handedOn(
	c *ast.CallExpr,
	f *types.Func,
	argument int,
	sinks map[*ast.CallExpr]call_sink.Sink,
	returned map[*ast.CallExpr]bool,
) bool {
	if len(s.yields[functionKey(f)][argument]) == 0 {
		return false
	}

	if returned[c] {
		return true
	}

	sink, okay := sinks[c]

	return okay && s.relays(
		fact.NewRelayCondition(
			functionKey(f),
			argument,
			functionKey(sink.Callee),
			sink.Argument,
		),
	)
}
