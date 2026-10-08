package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index/block"
)

func (s *Service) ReadConversation(
	session string,
	around string,
	count int,
) ([]*block.Block, error) {
	if count <= 0 {
		count = constant.ReadWindowDefault
	}

	result, e := s.search.Window(
		session,
		around,
		min(count, constant.ReadWindowMaximum),
	)

	if e != nil {
		return nil, e
	}

	for _, b := range result {
		text := []rune(b.Text)

		if len(text) > constant.BlockTextLimit {
			b.Text = fmt.Sprintf(
				"%s\n(%d more characters)",
				string(text[:constant.BlockTextLimit]),
				len(text)-constant.BlockTextLimit,
			)
		}
	}

	return result, nil
}
