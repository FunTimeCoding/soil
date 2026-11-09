package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/client"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/face"
	generated "github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/unit/base"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"reflect"
	"slices"
	"strings"
	"testing"
)

func TestClientCoversEveryOperation(t *testing.T) {
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

func TestClientAndMockHoldFace(t *testing.T) {
	c := client.New(locator.New("gosecretd.test"), "token")
	defer c.Close()
	m := mock_client.New()
	defer m.Close()

	for _, f := range []face.Client{c, m} {
		assert.NotNil(t, f)
	}
}

func TestClientRoundTripsThroughDaemon(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	c := client.New(
		locator.New("localhost").Port(s.ContextServer.Port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	x := t.Context()
	assert.FatalOnError(t, c.Reachable(x))
	put, e := c.PutSecret(
		x,
		&generated.PutSecretParams{Path: "alfa/bravo"},
		generated.PutSecretJSONRequestBody{
			Values: map[string]string{"CHARLIE": "delta-value"},
		},
	)
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"CHARLIE"}, put.Keys)
	revealed, f := c.RevealSecret(
		x,
		&generated.RevealSecretParams{Path: "alfa/bravo", Key: "CHARLIE"},
	)
	assert.FatalOnError(t, f)
	assert.String(t, "delta-value", revealed.Value)
	_, g := c.ShowSecret(x, &generated.ShowSecretParams{Path: "alfa/missing"})
	assert.True(t, not_found.Is(g))
	_, h := c.PutSecret(
		x,
		&generated.PutSecretParams{Path: "alfa/bravo"},
		generated.PutSecretJSONRequestBody{Values: map[string]string{}},
	)
	assert.True(t, validation.Is(h))
}
