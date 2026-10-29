package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/face"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	"testing"
)

func TestTheAlertmanagerClientAndItsMockShareTheFace(t *testing.T) {
	p := prometheus.New("127.0.0.1", 0, false, "", "")
	t.Cleanup(p.Close)
	c := alertmanager.New("127.0.0.1", 0, false, "", "", p)
	t.Cleanup(c.Close)
	m := mock_client.New()
	t.Cleanup(m.Close)
	faces := []face.Client{c, m}
	assert.Count(t, 2, faces)
}
