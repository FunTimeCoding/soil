//go:build browser

package mechanic

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func mark(name string) string {
	return join.Empty(constant.Hash, name)
}
