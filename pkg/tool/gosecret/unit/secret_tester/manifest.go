package secret_tester

import (
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosecret/constant"
	"path/filepath"
)

func (o *Tester) Manifest() string {
	o.t.Helper()
	directory := o.t.TempDir()
	system_tester.New(o.t).Write(
		directory,
		"example-secret.yaml",
		constant.TestManifest,
	)

	return filepath.Join(directory, "example-secret.yaml")
}
