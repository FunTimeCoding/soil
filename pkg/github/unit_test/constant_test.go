package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/github/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "X-GitHub-Event", constant.EventHeader)
	assert.String(t, "X-Hub-Signature-256", constant.SignatureHeader)
	assert.String(t, "closed", constant.Closed)
	assert.String(t, "open", constant.Open)
}
