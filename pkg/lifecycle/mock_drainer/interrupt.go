package mock_drainer

import "fmt"

func (d *MockDrainer) Interrupt() {
	*d.Log = append(*d.Log, fmt.Sprintf("interrupt:%s", d.Tag))

	if !d.ignore {
		d.finish()
	}
}
