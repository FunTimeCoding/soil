package web_tester

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/view"
	"net/http/httptest"
)

func RenderBrand(l *layout.Page) string {
	recorder := httptest.NewRecorder()
	view.New(l).RenderPage(recorder, "", constant.RootPath)

	return recorder.Body.String()
}
