package comparison

type Comparison struct {
	Changes []*Change
	Removed int
	Added   int
}
