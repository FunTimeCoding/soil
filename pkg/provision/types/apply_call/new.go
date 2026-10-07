package apply_call

func New(
	parameters map[string]any,
	triggerSource string,
) *Call {
	return &Call{Parameters: parameters, TriggerSource: triggerSource}
}
