package secure_shell

import (
	"github.com/funtimecoding/soil/pkg/system/secure_shell/mock_listener"
	"testing"
)

func TestAccept(t *testing.T) {
	Accept(mock_listener.New())
}
