package service

import (
	gw2Constant "github.com/funtimecoding/soil/pkg/gw2/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"os"
	"strings"
	"time"
)

func (s *Service) EnrichExisting() []error {
	entries, e := os.ReadDir(s.elitePath)

	if e != nil {
		return []error{e}
	}

	cutoff := time.Now().AddDate(0, 0, -constant.EnrichCutoffDays)
	var result []error

	for _, t := range entries {
		n := t.Name()

		if !strings.HasSuffix(n, gw2Constant.DetailedWvWKillSuffix) {
			continue
		}

		if len(n) < constant.FileDatePrefixLength {
			continue
		}

		date, f := time.Parse(
			constant.FileDateLayout,
			n[:constant.FileDatePrefixLength],
		)

		if f != nil || date.Before(cutoff) {
			continue
		}

		if g := s.EnrichFile(s.elitePath, n); g != nil {
			result = append(result, g)
		}
	}

	return result
}
