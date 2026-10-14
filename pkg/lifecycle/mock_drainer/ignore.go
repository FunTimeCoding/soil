package mock_drainer

func (d *MockDrainer) Ignore() *MockDrainer {
	d.ignore = true

	return d
}
