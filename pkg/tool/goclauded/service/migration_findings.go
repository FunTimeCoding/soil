package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/finding"
	"sort"
)

func (s *Service) migrationFindings() []*finding.Finding {
	unkeyed := s.store.CountUnkeyedRows()
	tables := make([]string, 0, len(unkeyed))

	for table := range unkeyed {
		tables = append(tables, table)
	}

	sort.Strings(tables)
	var result []*finding.Finding

	for _, table := range tables {
		if unkeyed[table] == 0 {
			continue
		}

		result = append(
			result,
			finding.New(
				constant.MigrationIncomplete,
				table,
				fmt.Sprintf(
					"%d rows carry no session and fall back to callsign delivery: the backfill declined or has not run",
					unkeyed[table],
				),
				unkeyed[table],
			),
		)
	}

	return result
}
