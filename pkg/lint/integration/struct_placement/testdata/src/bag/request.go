package bag

type Request struct {
	Name string
}

type Response struct {
	Body string
}

func NewRequest(name string) *Request {
	return &Request{Name: name}
}
