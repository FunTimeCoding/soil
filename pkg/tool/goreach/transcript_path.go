package goreach

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	lintConstant "github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/environment"
	goclaudeConstant "github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"os"
	"path/filepath"
)

func TranscriptPath(
	flag string,
	positionals []string,
) (string, error) {
	if flag != "" {
		return flag, nil
	}

	identifier := environment.Optional(
		goclaudeConstant.SessionIdentifierEnvironment,
	)

	if len(positionals) > 0 {
		resolved, e := resolveSession(positionals[0])

		if e != nil {
			return "", e
		}

		identifier = resolved
	}

	if identifier == "" {
		return "", not_selected.Format(
			"no session: name one, or run inside a session",
		)
	}

	home, e := os.UserHomeDir()

	if e != nil {
		return "", e
	}

	matches := system.Glob(
		filepath.Join(
			home,
			lintConstant.ProjectsDirectory,
			"*",
			join.Empty(identifier, lintConstant.TranscriptSuffix),
		),
	)

	if len(matches) == 0 {
		return "", not_found.Format("no transcript for session %s", identifier)
	}

	return matches[0], nil
}
