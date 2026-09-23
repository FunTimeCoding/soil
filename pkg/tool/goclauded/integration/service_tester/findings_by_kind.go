package service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/finding"
)

func (o *Tester) FindingsByKind(kind string) []*finding.Finding {
	o.t.Helper()
	all, e := o.Service.Findings()
	assert.FatalOnError(o.t, e)
	var result []*finding.Finding

	for _, i := range all {
		if i.Kind == kind {
			result = append(result, i)
		}
	}

	return result
}
