package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/memory_indexer_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face/search_option"
	"testing"
)

func TestASearchGoquerydCannotRunIsUnavailableWithItsReason(t *testing.T) {
	s := memory_indexer_tester.New(
		t,
		memory_indexer_tester.UnavailableFixture(),
		memory_indexer_tester.EmptyFixture(),
	)
	_, e := s.Search(search_option.New("alfa", constant.DefaultCollection, 10))
	assert.True(t, unavailable.Is(e))
	assert.String(
		t,
		"search: embedder unreachable since 14:02: no route",
		e.Error(),
	)
}
