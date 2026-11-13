package service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"os"
	"path/filepath"
	"time"
)

func (o *Tester) WriteRecentTranscript(
	identifier string,
	name string,
) {
	o.t.Helper()
	stamp := join.Empty(
		`"timestamp": "`,
		o.Store.Clock()().UTC().Format(time.RFC3339),
		`"`,
	)
	errors.PanicOnError(
		os.WriteFile(
			filepath.Join(
				o.Harbor,
				join.Empty(identifier, soil.NotationLogExtension),
			),
			[]byte(
				constant.FixtureTimestamp.ReplaceAllString(
					fixture.Read("claude", name),
					stamp,
				),
			),
			0o644,
		),
	)
}
