package packer

import (
	"fmt"
	"github.com/coreos/go-semver/semver"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/provision/packer/build"
	stringJoin "github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/debian/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
	"log"
	"runtime"
)

func (c *Client) Run(
	architecture string,
	imageName string,
	checksum string,
	v *semver.Version,
	preseed string,
	username string,
	password string,
	headless bool,
) {
	if headless {
		log.Panicf("headless not supported")
	}

	system.Move(
		preseed,
		join.Absolute(c.packerWebDirectory, constant.PreseedConfiguration),
	)
	// packer plugins install github.com/hashicorp/qemu
	b := build.New(architecture, 4444, username, password)
	b.SetBuilder(
		c.packerWebDirectory,
		c.packerOutputDirectory,
		headless,
		join.Absolute(c.workDirectory, imageName),
		checksum,
		v,
	)
	template := join.Absolute(c.workDirectory, "debian.json")
	system.SaveFile(template, fmt.Sprintln(notation.Format(b)))
	arguments := []string{
		"packer",
		"build",
		"-machine-readable",
		"-on-error",
		"cleanup",
		template,
	}
	fmt.Printf("Command: %s\n", stringJoin.Space(arguments...))

	if headless {
		var connectCommand string

		switch runtime.GOOS {
		case systemConstant.Linux:
			connectCommand = fmt.Sprintf(
				"nc -v 127.0.0.1 %d",
				b.CharacterDevicePort,
			)
		case systemConstant.Darwin:
			connectCommand = fmt.Sprintf(
				"socat -,raw,echo=0,escape=0x11 tcp:127.0.0.1:%d",
				b.CharacterDevicePort,
			)
		}

		// TODO: Typing happens via VNC, the character device does not reflect that
		//  Check VNC if typing actually happens
		//  Then proceed to actual unattended install
		if connectCommand != "" {
			fmt.Printf("Character device: %s\n", connectCommand)
		}
	}

	output, e := system.RunError(arguments...)

	if e != nil {
		fmt.Println("Error:")
	}

	fmt.Println(output)
}
