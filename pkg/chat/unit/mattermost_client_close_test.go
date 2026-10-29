package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/unit/mattermost_client_tester"
	"net/http"
	"testing"
	"time"
)

func TestCloseReleasesTheSocket(t *testing.T) {
	r := mattermost_client_tester.New(t, func(_ *http.ServeMux) {})
	deadline := time.Now().Add(2 * time.Second)

	for time.Now().Before(deadline) && r.Accepted() < 1 {
		time.Sleep(10 * time.Millisecond)
	}

	assert.Integer(t, 1, r.Accepted())
	r.Client.Close()
	deadline = time.Now().Add(2 * time.Second)

	for time.Now().Before(deadline) && r.Released() < 1 {
		time.Sleep(10 * time.Millisecond)
	}

	assert.Integer(t, 1, r.Released())
	r.Client.Close()
}
