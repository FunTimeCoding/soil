package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/directory/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/tester"
	"net/http"
	"testing"
)

func TestALoginWhileTheDirectoryIsAwaySaysSo(t *testing.T) {
	d := mock_client.New("")
	d.AddEntry("unique-1", "member", "member@example.test", "secret", true)
	o := tester.NewWith(t, d)
	reg := o.Register(t, []string{"http://localhost/callback"})
	d.Failure = directoryDown()
	status, body := o.Login(
		t,
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

func TestAWrongDirectoryPasswordIsStillRejected(t *testing.T) {
	d := mock_client.New("")
	d.AddEntry("unique-1", "member", "member@example.test", "secret", true)
	o := tester.NewWith(t, d)
	reg := o.Register(t, []string{"http://localhost/callback"})
	status, body := o.Login(
		t,
		reg.ClientIdentifier,
		"http://localhost/callback",
		"member",
		"wrong",
		constant.SourceDirectory,
	)
	assert.Integer(t, http.StatusOK, status)
	assert.StringContains(t, "rejected those credentials", body)
}

func TestALocalLoginWorksWhileTheDirectoryIsAway(t *testing.T) {
	d := mock_client.New("")
	d.Failure = directoryDown()
	o := tester.NewWith(t, d)
	reg := o.Register(t, []string{"http://localhost/callback"})
	status, _ := o.Login(
		t,
		reg.ClientIdentifier,
		"http://localhost/callback",
		constant.FixtureMail,
		constant.FixturePassword,
		constant.SourceLocal,
	)
	assert.Integer(t, http.StatusSeeOther, status)
}
