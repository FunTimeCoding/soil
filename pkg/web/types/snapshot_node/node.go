package snapshot_node

type Node struct {
	UID      string
	Role     string
	Name     string
	Value    string
	Children []*Node
}
