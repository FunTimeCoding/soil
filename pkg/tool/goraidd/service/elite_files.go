package service

import (
	gw2Constant "github.com/funtimecoding/soil/pkg/gw2/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"path/filepath"
	"strings"
)

func (s *Service) EliteFiles(names []string) []string {
	var result []string

	for _, n := range names {
		base := filepath.Base(n)
		timestamp := strings.TrimSuffix(base, filepath.Ext(base))
		result = append(
			result,
			filepath.Join(
				s.elitePath,
				join.Empty(timestamp, gw2Constant.DetailedWvWKillSuffix),
			),
		)
	}

	return result
}
