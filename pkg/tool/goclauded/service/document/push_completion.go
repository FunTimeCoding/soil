package document

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/completion"
)

func (s *Service) PushCompletion(
	slug string,
	c *completion.Completion,
	body string,
	metadata map[string][]string,
) error {
	return s.pushDocument(
		constant.CompletionCollection,
		constant.CompletionSourceType,
		fmt.Sprintf("%s/%d", slug, c.Sequence),
		body,
		completionMetadata(metadata, c),
	)
}
