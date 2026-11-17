package service_tester

import (
	confluenceMock "github.com/funtimecoding/soil/pkg/atlassian/confluence/mock_client"
	jiraMock "github.com/funtimecoding/soil/pkg/atlassian/jira/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"testing"
)

func New(
	t *testing.T,
	o ...Option,
) *Tester {
	t.Helper()
	result := &Tester{
		Jira:       jiraMock.New(t),
		Confluence: confluenceMock.New(t),
		Reporter:   memory.New(),
		notifier:   notifier.New(),
		t:          t,
	}

	for _, f := range o {
		f(result)
	}

	return result
}
