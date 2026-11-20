package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"testing"
)

func TestOllamaThatAnswersItsHeartbeatIsReachable(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithAnswer("Ollama is running"))
	host, port := upstream_server.HostPort(t, s.URL)
	c := ollama.New(ollama.WithHost(host), ollama.WithPort(port))
	defer c.Close()
	assert.FatalOnError(t, c.Reachable(t.Context()))
	assert.String(t, "/", s.Last().URL.Path)
}

func TestStoppedOllamaIsConnectionFailure(t *testing.T) {
	host, port := upstream_server.HostPort(
		t,
		outage_tester.New(t).RefusedAddress(),
	)
	c := ollama.New(ollama.WithHost(host), ollama.WithPort(port))
	defer c.Close()
	assert.NotNil(t, connection.Classify(c.Reachable(t.Context())))
}
