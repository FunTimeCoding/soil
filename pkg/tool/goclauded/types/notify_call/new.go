package notify_call

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func New(
	callsign string,
	source string,
	body string,
	d constant.Delivery,
) *Call {
	return &Call{Callsign: callsign, Source: source, Body: body, Delivery: d}
}
