package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/client"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/face"
	generated "github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/unit/base"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"reflect"
	"strings"
	"testing"
)

func TestTheClientCoversEveryOperation(t *testing.T) {
	s, e := server.GetSpec()
	assert.FatalOnError(t, e)
	methods := reflect.TypeFor[*client.Client]()
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

	assert.Strings(t, nil, missing)
}

func TestTheClientAndMockHoldTheFace(t *testing.T) {
	c := client.New(locator.New("gosublimed.test"), "token")
	defer c.Close()
	m := mock_client.New()
	defer m.Close()

	for _, f := range []face.Client{c, m} {
		assert.NotNil(t, f)
	}
}

func TestTheClientRoundTripsThroughTheDaemon(t *testing.T) {
	s := base.New(t)
	c := client.New(
		locator.New("localhost").Port(s.Port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	x := t.Context()
	assert.FatalOnError(t, c.Reachable(x))
	created, e := c.CreateView(
		x,
		generated.CreateViewJSONRequestBody{
			Title:   "notes",
			Content: "alfa alfa",
		},
	)
	assert.FatalOnError(t, e)
	read, f := c.GetView(x, created.ViewId)
	assert.FatalOnError(t, f)
	assert.String(t, "alfa alfa", *read.Text)
	_, g := c.EditView(
		x,
		created.ViewId,
		generated.EditViewJSONRequestBody{OldString: "alfa", NewString: "bravo"},
	)
	assert.True(t, conflict.Is(g))
	_, h := c.GetView(x, 99)
	assert.True(t, not_found.Is(h))
	assert.True(t, not_found.Is(c.CloseView(x, 99)))
}
