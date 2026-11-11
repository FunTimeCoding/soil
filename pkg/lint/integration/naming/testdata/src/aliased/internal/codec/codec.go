package codec

type Node struct {
	Value string
}

type Unmarshaler interface {
	UnmarshalYAML(node *Node) error
}
