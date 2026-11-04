package notifier

import (
	"context"
	goclaudedConstant "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (n *Notifier) Notify(
	callsign string,
	body string,
	d goclaudedConstant.Delivery,
) error {
	if callsign == "" {
		return nil
	}

	return n.guard.Attempt(
		context.Background(),
		func(context.Context) error {
			return n.connector.Notify(callsign, n.source, body, d)
		},
	)
}
