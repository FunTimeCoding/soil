package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/service_tester"
	"testing"
)

func TestServiceRefusesAnUnknownInstance(t *testing.T) {
	_, e := service_tester.Single(mock_client.New()).Projects("charlie", "")
	assert.True(t, not_found.Is(e))
	assert.String(t, "instance not found: charlie", e.Error())
}
