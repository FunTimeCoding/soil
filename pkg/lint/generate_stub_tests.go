package lint

import (
	"fmt"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"path/filepath"
	"sort"
	"strings"
)

func generateStubTests(
	v *virtual_file_system.System,
	fixes *virtual_file_system.System,
	paths []string,
	fix bool,
) {
	var generatedPaths []string

	for _, p := range v.Files() {
		if filepath.Base(p) == library.GeneratedFile {
			generatedPaths = append(generatedPaths, p)
		}
	}

	missing := MissingTestDirectories(paths, generatedPaths)
	directories := make([]string, 0, len(missing))

	for d := range missing {
		directories = append(directories, d)
	}

	sort.Strings(directories)
	useAssert := strings.Contains(
		v.ReadString("go.mod"),
		library.SoilModule,
	)

	for _, d := range directories {
		if fix {
			fmt.Printf(
				"%s: %s (auto-fixed)\n",
				constant.MissingTestFileText,
				d,
			)
		} else {
			fmt.Printf(
				"%s: %s (auto-fixable)\n",
				constant.MissingTestFileText,
				d,
			)
		}

		name := packageNameOf(v, missing[d])
		isToolPackage := strings.HasPrefix(d, "pkg/tool/") &&
			strings.Count(strings.TrimPrefix(d, "pkg/tool/"), "/") == 0
		filename := fmt.Sprintf("%s_test.go", stubTestSuffix(name))

		if isToolPackage {
			filename = "main_test.go"
		}

		fixes.WriteString(
			filepath.Join(d, filename),
			stubTestContent(
				name,
				strings.Contains(d, "/testdata/"),
				isToolPackage,
				useAssert,
			),
		)
	}
}
