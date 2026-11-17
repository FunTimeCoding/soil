package fixture

type Recorder struct {
	Primary   string
	Assigned  string
	Patched   map[string]any
	Deleted   []string
	Created   map[string]any
	Unmatched []string
}
