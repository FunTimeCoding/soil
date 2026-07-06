package godebian

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/debian"
	debianConstant "github.com/funtimecoding/soil/pkg/system/debian/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
	"github.com/funtimecoding/soil/pkg/tool/godebian/option"
)

func Run(o *option.Debian) {
	architecture := constant.AMD64
	packageName := debian.PackageVersion(
		o.Executable,
		o.PackageVersion,
		1,
		architecture,
	)
	packageDirectory := join.Absolute(system.WorkDirectory(), packageName)
	debianDirectory := join.Absolute(
		packageDirectory,
		debianConstant.PackageConfigurationDirectory,
	)
	system.MakeDirectory(debianDirectory)
	system.SaveFile(
		join.Absolute(debianDirectory, debianConstant.ControlFile),
		fmt.Sprintf(
			"Package: %s\nVersion: %s\nArchitecture: %s\nMaintainer: %s <%s>\nDescription: Short stub description.\n Long stub description.\n",
			o.Executable,
			o.PackageVersion,
			architecture,
			o.MaintainerName,
			o.MaintainerEmail,
		),
	)

	if o.SystemdUnit {
		unit := join.Absolute(
			packageDirectory,
			constant.Library,
			"systemd",
			"system",
		)
		system.MakeDirectory(unit)
		system.SaveFile(
			join.Absolute(unit, fmt.Sprintf("%s.service", o.Executable)),
			fmt.Sprintf(
				"[Unit]\nDescription=%s stub description\nAfter=network.target\n\n[Service]\nType=simple\nExecStart=/usr/local/bin/%s\n\n[Install]\nWantedBy=multi-user.target\n",
				o.Executable,
				o.Executable,
			),
		)
	}

	bin := join.Absolute(
		packageDirectory,
		constant.Resources,
		constant.Local,
		constant.Binary,
	)
	system.MakeDirectory(bin)
	system.Move(o.Executable, join.Absolute(bin, o.Executable))
	fmt.Println(
		system.Run(
			"dpkg-deb",
			"--build",
			"--root-owner-group",
			packageName,
		),
	)
}
