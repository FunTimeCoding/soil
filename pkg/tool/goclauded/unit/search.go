package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index/conversation"
	"testing"
)

func search(
	t *testing.T,
	x *search_index.Index,
	query string,
	kinds ...string,
) []*conversation.Conversation {
	t.Helper()
	result, e := x.Search(query, kinds, 20)
	assert.FatalOnError(t, e)

	return result
}
