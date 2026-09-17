package pointer

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func CommandPaths(s string) []string {
	name := strings.TrimPrefix(s, constant.Slash)

	if plugin, skill, found := strings.Cut(name, ":"); found {
		return []string{
			join.Empty(".claude/skills/", skill),
			join.Empty("strata/plugin/", plugin, "/skills/", skill),
			join.Empty(
				"../",
				plugin,
				"/strata/plugin/",
				plugin,
				"/skills/",
				skill,
			),
		}
	}

	return []string{join.Empty(".claude/skills/", name)}
}
