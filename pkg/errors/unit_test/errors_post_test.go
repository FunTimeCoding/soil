package unit_test

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestPost(t *testing.T) {
	errors.Post(
		locator.New(constant.Localhost).String(),
		"something went wrong",
	)
}
