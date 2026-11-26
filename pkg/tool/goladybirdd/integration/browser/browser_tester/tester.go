package browser_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/service"
	"net/http/httptest"
	"testing"
)

type Tester struct {
	t       *testing.T
	service *service.Service
	site    *httptest.Server
}
