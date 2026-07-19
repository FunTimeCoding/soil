package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gmail/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "GMAIL_DIRECTORY", constant.DirectoryEnvironment)
}
