package tool_call

func New(
	name string,
	identifier string,
	timestamp string,
) *Call {
	return &Call{Name: name, Identifier: identifier, Timestamp: timestamp}
}
