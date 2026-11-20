package event_tester

func (o *Tester) NextNamed(
	events <-chan Event,
	name string,
) *Event {
	o.t.Helper()

	for {
		if e := o.Next(events); e.Name == name {
			return e
		}
	}
}
