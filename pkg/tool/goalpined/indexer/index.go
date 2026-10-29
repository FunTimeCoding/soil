package indexer

import (
	"fmt"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"path/filepath"
)

func (i *Indexer) Index(
	directory string,
	keyPath string,
) error {
	packageFiles, e := filepath.Glob(filepath.Join(directory, "*.apk"))

	if e != nil {
		return fmt.Errorf("glob packages: %w", e)
	}

	arguments := []string{
		"apk",
		"index",
		"-vU",
		"-o",
		systemConstant.AlpineIndexArchive,
	}

	for _, f := range packageFiles {
		arguments = append(arguments, filepath.Base(f))
	}

	r := run.New().NoPanic()
	r.Directory = directory
	r.Execute(arguments...)

	if r.Error != nil {
		return fmt.Errorf("index packages: %w", r.Error)
	}

	s := run.New().NoPanic()
	s.Execute(
		"abuild-sign",
		"-k",
		keyPath,
		filepath.Join(directory, systemConstant.AlpineIndexArchive),
	)

	if s.Error != nil {
		return fmt.Errorf("sign index: %w", s.Error)
	}

	return nil
}
