package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument/create_machine"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/service_tester"
	"testing"
)

func TestBuildOptionsDefaults(t *testing.T) {
	m := create_machine.New()
	m.Name = "test-vm"
	options := m.BuildOptions()
	assert.String(
		t,
		"test-vm",
		service_tester.RequireOption(t, options, constant.NameOption).(string),
	)
	assert.Integer(
		t,
		2,
		service_tester.RequireOption(t, options, constant.CoresOption).(int),
	)
	assert.Integer(
		t,
		1,
		service_tester.RequireOption(t, options, constant.SocketsOption).(int),
	)
	assert.Integer(
		t,
		2048,
		service_tester.RequireOption(t, options, constant.MemoryOption).(int),
	)
	assert.String(
		t,
		"virtio-scsi-pci",
		service_tester.RequireOption(t, options, constant.DiskControllerOption).(string),
	)
	assert.String(
		t,
		"local-lvm:32,aio=io_uring,backup=1,cache=none,discard=on,iothread=1,replicate=1",
		service_tester.RequireOption(t, options, constant.PrimaryDiskOption).(string),
	)
	assert.String(
		t,
		"order=virtio0;net0",
		service_tester.RequireOption(t, options, constant.BootOption).(string),
	)
	assert.String(
		t,
		"virtio,bridge=vmbr0",
		service_tester.RequireOption(t, options, constant.PrimaryNetworkOption).(string),
	)
	assert.Integer(
		t,
		0,
		service_tester.RequireOption(t, options, constant.BalloonOption).(int),
	)
	assert.Integer(
		t,
		1,
		service_tester.RequireOption(t, options, constant.AgentOption).(int),
	)
	assert.String(
		t,
		"host",
		service_tester.RequireOption(t, options, constant.ProcessorOption).(string),
	)
	_, hasIDE := service_tester.FindOption(
		options,
		constant.RemovableDriveOption,
	)
	assert.Boolean(t, false, hasIDE)
	_, hasCI := service_tester.FindOption(
		options,
		constant.InternetConfigurationOption,
	)
	assert.Boolean(t, false, hasCI)
	_, hasOnBoot := service_tester.FindOption(options, constant.OnBootOption)
	assert.Boolean(t, false, hasOnBoot)
}

func TestBuildOptionsOnBoot(t *testing.T) {
	m := create_machine.New()
	m.Name = "boot-vm"
	m.OnBoot = new(true)
	options := m.BuildOptions()
	assert.Integer(
		t,
		1,
		service_tester.RequireOption(t, options, constant.OnBootOption).(int),
	)
}

func TestBuildOptionsOnBootDisabled(t *testing.T) {
	m := create_machine.New()
	m.Name = "no-boot-vm"
	m.OnBoot = new(false)
	options := m.BuildOptions()
	assert.Integer(
		t,
		0,
		service_tester.RequireOption(t, options, constant.OnBootOption).(int),
	)
}

func TestBuildOptionsAgentDisabled(t *testing.T) {
	m := create_machine.New()
	m.Name = "no-agent"
	m.Agent = new(false)
	options := m.BuildOptions()
	_, hasAgent := service_tester.FindOption(options, constant.AgentOption)
	assert.Boolean(t, false, hasAgent)
}

func TestBuildOptionsCustomValues(t *testing.T) {
	m := create_machine.New()
	m.Name = "custom"
	m.Cores = 8
	m.Memory = 16384
	m.Bridge = "vmbr0"
	m.OperatingSystemType = "l26"
	m.Tags = "prod;web"
	options := m.BuildOptions()
	assert.Integer(
		t,
		8,
		service_tester.RequireOption(t, options, constant.CoresOption).(int),
	)
	assert.Integer(
		t,
		16384,
		service_tester.RequireOption(t, options, constant.MemoryOption).(int),
	)
	assert.String(
		t,
		"virtio,bridge=vmbr0",
		service_tester.RequireOption(t, options, constant.PrimaryNetworkOption).(string),
	)
	assert.String(
		t,
		"l26",
		service_tester.RequireOption(t, options, constant.OperatingSystemOption).(string),
	)
	assert.String(
		t,
		"prod;web",
		service_tester.RequireOption(t, options, constant.TagsOption).(string),
	)
}

func TestBuildOptionsDiskImport(t *testing.T) {
	m := create_machine.New()
	m.Name = "import-vm"
	m.DiskImport = "local:import/debian-13-generic-amd64.qcow2"
	options := m.BuildOptions()
	assert.String(
		t,
		"local-lvm:0,import-from=local:import/debian-13-generic-amd64.qcow2,aio=io_uring,backup=1,cache=none,discard=on,iothread=1,replicate=1",
		service_tester.RequireOption(t, options, constant.PrimaryDiskOption).(string),
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
		service_tester.RequireOption(t, options, constant.PrimaryDiskOption).(string),
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
		service_tester.RequireOption(t, options, constant.RemovableDriveOption).(string),
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
		service_tester.RequireOption(t, options, constant.RemovableDriveOption).(string),
	)
}

func TestBuildOptionsCloudInit(t *testing.T) {
	m := create_machine.New()
	m.Name = "ci-vm"
	m.CIUser = "admin"
	options := m.BuildOptions()
	assert.String(
		t,
		"admin",
		service_tester.RequireOption(t, options, constant.CloudInitUserOption).(string),
	)
	assert.String(
		t,
		"ip=dhcp",
		service_tester.RequireOption(
			t,
			options,
			constant.InternetConfigurationOption,
		).(string),
	)
	assert.String(
		t,
		"local-lvm:cloudinit",
		service_tester.RequireOption(t, options, constant.RemovableDriveOption).(string),
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
	assert.String(
		t,
		"deploy",
		service_tester.RequireOption(t, options, constant.CloudInitUserOption).(string),
	)
	assert.String(
		t,
		"secret",
		service_tester.RequireOption(t, options, constant.CloudInitPasswordOption).(string),
	)
	assert.String(
		t,
		"ip=10.0.0.5/24,gw=10.0.0.1",
		service_tester.RequireOption(
			t,
			options,
			constant.InternetConfigurationOption,
		).(string),
	)
	_, hasKeys := service_tester.FindOption(
		options,
		constant.SecureShellKeysOption,
	)
	assert.Boolean(t, true, hasKeys)
	assert.String(
		t,
		"local-lvm:cloudinit",
		service_tester.RequireOption(t, options, constant.RemovableDriveOption).(string),
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
		service_tester.RequireOption(
			t,
			options,
			constant.InternetConfigurationOption,
		).(string),
	)
	assert.String(
		t,
		"local-lvm:cloudinit",
		service_tester.RequireOption(t, options, constant.RemovableDriveOption).(string),
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
		service_tester.RequireOption(t, options, constant.ProcessorOption).(string),
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
		service_tester.RequireOption(t, options, constant.SearchDomainOption).(string),
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
		service_tester.RequireOption(t, options, "serial0").(string),
	)
	assert.String(
		t,
		"serial0",
		service_tester.RequireOption(t, options, "vga").(string),
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
		service_tester.RequireOption(t, options, constant.PrimaryDiskOption).(string),
	)
}
