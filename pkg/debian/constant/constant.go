package constant

const (
	DpkgDeb        = "dpkg-deb"
	BuildArgument  = "--build"
	RootOwnerGroup = "--root-owner-group"

	SystemdDirectory = "systemd"
	SystemDirectory  = "system"

	ServiceExtension = "service"

	PostInstallScript = "postinst"
	PreRemoveScript   = "prerm"
	PostRemoveScript  = "postrm"

	UpgradeRestart = "restart"
	UpgradeKeep    = "keep"
)
