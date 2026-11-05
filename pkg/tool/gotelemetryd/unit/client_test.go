package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/client"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/face"
	generated "github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/unit/base"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"reflect"
	"slices"
	"strings"
	"testing"
)

func TestTheClientCoversEveryOperation(t *testing.T) {
	s, e := server.GetSpec()
	assert.FatalOnError(t, e)
	methods := reflect.TypeFor[face.Client]()
	var missing []string

	for _, p := range s.Paths.Map() {
		for _, o := range p.Operations() {
			name := join.Empty(
				strings.ToUpper(o.OperationID[:1]),
				o.OperationID[1:],
			)

			if _, okay := methods.MethodByName(name); !okay {
				missing = append(missing, name)
			}
		}
	}

	slices.Sort(missing)
	assert.String(t, "", strings.Join(missing, ", "))
}

func TestTheClientAndMockHoldTheFace(t *testing.T) {
	c := client.New(locator.New("gotelemetryd.test"), "token")
	defer c.Close()
	m := mock_client.New()
	defer m.Close()

	for _, f := range []face.Client{c, m} {
		assert.NotNil(t, f)
	}
}

func TestTheClientRoundTripsThroughTheDaemon(t *testing.T) {
	s := base.New(t)
	l := locator.New("localhost").Port(s.Port).Insecure()
	c := client.New(l, generativeConstant.ModelContextTestToken)
	defer c.Close()
	x := t.Context()
	assert.FatalOnError(t, c.Reachable(x))
	_, e := c.PostEvent(
		x,
		generated.PostEventJSONRequestBody{
			Tool:    "save_memory",
			Surface: "model_context",
			Actor:   "Blair",
			Outcome: "success",
		},
	)
	assert.FatalOnError(t, e)
	events, f := c.GetEvents(x, &generated.GetEventsParams{})
	assert.FatalOnError(t, f)
	assert.Count(t, 1, events)
	anonymous := client.New(l, "")
	defer anonymous.Close()
	_, g := anonymous.GetEvents(x, &generated.GetEventsParams{})
	assert.Error(t, g)
	assert.StringContains(t, "get events status: 401", g.Error())
}
