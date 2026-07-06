package web

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestGetList(t *testing.T) {
	assert.Any(
		t,
		[]string{"1", "2", "3"},
		GetList(
			NewGet(
				locator.New(
					constant.Localhost,
				).Insecure().Set("a", "1,2,3").String(),
			),
			"a",
		),
	)
}
