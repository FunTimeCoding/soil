package mock_client

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func key(
	namespace string,
	name string,
) string {
	return join.Empty(namespace, stringsConstant.Slash, name)
}
