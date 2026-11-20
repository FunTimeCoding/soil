package system_result

type Finding struct {
	Vulnerability string
	Package       string
	Source        string
	Installed     string
	Fixed         string
	Candidate     string
	Fix           string
	Urgency       string
	Upgradable    bool
}
