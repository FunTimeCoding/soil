package package_server

import (
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"net/http"
	"path/filepath"
	"strings"
)

func (s *Server) upload(
	w http.ResponseWriter,
	r *http.Request,
) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)

		return
	}

	parts := strings.Split(
		strings.Trim(r.URL.Path, stringsConstant.Slash),
		stringsConstant.Slash,
	)

	if len(parts) != 5 || parts[0] != constant.RoutePrefix {
		http.Error(w, "invalid path", http.StatusBadRequest)

		return
	}

	version := parts[1]
	repo := parts[2]
	arch := parts[3]
	filename := parts[4]

	if !strings.HasSuffix(filename, ".apk") {
		http.Error(w, "must end with .apk", http.StatusBadRequest)

		return
	}

	target := filepath.Join(constant.PackageRoot, version, repo, arch)
	system.MakeDirectory(target)
	file := system.Create(filepath.Join(target, filename))
	defer errors.LogClose(file)
	system.Copy(r.Body, file)
	packagePath := filepath.Join(target, filename)
	signPackage(packagePath, s.signatureKey)

	if f := rebuildIndex(target, s.signatureKey); f != nil {
		httpFail(w, "rebuild fail", f)

		return
	}

	w.WriteHeader(http.StatusCreated)
	writer.Print(w, "uploaded and indexed %s\n", filename)
}
