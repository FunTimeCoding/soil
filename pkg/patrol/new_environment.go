package patrol

import (
	"github.com/funtimecoding/soil/pkg/patrol/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"path/filepath"
)

func NewEnvironment(host string) *Client {
	key := environment.Required(constant.KeyEnvironment)

	return New(
		environment.Required(constant.UserEnvironment),
		host,
		filepath.Dir(key),
		filepath.Base(key),
	)
}
