package flagged

type Kind int

func (k Kind) String() string {
	return "kind"
}

type Alias = Payload

type Reader interface {
	Read() string
}
