package client_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/client"
	"testing"
)

type Tester struct {
	Client *client.Client
	t      *testing.T
}
