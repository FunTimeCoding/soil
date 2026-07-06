package tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func (t *Tester) WriteProcfile(content string) {
	errors.PanicOnError(
		os.WriteFile(t.ProcfilePath, []byte(content), 0644),
	)
}
