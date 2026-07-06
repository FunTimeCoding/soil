package store_tester

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"
)

func (o *Tester) ListRelated(identifier int64) []store.Related {
	o.t.Helper()
	result, e := o.Store.ListRelated(identifier)
	assert.FatalOnError(o.t, e)

	return result
}
