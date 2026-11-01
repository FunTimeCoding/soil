package sink

import "github.com/funtimecoding/soil/pkg/tool/gosourced/types/sink_operation"

type Sink struct {
	root       string
	rewrites   []*sink_operation.Operation
	operations []*sink_operation.Operation
	dropped    []string
}
