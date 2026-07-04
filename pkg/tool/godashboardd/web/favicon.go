package web

import (
	_ "embed"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

//go:embed favicon.png
var favicon []byte

func (*Server) favicon(
	w http.ResponseWriter,
	_ *http.Request,
) {
	w.Header().Set(constant.ContentType, constant.Graphic)
	web.WriteBytes(w, http.StatusOK, favicon)
}
