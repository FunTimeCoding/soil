package service_tester

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/ensure_result"

func (o *Tester) RegisterActive(
	sessionIdentifier string,
) *ensure_result.Result {
	result := o.Register(sessionIdentifier)
	o.Store.Store.UpdateFields(
		sessionIdentifier,
		map[string]any{"turn_count": 1, "lines": 10},
	)

	return result
}
