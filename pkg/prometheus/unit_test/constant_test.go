package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "up", constant.Up)
	assert.String(t, "daemonset", constant.DaemonSetLabel)
	assert.String(t, "deployment", constant.DeploymentLabel)
	assert.String(t, "endpoint", constant.EndpointLabel)
	assert.String(t, "scope", constant.ScopeLabel)
	assert.String(t, "service", constant.ServiceLabel)
	assert.String(t, "state", constant.StateLabel)
	assert.String(t, "statefulset", constant.StatefulSetLabel)
}
