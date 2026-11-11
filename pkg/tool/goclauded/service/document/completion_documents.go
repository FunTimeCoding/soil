package document

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/stored_document"
)

func (s *Service) completionDocuments() []*stored_document.Document {
	var result []*stored_document.Document

	for _, c := range s.store.ListCompletions() {
		slug, metadata, e := s.SessionMetadata(c.SessionIdentifier)
		errors.PanicOnError(e)

		if slug == "" {
			continue
		}

		result = append(
			result,
			stored_document.New(
				c.SessionIdentifier,
				fmt.Sprintf("%s/%d", slug, c.Sequence),
				c.Summary,
				completionMetadata(metadata, &c),
			),
		)
	}

	return result
}
