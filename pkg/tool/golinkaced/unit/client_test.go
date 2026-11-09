package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/client"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/face"
	generated "github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/unit/service_tester"
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
	c := client.New(locator.New("golinkaced.test"), "token")
	defer c.Close()
	m := mock_client.New()
	defer m.Close()

	for _, f := range []face.Client{c, m} {
		assert.NotNil(t, f)
	}
}

func TestClientRoundTripsThroughDaemon(t *testing.T) {
	source := service_tester.NewSource(t)
	c := client.New(
		locator.New("localhost").Port(
			servedPort(t, source, memory.New()),
		).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	x := t.Context()
	assert.FatalOnError(t, c.Reachable(x))
	_, e := c.GetLists(x, &generated.GetListsParams{})
	assert.FatalOnError(t, e)
	source.Failure = not_found.Format("resource not found: links/%d", 12)
	assert.True(t, not_found.Is(c.DeleteLink(x, 12)))
}

func TestClientReadsEveryDeleteAnsweredAsDone(t *testing.T) {
	c := client.New(
		locator.New("localhost").Port(
			servedPort(t, service_tester.NewSource(t), memory.New()),
		).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	x := t.Context()
	assert.FatalOnError(t, c.DeleteLink(x, 12))
	assert.FatalOnError(t, c.DeleteList(x, 12))
	assert.FatalOnError(t, c.DeleteBranch(x, 12))
	assert.FatalOnError(t, c.DeleteTag(x, 12))
	assert.FatalOnError(t, c.DeleteNote(x, 12))
}
