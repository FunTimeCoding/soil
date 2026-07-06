package run_if

import (
	"github.com/funtimecoding/soil/pkg/system/run"
	"strings"
)

func resolveBase(base string) string {
	if base != "" {
		return base
	}

	r := run.New()
	r.Start(
		"git",
		"rev-parse",
		"--abbrev-ref",
		"--symbolic-full-name",
		"@{upstream}",
	)

	if r.Error == nil {
		upstream := strings.TrimSpace(r.OutputString)

		if upstream != "" {
			return upstream
		}
	}

	return "origin/main"
}
