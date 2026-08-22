package service

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/coverage"
	"path/filepath"
	"time"
)

func (s *Service) Coverage() []*coverage.Server {
	root := coverageRoot()
	var configurationPaths []string

	for _, p := range constant.ConfigurationPaths {
		configurationPaths = append(configurationPaths, filepath.Join(root, p))
	}

	configuration := scan.LoadConfiguration(
		system.FirstFile(configurationPaths...),
	)
	registered := map[string][]string{}

	for server, path := range configuration.ModelContext {
		resolved := filepath.Join(root, path)

		if !system.DirectoryExists(resolved) {
			continue
		}

		names, _ := scan.RegisteredTools(
			virtual_file_system.From(resolved),
			path,
		)
		registered[server] = names
	}

	var calls []tool_call.Call

	for _, e := range s.Sessions() {
		calls = append(calls, s.ToolCalls(e.Identifier)...)
	}

	return coverage.Compute(
		calls,
		registered,
		configuration.ModelContext,
		configuredServers(root),
		configuration.Aliases,
		time.Now(),
	)
}
