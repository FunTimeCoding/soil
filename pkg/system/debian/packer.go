package debian

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/provision/packer"
	"github.com/funtimecoding/soil/pkg/provision/packer/build"
	"github.com/funtimecoding/soil/pkg/provision/packer/constant"
	"github.com/funtimecoding/soil/pkg/system"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/debian/image"
	"github.com/funtimecoding/soil/pkg/system/debian/image/checksum"
	"github.com/funtimecoding/soil/pkg/system/debian/preseed"
	"github.com/funtimecoding/soil/pkg/system/debian/release"
	"github.com/funtimecoding/soil/pkg/system/join"
	"log"
)

func (c *Client) Packer(
	r *release.Release,
	architecture string,
) {
	// Start before running this, not yet supported on M1:
	// src/virtualization-tools/bin/libvirt/run.sh
	// src/virtualization-tools/bin/libvirt/virtlogd.sh
	// virt-manager --debug
	switch architecture {
	case systemConstant.AMD64:
		// pass
	case systemConstant.ARM64:
		// pass
	default:
		log.Panicf("unsupported architecture: %s", architecture)
	}

	v := r.Version()
	outputImage := build.OutputFilename(v)
	outputImagePath := join.Absolute(
		c.workDirectory,
		constant.PackerDirectory,
		constant.OutputDirectory,
		outputImage,
	)

	if system.FileExists(outputImagePath) {
		fmt.Printf("Already exists: %s\n", outputImagePath)
		fmt.Printf(
			"To spawn a clone:\n  mkdir /tmp/test\n  cd /tmp/test\n  cp %s .\n  qemu-system-x86_64 -machine type=q35,accel=kvm -m 2G -cpu host -smp 2 -drive file=%s,format=qcow2 -netdev user,id=net0 -device virtio-net-pci,netdev=net0",
			outputImagePath,
			outputImage,
		)

		return
	}

	n := Codename(v)
	c.DownloadImage(r, architecture)
	preseedName := preseed.Name(n)
	preseedPath := join.Absolute(c.workDirectory, preseedName)
	preseed.Download(n, preseedPath)
	imageName := image.Name(v, architecture)
	sum, found := checksum.Map(c.workDirectory)[imageName]

	if !found {
		log.Panicf("Sum not found: %s", imageName)
	}

	rootPassword := "root"
	userFullName := "Debian User"
	username := "debian"
	password := "debian"
	late := fmt.Sprintf(
		"echo \"%s ALL=(ALL:ALL) NOPASSWD:ALL\" > /target/etc/sudoers.d/%s && chmod 440 /target/etc/sudoers.d/%s",
		username,
		username,
		username,
	)
	system.SaveFile(
		preseedPath,
		fmt.Sprintf(
			"%s\nd-i base-installer/install-recommends boolean false\nd-i grub-installer/bootdev string default\ntasksel tasksel/skip-tasks string standard\ntasksel tasksel/first multiselect ssh-server\nd-i pkgsel/include string sudo\nd-i passwd/root-password password %s\nd-i passwd/root-password-again password %s\nd-i passwd/user-fullname string %s\nd-i passwd/username string %s\nd-i passwd/user-password password %s\nd-i passwd/user-password-again password %s\nd-i preseed/late_command string \\\n  %s\n",
			system.ReadFile(c.workDirectory, preseedName),
			rootPassword,
			rootPassword,
			userFullName,
			username,
			password,
			password,
			late,
		),
	)
	packer.New(c.workDirectory).Run(
		architecture,
		imageName,
		sum,
		v,
		preseedPath,
		username,
		password,
		false,
	)
}
