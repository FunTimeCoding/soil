package message

func NewRequest(
	method string,
	parameters any,
	identifier int,
) *Request {
	return &Request{
		Method:     method,
		Parameters: parameters,
		Identifier: identifier,
	}
}
