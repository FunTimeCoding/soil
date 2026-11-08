package service_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) AddLabel(
	sessionIdentifier string,
	key string,
	value string,
) string {
	result, e := o.Label.AddLabel(sessionIdentifier, "", "", key, value)
	assert.FatalOnError(o.t, e)

	return result
}
