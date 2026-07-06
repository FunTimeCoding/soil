package web

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"net/http"
	"path/filepath"
)

func (s *Server) reportDelete(
	w http.ResponseWriter,
	r *http.Request,
) {
	system.RemoveFile(
		filepath.Join(s.outputPath, r.PathValue("fileName")),
	)
	http.Redirect(w, r, constant.ReportsPath, http.StatusSeeOther)
}
