package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/unit/base"
	"net/http"
	"testing"
)

func TestWebInterface(t *testing.T) {
	o := base.New(t)
	o.AssertStatus(http.StatusOK, constant.HeatmapPath)
	o.AssertStatus(http.StatusOK, constant.EventsPath)
}
