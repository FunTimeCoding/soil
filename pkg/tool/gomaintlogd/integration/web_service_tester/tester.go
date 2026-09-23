package web_service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/integration/base"
	"testing"
)

type Tester struct {
	server *base.Server
	Client *client.ClientWithResponses
	t      *testing.T
}
