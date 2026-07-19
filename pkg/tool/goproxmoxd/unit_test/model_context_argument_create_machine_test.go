package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument/create_machine"
	"github.com/luthermonson/go-proxmox"
	"testing"
)

func findOption(
	options []proxmox.VirtualMachineOption,
	name string,
) (any, bool) {
	for _, o := range options {
		if o.Name == name {
			return o.Value, true
		}
	}

	return nil, false
}

func requireOption(
	t *testing.T,
	options []proxmox.VirtualMachineOption,
	name string,
) any {
	t.Helper()
	v, okay := findOption(options, name)

	if !okay {
		t.Fatalf("option %q not found", name)
	}

	return v
}

func TestBuildOptionsDefaults(t *testing.T) {
	m := create_machine.New()
	m.Name = "test-vm"
	options := m.BuildOptions()
	assert.String(t, "test-vm", requireOption(t, options, "name").(string))
	assert.Integer(t, 2, requireOption(t, options, "cores").(int))
	assert.Integer(t, 1, requireOption(t, options, "sockets").(int))
	assert.Integer(t, 2048, requireOption(t, options, "memory").(int))
	assert.String(
		t,
		"virtio-scsi-pci",
		requireOption(t, options, "scsihw").(string),
	)
	assert.String(
		t,
		"local-lvm:32,aio=io_uring,backup=1,cache=none,discard=on,iothread=1,replicate=1",
		requireOption(t, options, "virtio0").(string),
	)
	assert.String(
		t,
		"order=virtio0;net0",
		requireOption(t, options, "boot").(string),
	)
	assert.String(
		t,
		"virtio,bridge=vmbr0",
		requireOption(t, options, "net0").(string),
	)
	assert.Integer(t, 0, requireOption(t, options, "balloon").(int))
	assert.Integer(t, 1, requireOption(t, options, "agent").(int))
	assert.String(t, "host", requireOption(t, options, "cpu").(string))
	_, hasIDE := findOption(options, "ide2")
	assert.Boolean(t, false, hasIDE)
	_, hasCI := findOption(options, "ipconfig0")
	assert.Boolean(t, false, hasCI)
	_, hasOnBoot := findOption(options, "onboot")
	assert.Boolean(t, false, hasOnBoot)
}

func TestBuildOptionsOnBoot(t *testing.T) {
	m := create_machine.New()
	m.Name = "boot-vm"
	m.OnBoot = new(true)
	options := m.BuildOptions()
	assert.Integer(t, 1, requireOption(t, options, "onboot").(int))
}

func TestBuildOptionsOnBootDisabled(t *testing.T) {
	m := create_machine.New()
	m.Name = "no-boot-vm"
	m.OnBoot = new(false)
	options := m.BuildOptions()
	assert.Integer(t, 0, requireOption(t, options, "onboot").(int))
}

func TestBuildOptionsAgentDisabled(t *testing.T) {
	m := create_machine.New()
	m.Name = "no-agent"
	m.Agent = new(false)
	options := m.BuildOptions()
	_, hasAgent := findOption(options, "agent")
	assert.Boolean(t, false, hasAgent)
}

func TestBuildOptionsCustomValues(t *testing.T) {
	m := create_machine.New()
	m.Name = "custom"
	m.Cores = 8
	m.Memory = 16384
	m.Bridge = "vmbr0"
	m.OSType = "l26"
	m.Tags = "prod;web"
	options := m.BuildOptions()
	assert.Integer(t, 8, requireOption(t, options, "cores").(int))
	assert.Integer(t, 16384, requireOption(t, options, "memory").(int))
	assert.String(
		t,
		"virtio,bridge=vmbr0",
		requireOption(t, options, "net0").(string),
	)
	assert.String(t, "l26", requireOption(t, options, "ostype").(string))
	assert.String(t, "prod;web", requireOption(t, options, "tags").(string))
}

func TestBuildOptionsDiskImport(t *testing.T) {
	m := create_machine.New()
	m.Name = "import-vm"
	m.DiskImport = "local:import/debian-13-generic-amd64.qcow2"
	options := m.BuildOptions()
	assert.String(
		t,
		"local-lvm:0,import-from=local:import/debian-13-generic-amd64.qcow2,aio=io_uring,backup=1,cache=none,discard=on,iothread=1,replicate=1",
		requireOption(t, options, "virtio0").(string),
	)
}

