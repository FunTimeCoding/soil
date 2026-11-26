package request

func NewCall(
	context string,
	declaration string,
	arguments []any,
) *Call {
	return &Call{
		Declaration:  declaration,
		Target:       NewTarget(context),
		Arguments:    arguments,
		AwaitPromise: true,
	}
}
