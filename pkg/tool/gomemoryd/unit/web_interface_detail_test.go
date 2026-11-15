package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web/memory_page"
	"strings"
	"testing"
	"time"
)

func TestHiddenDetailReadsLikeMissingOne(t *testing.T) {
	s := base.NewHidden(t, "private")
	o := save_option.New()
	o.Name = "zulu"
	o.Content = "zulu content"
	o.Type = "feedback"
	o.Tags = []string{"private"}
	identifier, e := s.Store().CreateMemory(o, time.Now())
	assert.FatalOnError(t, e)
	hidden := s.Body(t, memory_page.Path(identifier))
	missing := s.Body(t, memory_page.Path(identifier+1))
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
