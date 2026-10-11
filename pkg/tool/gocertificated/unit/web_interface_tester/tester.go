package web_interface_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/base"
	"testing"
)

type Tester struct {
	Server *base.Server
	t      *testing.T
}
