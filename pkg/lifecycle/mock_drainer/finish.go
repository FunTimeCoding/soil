package mock_drainer

func (d *MockDrainer) finish() {
	d.once.Do(
		func() {
			close(d.finished)
		},
	)
}
