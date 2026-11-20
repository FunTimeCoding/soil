package grant_tester

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/request"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/unit/command_tester"
	"os"
	"path/filepath"
)

func (r *Tester) Guard(
	session string,
	directory string,
	command string,
) (string, string, int) {
	r.t.Helper()
	raw, e := json.Marshal(
		request.Input{
			SessionIdentifier: session,
			WorkingDirectory:  directory,
			ToolName:          "Bash",
			ToolInput:         request.ToolInput{Command: command},
		},
	)
	assert.FatalOnError(r.t, e)
	path := filepath.Join(r.t.TempDir(), "input.json")
	assert.FatalOnError(r.t, os.WriteFile(path, raw, 0600))
	input, f := os.Open(path)
	assert.FatalOnError(r.t, f)
	previous := os.Stdin
	os.Stdin = input
	r.t.Cleanup(
		func() {
			os.Stdin = previous
			assert.FatalOnError(r.t, input.Close())
		},
	)

	return command_tester.New(r.t).Run(r.Mock, "guard")
}
