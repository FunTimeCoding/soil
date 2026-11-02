package double

import "time"

func (s *Set) AwaitSubscribers(count int) {
	s.t.Helper()
	deadline := time.Now().Add(2 * time.Second)

	for s.Notifier.Subscribers() < count {
		if time.Now().After(deadline) {
			s.t.Fatalf(
				"%d of %d subscribers after two seconds",
				s.Notifier.Subscribers(),
				count,
			)
		}

		time.Sleep(time.Millisecond)
	}
}
