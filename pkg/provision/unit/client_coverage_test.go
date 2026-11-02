package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/provision/client"
	"github.com/funtimecoding/soil/pkg/provision/client/mock_client"
	"github.com/funtimecoding/soil/pkg/provision/face"
	"github.com/funtimecoding/soil/pkg/provision/generated/server"
	"github.com/funtimecoding/soil/pkg/strings/join"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
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

	assert.String(t, "", strings.Join(missing, ", "))
}

func TestTheClientAndItsMockKeepTheFace(t *testing.T) {
	c := client.New(webConstant.Localhost, 1, "alfa")
	m := mock_client.New("bravo")
	clients := []face.Client{c, m}

	for _, v := range clients {
		v.Close()
	}

	assert.Count(t, 2, clients)
	assert.Integer(t, 1, m.Closed)
}
