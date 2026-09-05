package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/debian"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"testing"
)

func TestRenderControl(t *testing.T) {
	assert.String(
		t,
		`Package: goexample
Version: 1.0.0
Architecture: amd64
Maintainer: John Doe <john.doe@example.org>
Description: Short stub description.
 Long stub description.
`,
		debian.RenderControl(
			"goexample",
			constant.AMD64,
			library.DefaultVersion,
			"John Doe",
			"john.doe@example.org",
		),
	)
}
