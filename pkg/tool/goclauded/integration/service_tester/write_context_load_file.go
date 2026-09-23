package service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"os"
	"path/filepath"
)

func (o *Tester) WriteContextLoadFile(identifier string) {
	o.t.Helper()
	errors.PanicOnError(
		os.WriteFile(
			filepath.Join(
				o.Harbor,
				join.Empty(identifier, constant.NotationLogExtension),
			),
			[]byte(fixture.Read("claude", "context-loads.jsonl")),
			0o644,
		),
	)
}
