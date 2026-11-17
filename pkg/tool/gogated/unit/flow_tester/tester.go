package flow_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/base"
	"net/http"
	"testing"
)

type Tester struct {
	server *base.Server
	client *http.Client
	t      *testing.T
}
