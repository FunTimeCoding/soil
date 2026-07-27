package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (s *Service) pushCompletion(
	slug string,
	sequence int,
	body string,
	metadata map[string]string,
) error {
	return s.completionIndexer.Push(
		constant.CompletionCollection,
		fmt.Sprintf("%s/%d", slug, sequence),
		body,
		metadata,
	)
}
