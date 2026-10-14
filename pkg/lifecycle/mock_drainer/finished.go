package mock_drainer

func (d *MockDrainer) Finished() <-chan struct{} {
	return d.finished
}
