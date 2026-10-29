package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/kubernetes/face"
	"k8s.io/client-go/rest"
	"testing"
)

func TestTheClientAndItsMockShareTheFace(t *testing.T) {
	c, e := client.FromConfiguration(
		&rest.Config{Host: "https://kubernetes.test"},
		constant.InCluster,
	)
	assert.FatalOnError(t, e)
	faces := []face.Client{c, mock_client.New()}
	assert.Count(t, 2, faces)
}

func TestTheMockServesRawClients(t *testing.T) {
	m := mock_client.New()
	assert.NotNil(t, m.Nested())
	assert.NotNil(t, m.Dynamic())
	assert.NotNil(t, m.Configuration())
}
