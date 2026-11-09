package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/server_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/store_tester"
	"testing"
)

func TestRevokingTwiceFindsNothingSecondTime(t *testing.T) {
	v := store_tester.NewChain(t)
	r, _, e := v.IssueCertificate(
		server_tester.LeafBody(
			constant.FixtureCommonName,
			[]string{constant.FixtureHost},
		),
	)
	assert.FatalOnError(t, e)
	first, f := v.Revoke(r.Serial)
	assert.FatalOnError(t, f)
	assert.NotNil(t, first.Revoked)
	_, g := v.Revoke(r.Serial)
	assert.True(t, not_found.Is(g))
}
