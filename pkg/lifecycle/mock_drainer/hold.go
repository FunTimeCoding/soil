package mock_drainer

func (d *MockDrainer) Hold() *MockDrainer {
	d.hold = true

	return d
}
