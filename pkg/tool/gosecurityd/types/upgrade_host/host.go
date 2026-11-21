package upgrade_host

type Host struct {
	Name      string
	Priority  int
	Labels    map[string]string
	Installed string
	Candidate string
	Binaries  []string
}
