package notifier

import "context"

func (n *Notifier) Notify(
	callsign string,
	body string,
) error {
	if callsign == "" {
		return nil
	}

	return n.guard.Attempt(
		context.Background(),
		func(context.Context) error {
			return n.connector.Notify(callsign, n.source, body)
		},
	)
}
