package constant

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "up", Up)
	assert.String(t, "daemonset", DaemonSetLabel)
	assert.String(t, "deployment", DeploymentLabel)
	assert.String(t, "endpoint", EndpointLabel)
	assert.String(t, "scope", ScopeLabel)
	assert.String(t, "service", ServiceLabel)
	assert.String(t, "state", StateLabel)
	assert.String(t, "statefulset", StatefulSetLabel)
}
