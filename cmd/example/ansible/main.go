package main

import "github.com/funtimecoding/soil/pkg/provision/ansible"

func main() {
	c := ansible.NewEnvironment()
	c.Execute(c.Playbook("site.yml"))
}
