package message

func NewRequest(
	identifier int,
	method string,
	parameters any,
) *Request {
	return &Request{
		Identifier: identifier,
		Method:     method,
		Parameters: parameters,
	}
}
