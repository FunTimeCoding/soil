package mock_dump

type Program struct {
	Failure  error
	Tables   int
	Dumped   []string
	Restored []*Restoration
}
