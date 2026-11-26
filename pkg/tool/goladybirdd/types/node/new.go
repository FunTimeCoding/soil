package node

func New(
	uid string,
	role string,
	name string,
	value string,
) *Node {
	return &Node{UID: uid, Role: role, Name: name, Value: value}
}
