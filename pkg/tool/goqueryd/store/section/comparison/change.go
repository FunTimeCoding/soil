package comparison

type Change struct {
	Title        string
	Removed      []string
	RemovedCount int
	Added        []string
	AddedCount   int
}
