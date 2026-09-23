package base

import "time"

func (s *Stack) AssertTabGone(identifier string) {
	s.T.Helper()
	deadline := time.Now().Add(5 * time.Second)

	for time.Now().Before(deadline) {
		if !s.TabAlive(identifier) {
			return
		}

		time.Sleep(100 * time.Millisecond)
	}

	s.T.Fatalf("tab %s is still open", identifier)
}
