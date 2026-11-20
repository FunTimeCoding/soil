package fix_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix"
)

func (o *Tester) SecondRunQuiet(directory string) {
	o.t.Helper()
	before := o.GoFiles(directory)
	r := output.NewResultsWithDirectory(directory)
	gofix.RunFormatFix([]string{goModConstant.AllPackages}, directory, false, r)
	assert.Count(o.t, 0, Applied(r.Entries))
	assert.Any(o.t, before, o.GoFiles(directory))
}
