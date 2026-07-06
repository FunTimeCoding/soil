package store

import (
	"github.com/funtimecoding/soil/pkg/gw2/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"strings"
	"time"
)

func (s *Store) enrichExisting() {
	cutoff := time.Now().AddDate(0, 0, -enrichCutoffDays)

	for _, t := range system.ReadDirectory(s.elitePath) {
		n := t.Name()

		if !strings.HasSuffix(n, constant.DetailedWvWKillSuffix) {
			continue
		}

		if len(n) < fileDatePrefixLength {
			continue
		}

		date, dateFail := time.Parse(
			fileDateLayout,
			n[:fileDatePrefixLength],
		)

		if dateFail != nil {
			continue
		}

		if date.Before(cutoff) {
			continue
		}

		s.enrichFile(s.elitePath, n)
	}
}
