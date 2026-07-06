package tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func (t *Tester) WriteEnvrc(content string) {
	errors.PanicOnError(
		os.WriteFile(t.EnvrcPath, []byte(content), 0755),
	)
}
