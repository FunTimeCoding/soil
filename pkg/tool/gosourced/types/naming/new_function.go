package naming

func NewFunction(
	name string,
	receiver string,
) *Function {
	return &Function{Name: name, Receiver: receiver}
}
