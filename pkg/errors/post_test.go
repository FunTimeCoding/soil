package errors

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestPost(t *testing.T) {
	Post(
		locator.New(constant.Localhost).String(),
		"something went wrong",
	)
}
