package coverage

import (
	"github.com/funtimecoding/soil/pkg/constant"
	crap "github.com/funtimecoding/soil/pkg/crap/constant"
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
		join.Empty(filepath.Base(testPackage), crap.TestBinarySuffix),
	)
	r := run.New()
	r.Directory = root
	r.Start(
		constant.Go,
		crap.TestCommand,
		crap.CompileOnly,
		crap.CoverFlag,
		key_value.Equals(crap.CoverPackage, join.Comma(coverPackages)),
		constant.OutputArgument,
		binary,
		testPackage,
	)

	return binary
}
