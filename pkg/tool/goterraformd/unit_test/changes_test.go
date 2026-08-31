package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/runner"
	"testing"
)

func TestParseChanges(t *testing.T) {
	output := `module.proxmox.proxmox_virtual_environment_vm.alpha: Refreshing state... [id=101]

Terraform will perform the following actions:

module.proxmox.proxmox_virtual_environment_vm.alpha: Modifying... [id=101]
module.proxmox.proxmox_virtual_environment_vm.alpha: Still modifying... [id=101, 00m10s elapsed]
module.proxmox.proxmox_virtual_environment_vm.alpha: Modifications complete after 15s [id=101]
module.postgres.postgresql_role.charlie: Creating...
module.postgres.postgresql_role.charlie: Creation complete after 1s [id=charlie]
module.proxmox.proxmox_virtual_environment_vm.bravo: Destroying... [id=102]
module.proxmox.proxmox_virtual_environment_vm.bravo: Destruction complete after 3s

Apply complete! Resources: 1 added, 1 changed, 1 destroyed.`
	assert.Any(
		t,
		[]string{
			"module.proxmox.proxmox_virtual_environment_vm.alpha",
			"module.postgres.postgresql_role.charlie",
			"module.proxmox.proxmox_virtual_environment_vm.bravo",
		},
		runner.ParseChanges(output),
	)
}

func TestParseChangesNone(t *testing.T) {
	output := `module.proxmox.proxmox_virtual_environment_vm.alpha: Refreshing state... [id=101]

No changes. Your infrastructure matches the configuration.

Apply complete! Resources: 0 added, 0 changed, 0 destroyed.`
	assert.Count(t, 0, runner.ParseChanges(output))
}

func TestChangesError(t *testing.T) {
	record := &store.Run{} // goanalyze:ignore struct_literal
	record.Output = "module.a.b: Creation complete after 1s"
	assert.Count(t, 0, runner.Changes(record))
}

func TestChangesWrongType(t *testing.T) {
	assert.Count(t, 0, runner.Changes("not a run"))
}
