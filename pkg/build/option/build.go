package option

type Build struct {
	Name            string
	MainPath        string
	Output          string
	BuildTags       string
	Module          string
	Version         string
	CopyToBin       bool
	Native          bool
	OperatingSystem string
	Architecture    string
	LinuxAMD64      bool
	DarwinARM64     bool
	DarwinAMD64     bool
}
