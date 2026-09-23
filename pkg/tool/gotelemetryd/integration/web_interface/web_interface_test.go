package web_interface

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/integration/web_interface_tester"
	"net/http"
	"testing"
)

func TestWebInterface(t *testing.T) {
	o := web_interface_tester.New(t)
	o.AssertStatus(http.StatusOK, constant.HeatmapPath)
	o.AssertStatus(http.StatusOK, constant.EventsPath)
}
