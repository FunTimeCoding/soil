package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance"
	"go.yaml.in/yaml/v3"
	"testing"
)

func TestInstanceConnectionReadsEndpointAndLink(t *testing.T) {
	var i instance.Instance
	errors.PanicOnError(yaml.Unmarshal([]byte(constant.FixtureInstance), &i))
	assert.String(t, "alertmanager-operated.observability", i.Alertmanager.Host)
	assert.Integer(t, 9093, i.Alertmanager.Port)
	assert.False(t, i.Alertmanager.Secure)
	assert.String(t, "a.example", i.Alertmanager.Link.Host)
	assert.True(t, i.Alertmanager.Link.Secure)
	assert.Integer(t, 9090, i.Prometheus.Port)
	assert.String(t, "", i.Prometheus.Link.Host)
}
