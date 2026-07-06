package sed

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/split/key_value"
	"strings"
)

func ParseAction(raw string) *Action {
	colon := strings.IndexByte(raw, ':')

	if colon < 0 {
		panic(fmt.Sprintf("action missing colon separator: %s", raw))
	}

	path := raw[:colon]
	remainder := raw[colon+1:]
	prefix, value := key_value.Equals(remainder)

	return &Action{Path: path, Prefix: prefix, Value: value}
}
