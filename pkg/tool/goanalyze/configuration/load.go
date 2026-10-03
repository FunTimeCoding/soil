package configuration

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/constant"
	"go.yaml.in/yaml/v3"
	"io/fs"
	"os"
	"path/filepath"
)

func Load(root string) *Configuration {
	result := &Configuration{}
	b, e := os.ReadFile(filepath.Join(root, constant.ConfigurationPath))

	if errors.Is(e, fs.ErrNotExist) {
		return result
	}

	errors.PanicOnError(e)
	errors.PanicOnError(yaml.Unmarshal(b, result))

	return result
}
