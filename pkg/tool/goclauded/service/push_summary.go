package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (s *Service) pushSummary(
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
