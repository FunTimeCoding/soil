package service

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/model_context_configuration"
	"os"
	"path/filepath"
)

func configuredServers(root string) map[string]bool {
	result := map[string]bool{}
	b, e := os.ReadFile(filepath.Join(root, constant.ModelContextServersFile))

	if e != nil {
		return result
	}

	var s model_context_configuration.Configuration

	if e := notation.DecodeBytes(b, &s); e != nil {
		return result
	}

	for name := range s.Servers {
		result[name] = true
	}

	return result
}
