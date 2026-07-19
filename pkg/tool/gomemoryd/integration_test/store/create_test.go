package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration_test/store_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"testing"
)

func TestCreateAndGetMemory(t *testing.T) {
	s := store_tester.New(t)
	o := save_option.New()
	o.Content = "Retry failed requests with exponential backoff."
	o.Description = "retry with backoff; each caller gets its own policy"
	o.Type = "feedback"
	o.Tags = []string{"always", "http-client"}
	identifier := s.CreateMemory(o)
	m := s.GetMemory(identifier)
	assert.String(
		t,
		"Retry failed requests with exponential backoff.",
		m.Content,
	)
	assert.Count(t, 2, m.Tags)
}
