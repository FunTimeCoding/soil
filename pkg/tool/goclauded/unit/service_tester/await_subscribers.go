package service_tester

import "time"

func (o *Tester) AwaitSubscribers(count int) {
	o.t.Helper()
	deadline := time.Now().Add(2 * time.Second)

	for o.Notifier.Subscribers() < count {
		if time.Now().After(deadline) {
			o.t.Fatalf(
				"%d of %d subscribers after two seconds",
				o.Notifier.Subscribers(),
				count,
			)
		}

		time.Sleep(time.Millisecond)
	}
}
