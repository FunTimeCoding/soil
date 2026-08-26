package package_server

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/system/run"
	"path/filepath"
)

func rebuildIndex(
	directory string,
	keyName string,
) error {
	packageFiles, e := filepath.Glob(filepath.Join(directory, "*.apk"))

	if e != nil {
		return fmt.Errorf("glob: %w", e)
	}

	if len(packageFiles) == 0 {
		return not_found.Format("no .apk files found in %s", directory)
	}

	arguments := []string{"index", "-vU", "-o", constant.IndexArchive}

	for _, f := range packageFiles {
		arguments = append(arguments, filepath.Base(f))
	}

	r := run.New().NoPanic()
	r.Directory = directory
	r.Execute(append([]string{"apk"}, arguments...)...)

	if r.Error != nil {
		return fmt.Errorf("index: %w", r.Error)
	}

	s := run.New().NoPanic()
	s.Execute(
		"abuild-sign",
		"-k",
		filepath.Join(constant.KeyDirectory, keyName),
		filepath.Join(directory, constant.IndexArchive),
	)

	if s.Error != nil {
		return fmt.Errorf("abuild-sign: %w", s.Error)
	}

	return nil
}
