package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/provision/types/host_result"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/unit/service_tester"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestHostCommandRunsModuleOnGroupMember(t *testing.T) {
	o := service_tester.New(t, nil)
	o.Service()
	o.WaitForRun()
	o.DefineHostCommand()
	result, e := o.Service().HostCommand("alfa", "charlie", "delta")
	assert.FatalOnError(t, e)
	assert.String(t, "success", result.Record.Status)
	assert.String(t, "alfa", result.Record.Scope)
	assert.String(t, "delta", result.Record.TriggerSource)
	assert.Any(
		t,
		[]*host_result.Result{
			host_result.New("charlie", "success", 0, "echo restarted", ""),
		},
		result.Hosts,
	)
	arguments, okay := o.Captured(constant.FixtureAdHoc)
	assert.True(t, okay)
	assert.String(
		t,
		"bravo:&charlie\n-m\nansible.builtin.systemd_service\n-a\nname=echo state=restarted\n",
		arguments,
	)
	callback, found := o.Captured(constant.FixtureCallback)
	assert.True(t, found)
	assert.String(t, "ansible.posix.json 1\n", callback)
	stored, f := o.Service().Find(result.Record.Identifier)
	errors.PanicOnError(f)
	assert.String(t, "success", stored.Status)
}

func TestHostCommandWithoutHostRunsWholeGroup(t *testing.T) {
	o := service_tester.New(t, nil)
	o.Service()
	o.WaitForRun()
	o.DefineHostCommand()
	o.PlaceOutside()
	result, e := o.Service().HostCommand("alfa", "", "delta")
	assert.FatalOnError(t, e)
	assert.String(t, "error", result.Record.Status)
	assert.Any(
		t,
		[]*host_result.Result{
			host_result.New("charlie", "success", 0, "echo restarted", ""),
			host_result.New("delta", "unreachable", 0, "", "delta down"),
		},
		result.Hosts,
	)
	arguments, okay := o.Captured(constant.FixtureAdHoc)
	assert.True(t, okay)
	assert.String(
		t,
		"bravo\n-m\nansible.builtin.systemd_service\n-a\nname=echo state=restarted\n",
		arguments,
	)
}

func TestHostCommandRunsAgainstNetboxInventory(t *testing.T) {
	o := service_tester.New(
		t,
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				_ *http.Request,
			) {
				_, e := w.Write([]byte(constant.FixtureMarkup))
				errors.PanicOnError(e)
			},
		),
	)
	o.Service()
	o.WaitForRun()
	o.DefineHostCommand()
	result, e := o.Service().HostCommand("alfa", "charlie", "delta")
	assert.FatalOnError(t, e)
	assert.String(t, "success", result.Record.Status)
	arguments, okay := o.Captured(constant.FixtureAdHoc)
	assert.True(t, okay)
	a := strings.Split(strings.TrimSpace(arguments), "\n")
	assert.Count(t, 9, a)
	assert.String(t, "/etc/ansible/hosts", a[1])
	assert.True(t, strings.HasSuffix(a[3], "/netbox.yaml"))
	assert.String(t, "bravo:&charlie", a[4])
	_, f := os.Stat(a[3])
	assert.True(t, os.IsNotExist(f))
}

func TestHostCommandFailsWhenNetboxInventoryFails(t *testing.T) {
	o := service_tester.New(
		t,
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				_ *http.Request,
			) {
				w.WriteHeader(http.StatusInternalServerError)
			},
		),
	)
	o.Service()
	o.WaitForRun()
	o.DefineHostCommand()
	_, e := o.Service().HostCommand("alfa", "charlie", "delta")
	assert.Error(t, e)
	_, ran := o.Captured(constant.FixtureAdHoc)
	assert.False(t, ran)
}

func TestHostCommandOutsideGroupInvalid(t *testing.T) {
	o := service_tester.New(t, nil)
	o.Service()
	o.WaitForRun()
	o.DefineHostCommand()
	o.PlaceOutside()
	_, e := o.Service().HostCommand("alfa", "charlie", "delta")
	assert.True(t, validation.Is(e))
	_, okay := o.Captured(constant.FixtureAdHoc)
	assert.False(t, okay)
	recent, f := o.Service().Recent(10)
	errors.PanicOnError(f)
	assert.Count(t, 1, recent)
}

func TestUnknownHostCommandNotFound(t *testing.T) {
	o := service_tester.New(t, nil)
	o.Service()
	o.WaitForRun()
	o.DefineHostCommand()
	_, e := o.Service().HostCommand("echo", "charlie", "delta")
	assert.True(t, not_found.Is(e))
}

func TestHostCommandWithoutDefinitionsNotFound(t *testing.T) {
	o := service_tester.New(t, nil)
	o.Service()
	o.WaitForRun()
	_, e := o.Service().HostCommand("alfa", "charlie", "delta")
	assert.True(t, not_found.Is(e))
}

func TestFailedHostCommandRecordedAsError(t *testing.T) {
	o := service_tester.New(t, nil)
	o.Service()
	o.WaitForRun()
	o.DefineHostCommand()
	o.FailHostCommand()
	result, e := o.Service().HostCommand("alfa", "charlie", "delta")
	assert.FatalOnError(t, e)
	assert.String(t, "error", result.Record.Status)
	assert.Any(
		t,
		[]*host_result.Result{
			host_result.New("charlie", "error", 2, "", "echo failed"),
		},
		result.Hosts,
	)
}

func TestHostCommandAfterStopUnavailable(t *testing.T) {
	o := service_tester.New(t, nil)
	o.Service()
	o.WaitForRun()
	o.DefineHostCommand()
	o.Stop()
	_, e := o.Service().HostCommand("alfa", "charlie", "delta")
	assert.True(t, unavailable.Is(e))
}
