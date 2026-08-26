package gopackagedeb

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	debianConstant "github.com/funtimecoding/soil/pkg/debian/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/semver"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gopackagedeb/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopackagedeb/option"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version)
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.String(
		argumentConstant.MaintainerName,
		environment.Optional(constant.MaintainerNameEnvironment),
		"AN Other",
	)
	a.String(
		argumentConstant.MaintainerMail,
		environment.Optional(constant.MaintainerMailEnvironment),
		"another@example.org",
	)
	a.Boolean(argumentConstant.Unit, false, "Create systemd unit")
	a.String(
		argumentConstant.OnUpgrade,
		debianConstant.UpgradeRestart,
		"Systemd service handling on package upgrade: restart or keep",
	)
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Executable = a.RequiredPositional(0, "EXECUTABLE")
	o.PackageVersion = semver.Trim(a.RequiredPositional(1, "PACKAGE_VERSION"))
	o.MaintainerName = a.GetString(argumentConstant.MaintainerName)
	o.MaintainerMail = a.GetString(argumentConstant.MaintainerMail)
	o.SystemdUnit = a.GetBoolean(argumentConstant.Unit)
	o.UpgradeMode = a.GetString(argumentConstant.OnUpgrade)

	if o.UpgradeMode != debianConstant.UpgradeRestart &&
		o.UpgradeMode != debianConstant.UpgradeKeep {
		system.Exitf(
			1,
			"unknown upgrade mode: %s\nexpected %s or %s\n",
			o.UpgradeMode,
			debianConstant.UpgradeRestart,
			debianConstant.UpgradeKeep,
		)
	}

	if !o.SystemdUnit && a.Changed(argumentConstant.OnUpgrade) {
		system.Exitf(
			1,
			"--%s requires --%s\n",
			argumentConstant.OnUpgrade,
			argumentConstant.Unit,
		)
	}

	Run(o)
}
