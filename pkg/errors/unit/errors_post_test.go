package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestPost(t *testing.T) {
	errors.Post(
		locator.New(webConstant.Localhost).String(),
		"something went wrong",
	)
}
