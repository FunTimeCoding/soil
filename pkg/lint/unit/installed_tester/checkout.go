package installed_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"os"
	"path/filepath"
	"strings"
)

func (o *Tester) Checkout(
	directory string,
	module string,
	required []string,
	commands ...string,
) string {
	o.t.Helper()
	result := filepath.Join(o.parent, directory)
	assert.FatalOnError(o.t, os.MkdirAll(result, 0o755))

	for _, c := range commands {
		assert.FatalOnError(
			o.t,
			os.MkdirAll(
				filepath.Join(result, constant.CommandDirectory, c),
				0o755,
			),
		)
	}

	var b strings.Builder
	_, e := fmt.Fprintf(&b, "module %s\n", module)
	assert.FatalOnError(o.t, e)

	for _, r := range required {
		_, f := fmt.Fprintf(&b, "require %s v0.0.0\n", r)
		assert.FatalOnError(o.t, f)
	}

	assert.FatalOnError(
		o.t,
		os.WriteFile(
			filepath.Join(result, goModConstant.ModFile),
			[]byte(b.String()),
			0o644,
		),
	)

	return result
}
