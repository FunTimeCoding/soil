package service

import (
	gw2Constant "github.com/funtimecoding/soil/pkg/gw2/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"strings"
	"time"
)

func (s *Service) EnrichExisting() {
	cutoff := time.Now().AddDate(0, 0, -constant.EnrichCutoffDays)

	for _, t := range system.ReadDirectory(s.elitePath) {
		n := t.Name()

		if !strings.HasSuffix(n, gw2Constant.DetailedWvWKillSuffix) {
			continue
		}

		if len(n) < constant.FileDatePrefixLength {
			continue
		}

		date, e := time.Parse(
			constant.FileDateLayout,
			n[:constant.FileDatePrefixLength],
		)

		if e != nil || date.Before(cutoff) {
			continue
		}

		s.EnrichFile(s.elitePath, n)
	}
}
