package subscription

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func Query(names ...string) string {
	return key_value.Equals(constant.ParameterSubscribe, join.Comma(names))
}
