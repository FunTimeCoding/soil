package document

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (s *Service) PushSummary(
	name string,
	body string,
	metadata map[string][]string,
) error {
	return s.pushDocument(
		constant.SummaryCollection,
		constant.SummarySourceType,
		name,
		body,
		metadata,
	)
}
