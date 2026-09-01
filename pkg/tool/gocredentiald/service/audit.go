package service

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/service/audit_report"
	"github.com/tobischo/gokeepasslib/v3"
)

func (s *Service) Audit(staleYears int) *audit_report.Report {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.refresh()
	report := audit_report.New()
	cutoff := s.clock().AddDate(-staleYears, 0, 0)
	seen := map[string]int{}
	s.client.Walk(
		func(
			path string,
			_ *gokeepasslib.Group,
			entry *gokeepasslib.Entry,
		) {
			described := describe(path, entry)

			if described.ModifiedAt.Before(cutoff) {
				report.Stale = append(report.Stale, described)
			}

			if described.User == "" {
				report.EmptyUser = append(report.EmptyUser, described)
			}

			if entry.GetPassword() == "" {
				report.EmptyPassword = append(report.EmptyPassword, described)
			}

			seen[join.Empty(described.Title, "\n", described.User)]++
		},
	)
	s.client.Walk(
		func(
			path string,
			_ *gokeepasslib.Group,
			entry *gokeepasslib.Entry,
		) {
			described := describe(path, entry)

			if seen[join.Empty(described.Title, "\n", described.User)] > 1 {
				report.Duplicates = append(report.Duplicates, described)
			}
		},
	)

	return report
}
