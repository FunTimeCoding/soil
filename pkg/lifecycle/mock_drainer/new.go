package mock_drainer

func New(
	tag string,
	log *[]string,
) *MockDrainer {
	return &MockDrainer{
		Tag:      tag,
		Log:      log,
		finished: make(chan struct{}),
	}
}
