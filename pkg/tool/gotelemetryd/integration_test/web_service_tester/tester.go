package web_service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/integration_test/base"
)

type Tester struct {
	server *base.Server
	Client *client.ClientWithResponses
}
