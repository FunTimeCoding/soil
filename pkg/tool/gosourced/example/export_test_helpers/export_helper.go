package export_test_helpers

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/camel"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
	"path"
)

func exportHelper(
	s *service.Service,
	directory string,
	packagePath string,
	targetPath string,
	name string,
	dryRun bool,
) bool {
	exported := exportedName(name)
	file := fmt.Sprintf("%s.go", camel.ToSnake(exported))

	if dryRun {
		fmt.Printf("%s -> %s.%s\n", name, path.Base(targetPath), exported)

		return true
	}

	r, e := s.ChangeVisibility(directory, name, packagePath, "", false)
	errors.PanicOnError(e)

	if reason := refused(r); reason != "" {
		fmt.Printf("REFUSED export %s: %s\n", name, reason)

		return false
	}

	t, f := s.MoveSymbol(
		directory,
		packagePath,
		exported,
		targetPath,
		file,
		true,
		false,
	)
	errors.PanicOnError(f)

	if reason := refused(t); reason != "" {
		fmt.Printf("REFUSED move %s: %s\n", exported, reason)

		return false
	}

	fmt.Printf("%s -> %s.%s\n", name, path.Base(targetPath), exported)

	return true
}
