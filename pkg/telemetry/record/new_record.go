package record

import "fmt"

func newRecord(
	tool string,
	surface string,
	actor string,
	outcome string,
	kind string,
) *Record {
	if !validSurface(surface) {
		panic(fmt.Sprintf("invalid telemetry surface: %q", surface))
	}

	if !validKind(kind) {
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
