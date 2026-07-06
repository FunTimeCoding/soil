package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strconv"
	"strings"
)

func extractIdentifier(path string) (int64, error) {
	parts := strings.SplitN(path, separator.Slash, 2)

	if len(parts) != 2 {
		return 0, fmt.Errorf("unexpected path: %s", path)
	}

	return strconv.ParseInt(parts[1], 10, 64)
}
