//go:build browser

package mechanic

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"os"
)

func extendedSource() string {
	return os.Getenv(constant.ExtendedEnvironment)
}
