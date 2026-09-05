package web_interface_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/integration/base"
	"testing"
)

type Tester struct {
	t      *testing.T
	server *base.Server
	base   string
}
