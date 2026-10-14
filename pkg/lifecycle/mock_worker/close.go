package mock_worker

import "fmt"

func (w *MockWorker) Close() {
	*w.Log = append(*w.Log, fmt.Sprintf("stop:%s", w.Tag))
}
