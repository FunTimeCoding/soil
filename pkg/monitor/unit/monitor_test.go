package unit

import (
	"context"
	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/funtimecoding/soil/pkg/assert"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/monitor/coder/server"
	"github.com/funtimecoding/soil/pkg/monitor/collector"
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/helper"
	"github.com/funtimecoding/soil/pkg/monitor/item"
	"github.com/funtimecoding/soil/pkg/monitor/item/source"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	t.Parallel()
	s := httptest.NewServer(server.Server{Logf: t.Logf})
	defer s.Close()
	x, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	c, _, e := websocket.Dial(
		x,
		s.URL,
		&websocket.DialOptions{Subprotocols: []string{"echo"}},
	)
	assert.FatalOnError(t, e)
	defer func() {
		errors.LogOnError(c.CloseNow())
	}()

	for i := 0; i < 5; i++ {
		assert.FatalOnError(t, wsjson.Write(x, c, map[string]int{"i": i}))
		v := map[string]int{}
		assert.FatalOnError(t, wsjson.Read(x, c, &v))

		if v["i"] != i {
			t.Fatalf("expect %v but got %v", i, v)
		}
	}

	errors.LogOnError(c.Close(websocket.StatusNormalClosure, ""))
}

func TestConstant(t *testing.T) {
	assert.Count(t, 3, library.Severities)
	assert.Count(t, 4, library.Statuses)
}

func TestItemConstant(t *testing.T) {
	assert.Count(t, 15, monitorConstant.Collectors)
}

func TestSource(t *testing.T) {
	assert.NotNil(t, source.New(0, 0, 0, 0, 0, helper.SeverityWeights(0, 0, 0)))
}

func TestItem(t *testing.T) {
	c := collector.New("example", "example", "examples", 0, nil)
	assert.String(
		t,
		"example-1",
		item.New(
			c,
			c.IntegerIdentifier(1),
			library.Critical,
			constant.UpperAlfa,
			locator.New(webConstant.Example).Path("/1").String(),
			nil,
		).Identifier,
	)
}
