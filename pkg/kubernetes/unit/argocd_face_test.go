package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/kubernetes/argocd"
	"github.com/funtimecoding/soil/pkg/kubernetes/argocd/mock_client"
	"github.com/funtimecoding/soil/pkg/kubernetes/face"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"net/http"
	"testing"
)

func TestTheArgocdClientAndItsMockShareTheFace(t *testing.T) {
	c := argocd.New("argocd.test", 443, true, "token")
	defer c.Close()
	m := mock_client.New()
	defer m.Close()
	faces := []face.ArgocdClient{c, m}
	assert.Count(t, 2, faces)
}

func TestAnArgocdThatListsItsApplicationsIsReachable(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusOK, `{"items":[]}`)
	c := newArgocdClient(t, s.URL)
	defer c.Close()
	assert.FatalOnError(t, c.Reachable(t.Context()))
	assert.String(t, "/api/v1/applications", last().URL.Path)
}
