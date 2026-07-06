package build

import (
	"github.com/coreos/go-semver/semver"
	"github.com/funtimecoding/soil/pkg/provision/packer/build/builder"
	qemu "github.com/funtimecoding/soil/pkg/qemu/constant"
	"github.com/funtimecoding/soil/pkg/system/constant"
)

func (b *Build) SetBuilder(
	webDirectory string,
	outputDirectory string,
	headless bool,
	imageLink string,
	imageChecksum string,
	v *semver.Version,
) {
	var display string

	switch b.architecture {
	case constant.AMD64:
		display = "sdl"
	case constant.ARM64:
		display = "cocoa"
	}

	packerBuilder := builder.Builder{
		ImageLink:       imageLink,
		ImageChecksum:   imageChecksum,
		WebDirectory:    webDirectory,
		OutputDirectory: outputDirectory,
		Type:            "qemu",
		Name:            OutputFilename(v),
		QemuArguments:   b.arguments(headless),
		Cores:           2,
		Threads:         2,
		Memory:          2048,
		DiskSize:        81920,
		DiskFormat:      "qcow2",
		DiskInterface:   "virtio",
		NetDevice:       "virtio-net",
		Headless:        headless,
		Display:         display,
		BootWait:        "1s",
		BootCommand: []string{
			"<esc><wait>",
			"/install.amd/vmlinuz auto=true priority=critical",
			" url={{ .HTTPIP }}:{{ .HTTPPort }}/preseed.cfg",
			" initrd=/install.amd/initrd.gz",
			"<enter>",
		},
		SshUsername:     b.username,
		SshPassword:     b.password,
		SshWaitTimeout:  "10m",
		ShutdownCommand: "sudo -S shutdown -P now",
	}

	switch b.architecture {
	case constant.AMD64:
		packerBuilder.QemuBinary = qemu.AMD64Command
		packerBuilder.Accelerator = qemu.KVMAccelerator
	case constant.ARM64:
		packerBuilder.QemuBinary = qemu.ARM64Command
		packerBuilder.Accelerator = qemu.HVFAccelerator
	}

	b.Builders = []builder.Builder{packerBuilder}
}
