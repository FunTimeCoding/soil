package mock_drainer

import "fmt"

func (d *MockDrainer) Start() {
	*d.Log = append(*d.Log, fmt.Sprintf("start:%s", d.Tag))
}
