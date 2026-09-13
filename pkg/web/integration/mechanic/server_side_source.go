//go:build browser

package mechanic

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"os"
)

func serverSideSource() string {
	return os.Getenv(constant.ServerSideEnvironment)
}
