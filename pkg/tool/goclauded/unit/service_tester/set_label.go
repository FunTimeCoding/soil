package service_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) SetLabel(
	sessionIdentifier string,
	key string,
	value string,
) string {
	result, e := o.Label.SetLabel(sessionIdentifier, "", "", key, value)
	assert.FatalOnError(o.t, e)

	return result
}
