package document

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (s *Service) ReconcileCompletions() []error {
	s.backfillCompletionSequences()

	return s.reconcileCollection(
		constant.CompletionCollection,
		constant.CompletionSourceType,
		s.completionDocuments(),
	)
}
