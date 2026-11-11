package document

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/stored_document"
)

func (s *Service) summaryDocuments() []*stored_document.Document {
	var result []*stored_document.Document

	for _, m := range s.store.ListSummaries() {
		name, metadata, e := s.SessionMetadata(m.SessionIdentifier)
		errors.PanicOnError(e)

		if name == "" {
			continue
		}

		result = append(
			result,
			stored_document.New(m.SessionIdentifier, name, m.Body, metadata),
		)
	}

	return result
}
