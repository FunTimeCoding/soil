package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/directory/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/base"
	"testing"
)

func TestADirectoryOutageHoldsTheNextLoginOff(t *testing.T) {
	d := &countedDirectory{Client: mock_client.New("")}
	d.AddEntry("unique-1", "member", "member@example.test", "secret", true)
	d.Failure = directoryDown()
	b := base.NewWith(t, d)

	for range 2 {
		_, e := b.Service.AuthenticateUser(
			"member",
			"secret",
			constant.SourceDirectory,
		)
		assert.True(t, unavailable.Is(e))
	}

	assert.Integer(t, 1, d.authentications)
}
