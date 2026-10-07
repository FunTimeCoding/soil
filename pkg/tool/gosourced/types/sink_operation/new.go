package sink_operation

func New(
	kind string,
	path string,
) *Operation {
	return &Operation{Kind: kind, Path: path}
}
