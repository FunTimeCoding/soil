package option

type Ansible struct {
	Address         string
	ServiceTokens   []string
	Version         string
	Repository      string
	ClonePath       string
	AnsiblePath     string
	Playbook        []string
	PostgresLocator string
	LitePath        string
}
