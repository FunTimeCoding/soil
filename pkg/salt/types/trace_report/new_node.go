package trace_report

func NewNode(
	name string,
	orphaned bool,
) *Node {
	return &Node{Name: name, Orphaned: orphaned}
}
