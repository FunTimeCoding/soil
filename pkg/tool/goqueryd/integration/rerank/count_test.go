//go:build local

package rerank

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/rerank"
	"testing"
)

func TestBgeCountsTableRowAndAllowance(t *testing.T) {
	t.Setenv(constant.RerankSequenceEnvironment, "512")
	r, e := rerank.NewModel(constant.BgeRerankerBase)
	assert.FatalOnError(t, e)

	defer errors.PanicClose(r)
	assert.Integer(t, 0, r.Count(""))
	assert.Integer(t, 23, r.Count(constant.FixtureTableRow))
	assert.Integer(t, 476, r.Allowance())
}

func TestGteCountsTableRowAndAllowance(t *testing.T) {
	t.Setenv(constant.RerankSequenceEnvironment, "1024")
	r, e := rerank.NewModel(constant.GteRerankerModernBertBase)
	assert.FatalOnError(t, e)

	defer errors.PanicClose(r)
	assert.Integer(t, 0, r.Count(""))
	assert.Integer(t, 29, r.Count(constant.FixtureTableRow))
	assert.Integer(t, 989, r.Allowance())
}
