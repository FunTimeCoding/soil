package scan

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan/claude_settings"
	"strings"
)

func permissionEntries(s *claude_settings.Settings) []string {
	if s.Permissions == nil {
		return nil
	}

	var result []string

	for _, list := range [][]string{
		s.Permissions.Allow,
		s.Permissions.Deny,
		s.Permissions.Ask,
	} {
		for _, entry := range list {
			if strings.HasPrefix(entry, constant.ModelContextToolPrefix) {
				result = append(result, entry)
			}
		}
	}

	return result
}
