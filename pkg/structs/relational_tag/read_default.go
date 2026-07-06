package relational_tag

import (
	"github.com/funtimecoding/soil/pkg/strings/split/key_value"
	"github.com/funtimecoding/soil/pkg/structs/constant"
	"strings"
)

func ReadDefault(v []string) string {
	for _, e := range v {
		if strings.HasPrefix(e, constant.DefaultPrefix) {
			_, a := key_value.Colon(e)

			return a
		}
	}

	return ""
}
