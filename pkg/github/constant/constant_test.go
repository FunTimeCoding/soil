package constant

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "X-GitHub-Event", EventHeader)
	assert.String(t, "X-Hub-Signature-256", SignatureHeader)
	assert.String(t, "closed", Closed)
	assert.String(t, "open", Open)
}
