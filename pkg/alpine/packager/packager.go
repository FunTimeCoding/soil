package packager

type Packager struct {
	ExecutablePath   string
	ExecutableName   string
	PackageVersion   string
	PackageName      string
	WorkDirectory    string
	ControlDirectory string
	ArchiveDirectory string
	OutputFile       string
}
