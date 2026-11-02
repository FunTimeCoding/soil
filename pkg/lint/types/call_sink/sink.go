package call_sink

import "go/types"

type Sink struct {
	Callee   *types.Func
	Argument int
}
