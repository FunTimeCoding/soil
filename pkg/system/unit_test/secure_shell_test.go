package unit_test

import (
	"github.com/funtimecoding/soil/pkg/system/secure_shell"
	"github.com/funtimecoding/soil/pkg/system/secure_shell/mock_listener"
	"testing"
)

func TestAccept(t *testing.T) {
	secure_shell.Accept(mock_listener.New())
}
