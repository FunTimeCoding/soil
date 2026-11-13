package module_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix"
	"github.com/funtimecoding/soil/pkg/tool/gofix/option"
	"testing"
)

func RunDefault(
	t *testing.T,
	o *option.Fix,
	r *output.Results,
) {
	t.Helper()
	_, e := gofix.RunDefault(o, r)
	assert.FatalOnError(t, e)
}
