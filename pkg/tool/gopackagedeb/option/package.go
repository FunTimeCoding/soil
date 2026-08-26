package option

type Package struct {
	Executable     string
	PackageVersion string
	MaintainerName string
	MaintainerMail string
	SystemdUnit    bool
	UpgradeMode    string
}
