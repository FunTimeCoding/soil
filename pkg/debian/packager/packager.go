package packager

type Packager struct {
	ExecutablePath    string
	ExecutableName    string
	PackageName       string
	PackageVersion    string
	Root              string
	ConfigurationRoot string
	UnitRoot          string
	ControlFile       string
	BinaryRoot        string
	Architecture      string
	MaintainerName    string
	MaintainerMail    string
}
