package module_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix"
	"github.com/funtimecoding/soil/pkg/tool/gofix/option"
)

func (o *Tester) RunDefault(
	p *option.Fix,
	r *output.Results,
) {
	o.t.Helper()
	_, e := gofix.RunDefault(p, r)
	assert.FatalOnError(o.t, e)
}
