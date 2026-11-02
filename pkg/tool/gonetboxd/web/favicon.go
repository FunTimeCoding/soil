package web

import (
	_ "embed"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

//go:embed favicon.png
var favicon []byte

func (*Server) favicon(
	w http.ResponseWriter,
	_ *http.Request,
) {
	w.Header().Set(webConstant.ContentType, webConstant.Graphic)
	web.WriteBytes(w, http.StatusOK, favicon)
}
