package fixture

import (
	"github.com/funtimecoding/soil/pkg/project"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func Path(path ...string) string {
	return project.FixturePath(join.Relative(path...))
}
