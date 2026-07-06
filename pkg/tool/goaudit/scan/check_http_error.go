package scan

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/parse"
	"go/ast"
	"path/filepath"
)

func (s *Service) checkHttpError(
	f *ast.File,
	filePath string,
) {
	if parse.HasCall(f, "net/http", "Error") {
		s.Concerns = append(
			s.Concerns,
			concern.NewPackage(
				HttpErrorInStrictKey,
				fmt.Sprintf(
					"%s: http.Error bypasses strict server types",
					filepath.Base(filePath),
				),
				filepath.Dir(filepath.Dir(filePath)),
			),
		)
	}
}
