package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/rejecting_host"
	"testing"
)

func TestRejectedHostIsAnsweredQuietlyOverModelContext(t *testing.T) {
	b := base.New(t, rejecting_host.New(t))
	c := b.Tools()
	assert.String(
		t,
		"host rejected: ip: invalid address",
		c.MustCallToolError(
			constant.AddHost,
			map[string]any{constant.ParameterHost: "alfa"},
		),
	)
	assert.Count(t, 0, b.Reporter.Events())
}
