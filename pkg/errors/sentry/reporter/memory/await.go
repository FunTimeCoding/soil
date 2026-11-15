package memory

import "time"

func (m *Memory) Await(
	count int,
	within time.Duration,
) {
	deadline := time.Now().Add(within)

	for len(m.Events()) < count && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
}
