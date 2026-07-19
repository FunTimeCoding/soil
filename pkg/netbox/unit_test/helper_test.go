package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/helper"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestHelper(t *testing.T) {
	assert.String(
		t,
		"http://localhost/a",
		helper.ToWebLink(
			locator.New(
				webConstant.Localhost,
			).Insecure().Base(constant.Interface).Path("/a").String(),
		),
	)
}
