package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/client"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/face"
	generated "github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/unit/base"
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
	c := client.New(locator.New("godirectoryd.test"), "token")
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
		locator.New("localhost").Port(s.Port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	x := t.Context()
	assert.FatalOnError(t, c.Reachable(x))
	_, e := c.PostUser(
		x,
		generated.PostUserJSONRequestBody{
			Account: "alfa",
			Name:    "Alfa",
			Surname: "Bravo",
		},
	)
	assert.FatalOnError(t, e)
	found, f := c.GetUserAccount(x, "alfa")
	assert.FatalOnError(t, f)
	assert.String(t, "alfa", found.Account)
	_, g := c.GetUserAccount(x, "zulu")
	assert.True(t, not_found.Is(g))
}
