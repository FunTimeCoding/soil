package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/worker"
	"testing"
)

func TestNewestQuerySingleProject(t *testing.T) {
	assert.String(
		t,
		"project in ('ABC') AND statusCategory != Done AND status NOT IN ('Closed') ORDER BY created DESC",
		worker.NewestQuery([]string{"ABC"}, "'Closed'"),
	)
}

func TestNewestQueryQuotesEveryProject(t *testing.T) {
	assert.String(
		t,
		"project in ('ABC','DE F') AND statusCategory != Done AND status NOT IN ('Closed') ORDER BY created DESC",
		worker.NewestQuery([]string{"ABC", "DE F"}, "'Closed'"),
	)
}

func TestNewestQueryCarriesEveryClosedStatus(t *testing.T) {
	assert.String(
		t,
		"project in ('ABC') AND statusCategory != Done AND status NOT IN ('Closed','In Review') ORDER BY created DESC",
		worker.NewestQuery([]string{"ABC"}, "'Closed','In Review'"),
	)
}
