package search_index

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index/conversation"
	"strings"
)

func (x *Index) Search(
	query string,
	kinds []string,
	limit int,
) ([]*conversation.Conversation, error) {
	terms := strings.Fields(query)

	if len(terms) == 0 {
		return nil, nil
	}

	return x.matches(terms, kindsOrDefault(kinds), limit)
}
