package scratch

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"path/filepath"
	"strings"
)

func Prefix(callsign string) string {
	return join.Empty(
		filepath.Join(
			constant.SystemTemporaryDirectory,
			strings.ToLower(callsign),
		),
		constant.ScratchPrefixSeparator,
	)
}
