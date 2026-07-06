package system

import (
	"github.com/funtimecoding/soil/pkg/system/join"
	"os"
)

func OpenHome(subPath string) *os.File {
	return Open(join.Absolute(Home(), subPath))
}
