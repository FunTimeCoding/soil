package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/client"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/face"
	generated "github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/unit/base"
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
	c := client.New(locator.New("gojellyfind.test"), "token")
	defer c.Close()
	m := mock_client.New()
	defer m.Close()

	for _, f := range []face.Client{c, m} {
		assert.NotNil(t, f)
	}
}

func TestTheClientRoundTripsThroughTheDaemon(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	c := client.New(
		locator.New("localhost").Port(s.ContextServer.Port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	x := t.Context()
	assert.FatalOnError(t, c.Reachable(x))
	_, e := c.ListSessions(x)
	assert.FatalOnError(t, e)
	_, f := c.GetItem(x, "charlie")
	assert.True(t, not_found.Is(f))
	g := c.SetVolume(x, "delta", generated.SetVolumeJSONRequestBody{Level: 101})
	assert.True(t, validation.Is(g))
}
