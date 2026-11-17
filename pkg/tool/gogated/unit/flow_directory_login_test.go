package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/directory/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/directory_tester"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/flow_tester"
	"net/http"
	"testing"
)

func TestLoginWhileDirectoryIsAwaySaysSo(t *testing.T) {
	d := mock_client.New("")
	d.AddEntry("unique-1", "member", "member@example.test", "secret", true)
	o := flow_tester.New(t, base.WithDirectory(d))
	reg := o.Register([]string{"http://localhost/callback"})
	d.Failure = directory_tester.Down()
	status, body := o.Login(
		reg.ClientIdentifier,
		"http://localhost/callback",
		"member",
		"secret",
		constant.SourceDirectory,
	)
	assert.Integer(t, http.StatusServiceUnavailable, status)
	assert.StringContains(t, "can&#39;t be reached", body)
	assert.StringNotContains(t, "rejected those credentials", body)
}

func TestWrongDirectoryPasswordIsStillRejected(t *testing.T) {
	d := mock_client.New("")
	d.AddEntry("unique-1", "member", "member@example.test", "secret", true)
	o := flow_tester.New(t, base.WithDirectory(d))
	reg := o.Register([]string{"http://localhost/callback"})
	status, body := o.Login(
		reg.ClientIdentifier,
		"http://localhost/callback",
		"member",
		"wrong",
		constant.SourceDirectory,
	)
	assert.Integer(t, http.StatusOK, status)
	assert.StringContains(t, "rejected those credentials", body)
}

func TestLocalLoginWorksWhileDirectoryIsAway(t *testing.T) {
	d := mock_client.New("")
	d.Failure = directory_tester.Down()
	o := flow_tester.New(t, base.WithDirectory(d))
	reg := o.Register([]string{"http://localhost/callback"})
	status, _ := o.Login(
		reg.ClientIdentifier,
		"http://localhost/callback",
		constant.FixtureMail,
		constant.FixturePassword,
		constant.SourceLocal,
	)
	assert.Integer(t, http.StatusSeeOther, status)
}
