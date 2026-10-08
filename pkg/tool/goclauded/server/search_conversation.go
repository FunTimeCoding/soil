package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index/conversation"
)

func searchConversation(
	c *conversation.Conversation,
) server.SearchConversation {
	hits := []server.SearchHit{}

	for _, h := range c.Hits {
		hits = append(
			hits,
			server.SearchHit{
				Identifier: h.Identifier,
				Role:       h.Role,
				Kind:       h.Kind,
				At:         h.At,
				Snippet:    h.Snippet,
			},
		)
	}

	return server.SearchConversation{
		Session: c.Session,
		Name:    c.Name,
		Latest:  c.Latest,
		Count:   c.Count,
		Hits:    hits,
	}
}
