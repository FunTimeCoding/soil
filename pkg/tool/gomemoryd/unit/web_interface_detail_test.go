package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web/memory_page"
	"strings"
	"testing"
)

func TestHiddenDetailReadsLikeAMissingOne(t *testing.T) {
	s := base.New(t)
	s.Service().WithHiddenTag("private")
	o := save_option.New()
	o.Name = "zulu"
	o.Content = "zulu content"
	o.Type = "feedback"
	o.Tags = []string{"private"}
	identifier, e := s.Store().CreateMemory(o)
	assert.FatalOnError(t, e)
	hidden := body(t, s.Server, memory_page.Path(identifier), false)
	missing := body(t, s.Server, memory_page.Path(identifier+1), false)
	assert.StringContains(
		t,
		fmt.Sprintf("memory not found: %d", identifier),
		hidden,
	)
	assert.StringContains(
		t,
		fmt.Sprintf("memory not found: %d", identifier+1),
		missing,
	)
	assert.True(t, !strings.Contains(hidden, "zulu"))
}