func TestBuildOptionsDiskImportCustomStorage(t *testing.T) {
	m := create_machine.New()
	m.Name = "import-vm"
	m.DiskImport = "local:import/debian-13.qcow2"
	m.DiskStorage = "ceph-pool"
	options := m.BuildOptions()
	assert.String(
		t,
		"ceph-pool:0,import-from=local:import/debian-13.qcow2,aio=io_uring,backup=1,cache=none,discard=on,iothread=1,replicate=1",
		requireOption(t, options, "virtio0").(string),
	)
}

func TestBuildOptionsCDROM(t *testing.T) {
	m := create_machine.New()
	m.Name = "iso-vm"
	m.CDROM = "local:iso/debian-13.iso"
	options := m.BuildOptions()
	assert.String(
		t,
		"local:iso/debian-13.iso,media=cdrom",
		requireOption(t, options, "ide2").(string),
	)
}

func TestBuildOptionsCloudInitTakesIDE2OverCDROM(t *testing.T) {
	m := create_machine.New()
	m.Name = "ci-with-cdrom"
	m.CDROM = "local:iso/debian-13.iso"
	m.CIUser = "admin"
	options := m.BuildOptions()
	assert.String(
		t,
		"local-lvm:cloudinit",
		requireOption(t, options, "ide2").(string),
	)
}

func TestBuildOptionsCloudInit(t *testing.T) {
	m := create_machine.New()
	m.Name = "ci-vm"
	m.CIUser = "admin"
	options := m.BuildOptions()
	assert.String(t, "admin", requireOption(t, options, "ciuser").(string))
	assert.String(
		t,
		"ip=dhcp",
		requireOption(t, options, "ipconfig0").(string),
	)
	assert.String(
		t,
		"local-lvm:cloudinit",
		requireOption(t, options, "ide2").(string),
	)
}

func TestBuildOptionsCloudInitFull(t *testing.T) {
	m := create_machine.New()
	m.Name = "ci-vm"
	m.CIUser = "deploy"
	m.CIPassword = "secret"
	m.SSHKeys = "ssh-ed25519 AAAA key1\nssh-ed25519 BBBB key2"
	m.IPConfiguration = "ip=10.0.0.5/24,gw=10.0.0.1"
	options := m.BuildOptions()
	assert.String(t, "deploy", requireOption(t, options, "ciuser").(string))
	assert.String(
		t,
		"secret",
		requireOption(t, options, "cipassword").(string),
	)
	assert.String(
		t,
		"ip=10.0.0.5/24,gw=10.0.0.1",
		requireOption(t, options, "ipconfig0").(string),
	)
	_, hasKeys := findOption(options, "sshkeys")
	assert.Boolean(t, true, hasKeys)
	assert.String(
		t,
		"local-lvm:cloudinit",
		requireOption(t, options, "ide2").(string),
	)
}

func TestBuildOptionsCloudInitSSHKeysOnly(t *testing.T) {
	m := create_machine.New()
	m.Name = "ssh-vm"
	m.SSHKeys = "ssh-ed25519 AAAA key1"
	options := m.BuildOptions()
	assert.String(
		t,
		"ip=dhcp",
		requireOption(t, options, "ipconfig0").(string),
	)
	assert.String(
		t,
		"local-lvm:cloudinit",
		requireOption(t, options, "ide2").(string),
	)
}

func TestBuildOptionsCustomCPUType(t *testing.T) {
	m := create_machine.New()
	m.Name = "cpu-vm"
	m.CPUType = "x86-64-v2-AES"
	options := m.BuildOptions()
	assert.String(
		t,
		"x86-64-v2-AES",
		requireOption(t, options, "cpu").(string),
	)
}

func TestBuildOptionsSearchDomain(t *testing.T) {
	m := create_machine.New()
	m.Name = "dns-vm"
	m.CIUser = "admin"
	m.SearchDomain = "local"
	options := m.BuildOptions()
	assert.String(
		t,
		"local",
		requireOption(t, options, "searchdomain").(string),
	)
}

func TestBuildOptionsExtras(t *testing.T) {
	m := create_machine.New()
	m.Name = "extras-vm"
	m.Extras = "serial0=socket,vga=serial0"
	options := m.BuildOptions()
	assert.String(
		t,
		"socket",
		requireOption(t, options, "serial0").(string),
	)
	assert.String(
		t,
		"serial0",
		requireOption(t, options, "vga").(string),
	)
}

func TestBuildOptionsCustomDiskSize(t *testing.T) {
	m := create_machine.New()
	m.Name = "big-vm"
	m.DiskSize = 100
	options := m.BuildOptions()
	assert.String(
		t,
		"local-lvm:100,aio=io_uring,backup=1,cache=none,discard=on,iothread=1,replicate=1",
		requireOption(t, options, "virtio0").(string),
	)
}
