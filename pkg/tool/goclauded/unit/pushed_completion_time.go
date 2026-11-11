package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"testing"
	"time"
)

func pushedCompletionTime(
	t *testing.T,
	pushed client.PostDocumentJSONRequestBody,
) time.Time {
	t.Helper()
	value := (*pushed.Metadata)[constant.CompletedAtMetadata]
	assert.Count(t, 1, value)
	at, e := time.Parse(time.RFC3339, value[0])
	assert.FatalOnError(t, e)

	return at
}
