package call

func New(
	context string,
	declaration string,
	arguments []any,
) *Call {
	return &Call{
		Context:     context,
		Declaration: declaration,
		Arguments:   arguments,
	}
}
