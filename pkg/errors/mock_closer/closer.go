package mock_closer

type Closer struct {
	result error
	closed int
}
