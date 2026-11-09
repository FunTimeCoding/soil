package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func IsBareSymbol(s string) bool {
	return !IsBareName(s) && constant.BareSymbolSpan.MatchString(
		strings.TrimSuffix(s, constant.CallSuffix),
	)
}
