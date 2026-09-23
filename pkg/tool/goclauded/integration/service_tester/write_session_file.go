package service_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func (o *Tester) WriteSessionFile(
	identifier string,
	slug string,
) {
	path := filepath.Join(o.Harbor, fmt.Sprintf("%s.jsonl", identifier))
	f := system.Create(path)
	_, e := fmt.Fprintf(
		f,
		"{\"type\":\"user\",\"timestamp\":\"2026-05-21T10:00:00Z\",\"sessionId\":\"%s\",\"slug\":\"%s\",\"cwd\":\"/home/user\",\"gitBranch\":\"main\",\"message\":{\"role\":\"user\",\"content\":\"can you help me fix the login bug\"}}\n{\"type\":\"user\",\"timestamp\":\"2026-05-21T10:02:00Z\",\"sessionId\":\"%s\",\"message\":{\"role\":\"user\",\"content\":\"now update the tests to match\"}}\n",
		identifier,
		slug,
		identifier,
	)
	errors.PanicOnError(e)
	errors.PanicClose(f)
}
