package document

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (s *Service) ReconcileSummaries() []error {
	return s.reconcileCollection(
		constant.SummaryCollection,
		constant.SummarySourceType,
		s.summaryDocuments(),
	)
}
