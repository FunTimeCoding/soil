package export_test_helpers

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func ExportTestHelpers() {
	if len(os.Args) < 6 {
		fmt.Println(
			"Usage: go run main.go DIRECTORY MODULE RELATIVE TARGET NAMES [apply]",
		)

		return
	}

	directory := os.Args[1]
	module := os.Args[2]
	relative := os.Args[3]
	target := os.Args[4]
	content, e := os.ReadFile(os.Args[5])
	errors.PanicOnError(e)
	dryRun := len(os.Args) < 7 || os.Args[6] != "apply"
	packagePath := path.Join(module, filepath.ToSlash(relative))
	targetPath := path.Join(packagePath, target)
	moved := 0
	asked := 0

	for _, line := range strings.Split(string(content), "\n") {
		name := strings.TrimSpace(line)

		if name == "" {
			continue
		}

		asked++

		if exportHelper(
			service.New(nil),
			directory,
			packagePath,
			targetPath,
			name,
			dryRun,
		) {
			moved++
		}
	}

	verb := "would move"

	if !dryRun {
		verb = "moved"
	}

	fmt.Printf("%s %d of %d helpers into %s\n", verb, moved, asked, target)
}
