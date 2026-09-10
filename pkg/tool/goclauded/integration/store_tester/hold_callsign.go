package store_tester

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (o *Tester) HoldCallsign(
	sessionIdentifier string,
	callsign string,
) {
	o.EnsureSession(sessionIdentifier)
	o.Store.UpdateFields(
		sessionIdentifier,
		map[string]any{"name": callsign, constant.Callsign: callsign},
	)
}
