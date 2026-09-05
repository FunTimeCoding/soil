package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/monitor/collector"
	"github.com/funtimecoding/soil/pkg/monitor/item"
	strings "github.com/funtimecoding/soil/pkg/strings/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestItem(t *testing.T) {
	c := collector.New("example", "example", "examples", 0, nil)
	assert.String(
		t,
		"example-1",
		item.New(
			c,
			c.IntegerIdentifier(1),
			constant.Critical,
			strings.UpperAlfa,
			locator.New(web.Example).Path("/1").String(),
			nil,
		).Identifier,
	)
}
