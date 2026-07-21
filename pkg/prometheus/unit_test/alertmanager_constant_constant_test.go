package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"
	"testing"
)

func TestAlertmanagerConstant(t *testing.T) {
	assert.String(t, "ALERTMANAGER_HOST", constant.HostEnvironment)
}
