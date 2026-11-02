package scope_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/constant"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/option"
	"os"
	"path/filepath"
)

func ChildOptions(patterns ...string) *option.Analyze {
	directory := os.Getenv(constant.FixtureDirectoryEnvironment)

	if directory == "" {
		return nil
	}

	result := option.New()
	result.Root = directory
	result.Full = os.Getenv(constant.FixtureFullEnvironment) == "true"
	result.Index = filepath.Join(directory, "_index")
	result.Patterns = patterns

	return result
}
