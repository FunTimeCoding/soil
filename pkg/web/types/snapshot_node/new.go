package snapshot_node

func New(
	uIdentifier string,
	role string,
	name string,
	value string,
) *Node {
	return &Node{UID: uIdentifier, Role: role, Name: name, Value: value}
}
