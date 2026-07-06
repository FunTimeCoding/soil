package record

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/telemetry/constant"
)

func newRecord(
	tool string,
	surface string,
	actor string,
	outcome string,
	kind string,
) *Record {
	if !constant.ValidSurface(surface) {
		panic(fmt.Sprintf("invalid telemetry surface: %q", surface))
	}

	if !constant.ValidKind(kind) {
		panic(fmt.Sprintf("invalid telemetry kind: %q", kind))
	}

	return &Record{
		Tool:    tool,
		Surface: surface,
		Actor:   actor,
		Outcome: outcome,
		Kind:    kind,
	}
}
