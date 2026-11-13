package service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"os"
	"path/filepath"
)

func (o *Tester) WriteTranscript(
	identifier string,
	name string,
) {
	o.t.Helper()
	errors.PanicOnError(
		os.WriteFile(
			filepath.Join(
				o.Harbor,
				join.Empty(identifier, soil.NotationLogExtension),
			),
			[]byte(fixture.Read("claude", name)),
			0o644,
		),
	)
}
