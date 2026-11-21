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
	Score         float64
	Severity      string
	Upgradable    bool
	Reboot        bool
	Kernel        bool
	Backports     bool
}
