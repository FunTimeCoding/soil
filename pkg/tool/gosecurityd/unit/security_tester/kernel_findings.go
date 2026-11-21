package security_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
)

func (o *Tester) KernelFindings(
	d *system_result.Detail,
) map[string]*system_result.Finding {
	o.t.Helper()
	result := map[string]*system_result.Finding{}

	for _, f := range d.Findings {
		assert.True(o.t, f.Kernel)
		result[f.Vulnerability] = f
	}

	return result
}
