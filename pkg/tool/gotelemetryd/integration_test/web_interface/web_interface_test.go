package web_interface

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/integration_test/web_interface_tester"
	"net/http"
	"testing"
)

func TestWebInterface(t *testing.T) {
	o := web_interface_tester.New(t)
	defer o.Close()
	o.AssertStatus(constant.HeatmapPath, http.StatusOK)
	o.AssertStatus(constant.EventsPath, http.StatusOK)
}
