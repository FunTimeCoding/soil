package service

import (
	"github.com/funtimecoding/soil/pkg/errors/load_failure"
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"golang.org/x/mod/modfile"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func markSettled(
	directory string,
	e error,
) error {
	failure, okay := load_failure.As(e)

	if !okay {
		return e
	}

	b, f := os.ReadFile(filepath.Join(directory, goModConstant.ModFile))

	if f != nil {
		return e
	}

	relative, inside := strings.CutPrefix(
		failure.Package,
		modfile.ModulePath(b),
	)

	if !inside {
		return e
	}

	entries, g := os.ReadDir(filepath.Join(directory, relative))

	if g != nil {
		return e
	}

	newest := time.Time{}

	for _, entry := range entries {
		if i, h := entry.Info(); h == nil && i.ModTime().After(newest) {
			newest = i.ModTime()
		}
	}

	failure.Settled = time.Since(newest) >= constant.LoadSettled

	return e
}
