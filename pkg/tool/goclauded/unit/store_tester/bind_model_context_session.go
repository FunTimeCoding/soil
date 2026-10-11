package store_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) BindModelContextSession(
	name string,
	modelContextSessionIdentifier string,
) {
	assert.FatalOnError(
		o.t,
		o.Store.BindModelContextSession(name, modelContextSessionIdentifier),
	)
}
