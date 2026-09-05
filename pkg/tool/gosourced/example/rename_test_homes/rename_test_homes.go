package rename_test_homes

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
	"os"
)

func RenameTestHomes() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go DIRECTORY MODULE")

		return
	}

	directory := os.Args[1]
	module := os.Args[2]
	s := service.New(nil)

	for _, relative := range testHomeDirectories(directory, "unit_test") {
		renameTestHome(s, directory, module, relative, "unit")
	}

	for _, relative := range testHomeDirectories(
		directory,
		"integration_test",
	) {
		if hasGoFiles(directory, relative) {
			renameTestHome(s, directory, module, relative, "integration")

			continue
		}

		moveFacets(s, directory, module, relative)
	}
}
