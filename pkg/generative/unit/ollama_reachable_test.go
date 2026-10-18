package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"net/http"
	"testing"
)

func TestAnOllamaThatAnswersItsHeartbeatIsReachable(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusOK, "Ollama is running")
	host, port := upstream_tester.HostPort(t, s.URL)
	c := ollama.New(ollama.WithHost(host), ollama.WithPort(port))
	defer c.Close()
	assert.FatalOnError(t, c.Reachable(t.Context()))
	assert.String(t, "/", last().URL.Path)
}

func TestAStoppedOllamaIsAConnectionFailure(t *testing.T) {
	s, _ := upstream_tester.New(t, http.StatusOK, "")
	host, port := upstream_tester.HostPort(t, s.URL)
	s.Close()
	c := ollama.New(ollama.WithHost(host), ollama.WithPort(port))
	defer c.Close()
	assert.NotNil(t, connection.Classify(c.Reachable(t.Context())))
}
