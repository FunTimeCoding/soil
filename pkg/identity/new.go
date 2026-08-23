package identity

func New(
	name string,
	description string,
	usage string,
) *Tool {
	return &Tool{name: name, description: description, usage: usage}
}
