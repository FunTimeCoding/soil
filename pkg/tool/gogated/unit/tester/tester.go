package tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/base"
	"net/http"
)

type Tester struct {
	server *base.Server
	client *http.Client
}
