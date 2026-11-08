package rename_test_homes

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/example/offline_service"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/packages"
	"os"
)

func RenameTestHomes() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go DIRECTORY MODULE")

		return
	}

	directory := os.Args[1]
	module := os.Args[2]
	w, c, closeService := offline_service.New()
	defer closeService()
	s := packages.New(w, c)

	for _, relative := range testHomeDirectories(directory, "unit_test") {
		renameTestHome(s, directory, module, relative, constant.UnitSegment)
	}

	for _, relative := range testHomeDirectories(
		directory,
		"integration_test",
	) {
		if hasGoFiles(directory, relative) {
			renameTestHome(
				s,
				directory,
				module,
				relative,
				constant.IntegrationSegment,
			)

			continue
		}

		moveFacets(s, directory, module, relative)
	}
}
