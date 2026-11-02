package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/option"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/double"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/store_tester"
	"testing"
)

type Server struct {
	*double.Set
	t      *testing.T
	daemon *daemon.Daemon
	option *option.Option
	Store  *store_tester.Tester
	*model_context_server.Server
}
