package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/search_index/block"

func (s *Service) ConversationBlocks(
	session string,
	kinds []string,
) ([]*block.Block, error) {
	return s.search.ConversationBlocks(session, kinds)
}
