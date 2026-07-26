package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strconv"
	"strings"
)

func extractIdentifier(path string) (int64, error) {
	parts := strings.SplitN(path, constant.Slash, 2)

	if len(parts) != 2 {
		return 0, fmt.Errorf("unexpected path: %s", path)
	}

	return strconv.ParseInt(parts[1], 10, 64)
}
