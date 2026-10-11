package web_service_tester

import (
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/unit/base"
	"testing"
)

type Tester struct {
	server   *base.Server
	Client   *client.ClientWithResponses
	Recorder *mock_recorder.Recorder
	t        *testing.T
}
