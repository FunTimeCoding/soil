package mock_drainer

import "fmt"

func (d *MockDrainer) Close() {
	*d.Log = append(*d.Log, fmt.Sprintf("stop:%s", d.Tag))

	if !d.hold {
		d.finish()
	}
}
