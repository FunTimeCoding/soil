package fixture

import (
	"github.com/funtimecoding/soil/pkg/system"
	"os"
)

func File(path ...string) *os.File {
	return system.Open(Path(path...))
}
