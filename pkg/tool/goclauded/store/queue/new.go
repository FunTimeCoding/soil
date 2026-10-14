package queue

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func New(
	sessionIdentifier string,
	callsign string,
	kind string,
	body string,
) *Entry {
	return &Entry{
		SessionIdentifier: sessionIdentifier,
		Callsign:          callsign,
		Kind:              kind,
		Body:              body,
		Delivery:          constant.DeliveryQueued,
	}
}
