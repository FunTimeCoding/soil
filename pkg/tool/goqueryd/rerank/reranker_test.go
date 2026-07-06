//go:build local

package rerank

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"testing"
)

func TestRank(t *testing.T) {
	r, e := New(
		environment.Required(constant.ModelEnvironment),
		environment.Required(constant.TokenizerEnvironment),
	)
	assert.FatalOnError(t, e)

	defer errors.PanicClose(r)
	results, f := r.Rank(
		"database migration strategy",
		[]string{
			"Schema migrations use ALTER TABLE statements with version tracking to evolve the database safely.",
			"The HTTP client retries failed requests with exponential backoff and jitter.",
			"PostgreSQL supports transactional DDL, allowing migrations to be rolled back on failure.",
		},
	)
	assert.FatalOnError(t, f)
	assert.Greater(t, results[1].Score, results[0].Score)
}
