package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"strings"
)

func sanitizeTerm(term string) string {
	return constant.NonAlphanumeric.ReplaceAllString(strings.ToLower(term), "")
}
