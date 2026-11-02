package goreach

import (
	lintConstant "github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/environment"
	goclaudeConstant "github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"os"
	"path/filepath"
)

func transcriptPath(
	flag string,
	positionals []string,
) string {
	if flag != "" {
		return flag
	}

	identifier := environment.Optional(
		goclaudeConstant.SessionIdentifierEnvironment,
	)

	if len(positionals) > 0 {
		identifier = resolveSession(positionals[0])
	}

	if identifier == "" {
		system.Exitf(1, "no session: name one, or run inside a session\n")
	}

	home, e := os.UserHomeDir()

	if e != nil {
		system.Exitf(1, "%s\n", e)
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
		system.Exitf(1, "no transcript for session %s\n", identifier)
	}

	return matches[0]
}
