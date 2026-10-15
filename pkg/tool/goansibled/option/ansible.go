package option

import "github.com/funtimecoding/soil/pkg/web/locator"

type Ansible struct {
	Address          string
	ServiceTokens    []string
	Repository       string
	ClonePath        string
	AnsiblePath      string
	Playbook         []string
	PostgresLocator  string
	LitePath         string
	InventoryLocator *locator.Locator
	InventoryToken   string
}
