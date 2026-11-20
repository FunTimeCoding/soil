package scan_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan/audit_configuration"
)

func (o *Tester) One(
	v *virtual_file_system.System,
	daemon bool,
) *scan.Service {
	o.t.Helper()
	c := audit_configuration.New()
	c.Daemon = daemon
	result := scan.Services(v, "test", c)
	assert.Integer(o.t, 1, len(result))

	return result[0]
}
