package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/secure_shell"
	"github.com/funtimecoding/soil/pkg/system/secure_shell/mock_listener"
	"testing"
)

func TestAccept(t *testing.T) {
	owned := mock_listener.New()
	defer errors.PanicClose(owned)
	secure_shell.Accept(owned)
}
