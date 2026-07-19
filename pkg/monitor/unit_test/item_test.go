package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/monitor/collector"
	"github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/item"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestItem(t *testing.T) {
	c := collector.New(
		"example",
		"example",
		"examples",
		0,
		nil,
	)
	assert.String(
		t,
		"example-1",
		item.New(
			c,
			c.IntegerIdentifier(1),
			constant.Critical,
			upper.Alfa,
			locator.New(web.Example).Path("/1").String(),
			nil,
		).Identifier,
	)
}
