package installed_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"os"
	"path/filepath"
)

func (o *Tester) Main(
	checkout string,
	name string,
	imported string,
) {
	o.t.Helper()
	assert.FatalOnError(
		o.t,
		os.WriteFile(
			filepath.Join(
				checkout,
				constant.CommandDirectory,
				name,
				soil.MainFile,
			),
			[]byte(
				fmt.Sprintf(
					"package main\n\nimport %q\n\nfunc main() {}\n",
					imported,
				),
			),
			0o644,
		),
	)
}
