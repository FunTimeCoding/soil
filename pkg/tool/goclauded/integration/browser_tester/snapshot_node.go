package browser_tester

type SnapshotNode struct {
	UID      string
	Role     string
	Name     string
	Value    string
	Children []*SnapshotNode
}
