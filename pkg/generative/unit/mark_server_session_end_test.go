package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	markServer "github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/generative/model_context/server"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"sync"
	"testing"
)

func TestMarkServerSessionEndFiresOnClose(t *testing.T) {
	var lock sync.Mutex
	var ended []string
	s := markServer.New(identity.Example()).WithSessionEnd(
		func(session string) {
			lock.Lock()
			defer lock.Unlock()
			ended = append(ended, session)
		},
	).Server()
	v := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			server.New(s).Setup(g)
		},
	)
	defer v.Close()
	c := v.Connect(t)
	session := c.SessionIdentifier()
	assert.True(t, session != "")
	c.Close()
	lock.Lock()
	defer lock.Unlock()
	assert.Strings(t, []string{session}, ended)
}
