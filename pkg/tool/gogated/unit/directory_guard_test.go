package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/counted_directory"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/directory_tester"
	"testing"
)

func TestDirectoryOutageHoldsNextLoginOff(t *testing.T) {
	d := counted_directory.New()
	d.AddEntry("unique-1", "member", "member@example.test", "secret", true)
	d.Failure = directory_tester.Down()
	b := base.NewWith(t, d)

	for range 2 {
		_, e := b.Service.AuthenticateUser(
			"member",
			"secret",
			constant.SourceDirectory,
		)
		assert.True(t, unavailable.Is(e))
	}

	assert.Integer(t, 1, d.Authentications)
}
