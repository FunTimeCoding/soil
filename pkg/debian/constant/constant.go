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

	PackageKeyFields = 3

	HostEnvironment     = "APTLY_HOST"
	PortEnvironment     = "APTLY_PORT"
	InsecureEnvironment = "APTLY_INSECURE"
	UsernameEnvironment = "APTLY_USERNAME"
	PasswordEnvironment = "APTLY_PASSWORD"
)
