package request

func NewEvaluate(
	context string,
	expression string,
) *Evaluate {
	return &Evaluate{
		Expression:   expression,
		Target:       NewTarget(context),
		AwaitPromise: true,
	}
}
