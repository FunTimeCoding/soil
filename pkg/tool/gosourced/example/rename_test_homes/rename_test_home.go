package rename_test_homes

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
	"log"
	"os"
	"path"
	"path/filepath"
)

func renameTestHome(
	s *service.Service,
	directory string,
	module string,
	relative string,
	newName string,
) {
	packagePath := path.Join(module, filepath.ToSlash(relative))
	_, e := s.RenamePackage(directory, packagePath, newName, false)
	errors.PanicOnError(e)

	if _, f := os.Stat(filepath.Join(directory, relative)); f == nil {
		log.Panicf("refused: %s", packagePath)
	}

	fmt.Printf("renamed %s\n", packagePath)
}
