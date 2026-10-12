package coverage

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/system/run"
	"path/filepath"
)

func Compile(
	root string,
	directory string,
	testPackage string,
	coverPackages []string,
) string {
	binary := filepath.Join(
		directory,
		join.Empty(filepath.Base(testPackage), constant.TestBinarySuffix),
	)
	r := run.New()
	r.Directory = root
	r.Start(
		soil.Go,
		constant.TestCommand,
		constant.CompileOnly,
		constant.CoverFlag,
		key_value.Equals(constant.CoverPackage, join.Comma(coverPackages)),
		soil.OutputArgument,
		binary,
		testPackage,
	)

	return binary
}
