package lint

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/run"
	"path/filepath"
	"strings"
)

func stdlibMatcher() func(string) bool {
	r := run.New()
	r.Panic = false
	root := strings.TrimSpace(r.Start(constant.Go, "env", "GOROOT"))

	if r.Error != nil || root == "" {
		return func(string) bool { return false }
	}

	return func(p string) bool {
		return p != "" &&
			system.DirectoryExists(filepath.Join(root, "src", p))
	}
}
