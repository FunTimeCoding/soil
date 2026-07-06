package guilds

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/gw2/constant"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
)

func Parse(path string) map[string][]string {
	var result map[string][]string
	s := system.ReadFile(path, constant.GuildFile)

	if false {
		fmt.Printf("Parsing: %s\n", s)
	}

	notation.MustDecode(s, &result, true)

	return result
}
