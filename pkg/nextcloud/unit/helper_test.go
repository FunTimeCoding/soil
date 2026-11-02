package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/nextcloud/helper"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestBase(t *testing.T) {
	assert.String(t, "https://example.org", helper.Base(webConstant.Example))
}
