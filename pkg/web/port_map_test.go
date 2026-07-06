package web

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestPortMap(t *testing.T) {
	PortMap(constant.ListenPort, 80)
}
