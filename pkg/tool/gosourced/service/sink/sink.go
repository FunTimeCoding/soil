package sink

import (
	"github.com/funtimecoding/soil/pkg/lint/types/symbol_move"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/sink_operation"
)

type Sink struct {
	root        string
	rewrites    []*sink_operation.Operation
	operations  []*sink_operation.Operation
	relocations []*sink_operation.Operation
	symbols     []*symbol_move.Move
	dropped     []string
}
