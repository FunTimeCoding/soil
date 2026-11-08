package rename_test_homes

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/packages"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func moveFacets(
	s *packages.Service,
	directory string,
	module string,
	relative string,
) {
	source := filepath.Join(directory, relative)
	target := filepath.Join(filepath.Dir(source), constant.IntegrationSegment)
	errors.PanicOnError(os.MkdirAll(target, 0750))
	entries, e := os.ReadDir(source)
	errors.PanicOnError(e)

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		packagePath := path.Join(
			module,
			filepath.ToSlash(relative),
			entry.Name(),
		)
		targetPath := strings.Replace(
			packagePath,
			"/integration_test/",
			"/integration/",
			1,
		)
		_, f := s.MovePackage(
			context.Background(),
			directory,
			packagePath,
			targetPath,
			false,
		)
		errors.PanicOnError(f)

		if _, g := os.Stat(filepath.Join(source, entry.Name())); g == nil {
			log.Panicf("refused: %s", packagePath)
		}

		fmt.Printf("moved %s\n", packagePath)
	}

	errors.PanicOnError(os.Remove(source))
}
