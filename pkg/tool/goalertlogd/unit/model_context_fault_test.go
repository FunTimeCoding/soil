package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/base"
	"testing"
)

func TestStoreFaultCarriesItsOwnTextOverModelContext(t *testing.T) {
	s := base.New(t)
	c := s.Tools()
	s.Daemon.Close()
	assert.StringContains(
		t,
		`"error": "sql: database is closed"`,
		c.MustCallToolError(
			constant.GetAlerts,
			map[string]any{generativeConstant.ParameterName: "HighMemory"},
		),
	)
}
