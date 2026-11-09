package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"testing"
)

func TestAlertmanagerClientAndItsMockShareFace(t *testing.T) {
	c := alertmanager.New("127.0.0.1", 0, false, "", "")
	t.Cleanup(c.Close)
	m := mock_client.New()
	t.Cleanup(m.Close)
	faces := []face.AlertmanagerClient{c, m}
	assert.Count(t, 2, faces)
}
