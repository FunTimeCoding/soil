package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/nextcloud/helper"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestBase(t *testing.T) {
	assert.String(t, "https://example.org", helper.Base(constant.Example))
}
