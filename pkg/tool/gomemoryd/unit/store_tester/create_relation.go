package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"time"
)

func (o *Tester) CreateRelation(
	sourceIdentifier int64,
	targetIdentifier int64,
) {
	o.t.Helper()
	assert.FatalOnError(
		o.t,
		o.Store.CreateRelation(
			sourceIdentifier,
			targetIdentifier,
			"",
			time.Now(),
		),
	)
}
