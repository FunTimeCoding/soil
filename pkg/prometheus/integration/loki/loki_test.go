//go:build local

package loki

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/loki"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"testing"
	"time"
)

func TestLabelsAndRangeRead(t *testing.T) {
	c := loki.NewEnvironment(false)
	end := time.Now()
	start := end.Add(-time.Hour)
	labels, e := c.Labels(start, end)
	assert.FatalOnError(t, e)
	assert.NotEmpty(t, labels)
	_, _, f := c.QueryRange(
		fmt.Sprintf(
			`{namespace="%s"}`,
			environment.Required(constant.LokiNamespaceEnvironment),
		),
		start,
		end,
		1,
	)
	assert.FatalOnError(t, f)
}

func TestBrokenQueryIsRefusedWithTheReason(t *testing.T) {
	_, e := loki.NewEnvironment(false).Query("{")
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "parse error", e.Error())
}
