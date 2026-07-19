package unit_test

import (
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestPortMap(t *testing.T) {
	web.PortMap(constant.ListenPort, 80)
}
