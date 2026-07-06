package join

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func Space(s ...string) string {
	return strings.Join(s, separator.Space)
}
