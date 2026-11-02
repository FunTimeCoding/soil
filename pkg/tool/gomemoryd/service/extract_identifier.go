package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strconv"
	"strings"
)

func extractIdentifier(path string) (int64, error) {
	parts := strings.SplitN(path, stringsConstant.Slash, 2)

	if len(parts) != 2 {
		return 0, validation.New("unexpected path: %s", path)
	}

	return strconv.ParseInt(parts[1], 10, 64)
}
