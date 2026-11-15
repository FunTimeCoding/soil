package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/proxmox"
	"github.com/funtimecoding/soil/pkg/proxmox/unit/proxmox_server"
	luthermonson "github.com/luthermonson/go-proxmox"
	"testing"
)

func TestGuestConfigurationThatDoesNotExistIsNotFound(t *testing.T) {
	s := proxmox_server.New(t)
	s.AnswerMessage(
		"/nodes/alfa/qemu/999/status/current",
		"500 Configuration file 'nodes/alfa/qemu-server/999.conf' does not exist",
		"",
	)
	var m map[string]any
	e := proxmox.Get(s.Client(), "/nodes/alfa/qemu/999/status/current", &m)
	assert.True(t, not_found.Is(e))
	assert.StringContains(t, "999.conf' does not exist", e.Error())
}

func TestLockedGuestIsConflict(t *testing.T) {
	s := proxmox_server.New(t)
	s.AnswerMessage(
		"/nodes/alfa/qemu/101/status/start",
		"500 VM is locked (backup)",
		"",
	)
	var m map[string]any
	e := proxmox.Get(s.Client(), "/nodes/alfa/qemu/101/status/start", &m)
	assert.True(t, conflict.Is(e))
}

func TestPermissionCheckIsForbiddenAndStillNotAuthorized(t *testing.T) {
	s := proxmox_server.New(t)
	s.AnswerMessage(
		"/nodes/alfa/qemu/101/status/current",
		"403 Permission check failed (/vms/101, VM.Audit)",
		"",
	)
	var m map[string]any
	e := proxmox.Get(s.Client(), "/nodes/alfa/qemu/101/status/current", &m)
	assert.True(t, forbidden.Is(e))
	assert.True(t, luthermonson.IsNotAuthorized(e))
}

func TestNodeIsFoundWithoutAuditPermission(t *testing.T) {
	s := proxmox_server.New(t)
	s.AnswerMessage(
		"/nodes/alfa/status",
		"403 Permission check failed (/, Sys.Audit)",
		"",
	)
	n, e := s.Client().Node("alfa")
	assert.FatalOnError(t, e)
	assert.String(t, "alfa", n.Name)
}

func TestBadRequestCallerConfiguredIsValidation(t *testing.T) {
	s := proxmox_server.New(t)
	s.AnswerMessage("/nodes/alfa/status", "200 OK", `{"data":{}}`)
	s.AnswerMessage(
		"/nodes/alfa/qemu",
		"400 Parameter verification failed.",
		`{"data":null,"errors":{"memory":"type check ('integer') failed"}}`,
	)
	c := s.Client()
	n, e := c.Node("alfa")
	assert.FatalOnError(t, e)
	_, f := c.CreateMachine(
		n,
		101,
		luthermonson.VirtualMachineOption{Name: "memory", Value: "bravo"},
	)
	assert.True(t, validation.Is(f))
	assert.StringContains(t, "memory: type check", f.Error())
}

func TestBadRequestOnUpdateOrCloneIsValidation(t *testing.T) {
	s := proxmox_server.New(t)
	s.AnswerMessage("/nodes/alfa/status", "200 OK", `{"data":{}}`)
	s.AnswerMessage(
		"/nodes/alfa/qemu/101/status/current",
		"200 OK",
		`{"data":{"vmid":101,"name":"alfa"}}`,
	)
	s.AnswerFor("GET", "/nodes/alfa/qemu/101/config", "200 OK", `{"data":{}}`)
	s.AnswerFor(
		"PUT",
		"/nodes/alfa/qemu/101/config",
		"400 Parameter verification failed.",
		`{"data":null,"errors":{"cores":"value must be at least 1"}}`,
	)
	s.AnswerMessage(
		"/nodes/alfa/qemu/101/clone",
		"400 Parameter verification failed.",
		`{"data":null,"errors":{"name":"invalid format"}}`,
	)
	c := s.Client()
	n, h := c.Node("alfa")
	assert.FatalOnError(t, h)
	m, h := c.Machine(n, 101)
	assert.FatalOnError(t, h)
	f := c.UpdateMachineConfiguration(
		m,
		luthermonson.VirtualMachineOption{Name: "cores", Value: 0},
	)
	assert.True(t, validation.Is(f))
	_, _, g := c.CloneMachine(
		m,
		&luthermonson.VirtualMachineCloneOptions{NewID: 102, Name: "bravo"},
	)
	assert.True(t, validation.Is(g))
}

func TestBadRequestDaemonWroteIsUnexpected(t *testing.T) {
	s := proxmox_server.New(t)
	s.AnswerMessage(
		"/nodes/alfa/qemu",
		"400 Parameter verification failed.",
		`{"data":null,"errors":{"full":"value must be boolean"}}`,
	)
	var m map[string]any
	e := proxmox.Get(s.Client(), "/nodes/alfa/qemu", &m)
	assert.False(t, validation.Is(e))
	assert.True(t, unexpected.Is(e))
}

func TestAnyOtherServerErrorIsUnexpectedWithItsReason(t *testing.T) {
	s := proxmox_server.New(t)
	s.AnswerMessage("/cluster/resources", "500 cfs lock timeout", "")
	var m map[string]any
	e := proxmox.Get(s.Client(), "/cluster/resources", &m)
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "cfs lock timeout", e.Error())
}

func TestVersionReadSaysHypervisorIsReachable(t *testing.T) {
	s := proxmox_server.New(t)
	s.AnswerMessage("/version", "200 OK", `{"data":{"version":"9.2.3"}}`)
	assert.FatalOnError(t, s.Client().Reachable(t.Context()))
}

func TestClosedPortIsUnreachable(t *testing.T) {
	c := proxmox.New(
		"127.0.0.1",
		proxmox.WithToken("alfa@pve!test", "secret"),
		proxmox.WithPort(1),
		proxmox.WithUntrusted(),
	)
	defer c.Close()
	assert.True(t, unreachable.Is(c.Reachable(t.Context())))
}
