package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/provision/client/mock_client"
	"github.com/funtimecoding/soil/pkg/provision/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/security_tester"
	"testing"
)

func TestFeedStoresEachProbedHost(t *testing.T) {
	o := security_tester.New(t)
	a := mock_client.New(t, constant.AnsiblePeer)
	a.SetHostResults(
		[]client.HostCommandResult{
			{
				Host:   "alfa",
				Status: "success",
				Stdout: security_tester.ProbeFixture("codename"),
			},
			{Host: "bravo", Status: "unreachable", Stderr: "bravo down"},
			{Host: "charlie", Status: "success", Stdout: "charlie garbled"},
		},
	)
	errors := o.Feed(a).Run(t.Context())
	assert.Count(t, 1, errors)
	assert.StringContains(t, "probe charlie", errors[0].Error())
	alfa, e := o.Store().SystemByName("alfa")
	assert.FatalOnError(t, e)
	assert.String(t, "trixie", alfa.Release)
	_, f := o.Store().SystemByName("bravo")
	assert.True(t, not_found.Is(f))
	_, g := o.Store().SystemByName("charlie")
	assert.True(t, not_found.Is(g))
	assert.Any(
		t,
		[]client.HostCommandRequest{
			{Command: "security_probe", Source: "gosecurityd"},
		},
		a.HostCommands(),
	)
}

func TestFeedReturnsAnsibleRefusal(t *testing.T) {
	o := security_tester.New(t)
	a := mock_client.New(t, constant.AnsiblePeer)
	a.Fail(
		a.PostHostCommand,
		not_found.New("host command", constant.ProbeHostCommand),
	)
	errors := o.Feed(a).Run(t.Context())
	assert.Count(t, 1, errors)
	assert.True(t, not_found.Is(errors[0]))
}
