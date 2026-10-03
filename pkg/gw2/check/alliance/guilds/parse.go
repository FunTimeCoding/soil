package guilds

import (
	"github.com/funtimecoding/soil/pkg/gw2/constant"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
)

func Parse(path string) map[string][]string {
	var result map[string][]string
	notation.MustDecode(
		system.ReadFile(path, constant.GuildFile),
		&result,
		true,
	)

	return result
}
