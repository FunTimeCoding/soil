package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/kubernetes/argocd"
	"github.com/funtimecoding/soil/pkg/kubernetes/argocd/mock_client"
	"github.com/funtimecoding/soil/pkg/kubernetes/face"
	"github.com/funtimecoding/soil/pkg/kubernetes/unit/argocd_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"testing"
)

func TestArgocdClientAndItsMockShareFace(t *testing.T) {
	c := argocd.New("argocd.test", 443, true, "token")
	defer c.Close()
	m := mock_client.New(t)
	faces := []face.ArgocdClient{c, m}
	assert.Count(t, 2, faces)
}

func TestArgocdThatListsItsApplicationsIsReachable(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithAnswer(`{"items":[]}`))
	c := argocd_tester.New(t, s.URL)
	defer c.Close()
	assert.FatalOnError(t, c.Reachable(t.Context()))
	assert.String(t, "/api/v1/applications", s.Last().URL.Path)
}
