package goclaude

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"path/filepath"
)

func labelWarning(directory string) string {
	configuration := environment.Optional(
		constant.ConfigurationDirectoryEnvironment,
	)

	if configuration == "" {
		configuration = filepath.Join(system.Home(), constant.ClaudeDirectory)
	}

	project := environment.Optional(constant.ProjectDirectoryEnvironment)

	if project == "" {
		project = directory
	}

	files := LabelAllowance(configuration, project)

	if len(files) == 0 {
		return ""
	}

	return fmt.Sprintf(constant.LabelAllowedWarning, join.Comma(files))
}
