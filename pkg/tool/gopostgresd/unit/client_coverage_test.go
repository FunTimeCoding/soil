package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/client"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/face"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/server"
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

func TestTheClientAndItsMockHoldTheFace(t *testing.T) {
	assert.Count(t, 2, []face.Client{&client.Client{}, mock_client.New()})
}
