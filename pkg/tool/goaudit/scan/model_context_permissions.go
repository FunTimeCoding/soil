package scan

import (
	"fmt"
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func ModelContextPermissions(
	base string,
	configuration *Configuration,
) []*concern.Concern {
	if len(configuration.ModelContext) == 0 {
		return nil
	}

	b, e := os.ReadFile(filepath.Join(base, constant.ClaudeSettingsPath))

	if e != nil {
		return nil
	}

	var s Settings

	if e := notation.DecodeBytes(b, &s); e != nil {
		return nil
	}

	servers := make(map[string][]string)

	for _, entry := range permissionEntries(&s) {
		rest := strings.TrimPrefix(entry, generative.ModelContextToolPrefix)
		server, tool, found := strings.Cut(
			rest,
			generative.ModelContextToolSeparator,
		)

		if !found {
			continue
		}

		servers[server] = append(servers[server], tool)
	}

	keys := make([]string, 0, len(servers))

	for server := range servers {
		if _, mapped := configuration.ModelContext[server]; mapped {
			keys = append(keys, server)
		}
	}

	sort.Strings(keys)
	var result []*concern.Concern

	for _, server := range keys {
		path := configuration.ModelContext[server]
		resolved := filepath.Join(base, path)

		if !system.DirectoryExists(resolved) {
			result = append(
				result,
				concern.NewPackage(
					constant.MappedServiceMissingKey,
					constant.MappedServiceMissingText,
					path,
				),
			)

			continue
		}

		names, concerns := RegisteredTools(
			virtual_file_system.From(resolved),
			path,
		)
		result = append(result, concerns...)
		registered := make(map[string]bool)

		for _, name := range names {
			registered[name] = true
		}

		for _, tool := range servers[server] {
			if registered[tool] {
				continue
			}

			result = append(
				result,
				concern.NewPackage(
					constant.StaleToolPermissionKey,
					fmt.Sprintf(
						"%s%s%s%s not registered by %s",
						generative.ModelContextToolPrefix,
						server,
						generative.ModelContextToolSeparator,
						tool,
						path,
					),
					constant.ClaudeSettingsPath,
				),
			)
		}
	}

	return result
}
