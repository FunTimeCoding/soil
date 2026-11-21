package upgrade_group

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/upgrade_host"

type Group struct {
	Release          string
	Source           string
	Kernel           bool
	Hosts            []*upgrade_host.Host
	Vulnerabilities  []string
	Fixed            string
	InstalledLowest  string
	InstalledHighest string
	Candidate        string
	TopScore         float64
}
