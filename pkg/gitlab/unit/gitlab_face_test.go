package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/gitlab/unit/client_tester"
	"net/http"
	"testing"
)

func TestTheClientAndItsMockKeepTheForge(t *testing.T) {
	r := client_tester.New(t, func(*http.ServeMux) {})
	m := mock_client.New()
	forges := []face.Forge{r.Client, m}
	m.Close()
	assert.Count(t, 2, forges)
	assert.Integer(t, 1, m.Closed)
}
