package turn

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/turn/session_key"
)

func Session(x context.Context) string {
	result, _ := x.Value(session_key.Key{}).(string)

	return result
}
