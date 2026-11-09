package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
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
	c := client.New(locator.New("goatlassiand.test"), "token")
	defer c.Close()
	m := mock_client.New()
	defer m.Close()
	faces := []face.Client{c, m}
	assert.Count(t, 2, faces)
}
