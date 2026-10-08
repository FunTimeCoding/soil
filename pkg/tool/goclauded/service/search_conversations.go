package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index/conversation"
)

func (s *Service) SearchConversations(
	query string,
	kinds []string,
	limit int,
) ([]*conversation.Conversation, error) {
	if limit <= 0 {
		limit = constant.SearchConversationLimit
	}

	result, e := s.search.Search(
		query,
		kinds,
		min(limit, constant.SearchConversationMaximum),
	)

	if e != nil {
		return nil, e
	}

	for _, c := range result {
		c.Name = s.sessionName(c.Session)
	}

	return result, nil
}
