package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan/audit_configuration"
	"testing"
)

func scanOne(
	t *testing.T,
	v *virtual_file_system.System,
	daemon bool,
) *scan.Service {
	t.Helper()
	c := audit_configuration.New()
	c.Daemon = daemon
	result := scan.Services(v, "test", c)
	assert.Integer(t, 1, len(result))

	return result[0]
}
