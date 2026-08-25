package snapshot

type Node struct {
	UID                      string
	Role                     string
	Name                     string
	Value                    string
	BackendDOMNodeIdentifier int64
	Children                 []*Node
}
