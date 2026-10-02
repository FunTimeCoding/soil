package license

import (
	"github.com/funtimecoding/soil/pkg/strings/split"
	"slices"
)

func hasPrefix(
	identifier string,
	prefixes []string,
) bool {
	return slices.Contains(prefixes, split.Dash(identifier)[0])
}
