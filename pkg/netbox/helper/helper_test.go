package helper

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestHelper(t *testing.T) {
	assert.String(
		t,
		"http://localhost/a",
		ToWebLink(
			locator.New(
				webConstant.Localhost,
			).Insecure().Base(constant.Interface).Path("/a").String(),
		),
	)
}
