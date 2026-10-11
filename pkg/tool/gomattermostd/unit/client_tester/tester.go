package client_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/client"
	"testing"
)

type Tester struct {
	Client *client.ClientWithResponses
	t      *testing.T
}
