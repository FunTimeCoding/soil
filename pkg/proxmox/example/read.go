package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/proxmox"
)

func Read() {
	p := proxmox.NewEnvironment()
	fmt.Printf("Version: %+v\n", p.MustVersion())
	fmt.Printf("Cluster: %+v\n", p.MustCluster())

	for _, n := range p.MustNodes() {
		fmt.Printf("Node list: %+v\n", n)
		o := p.MustNode(n.Node)
		fmt.Printf("Node: %+v\n", o)

		for _, m := range p.MustMachines(o) {
			fmt.Printf("  Machine list: %+v\n", m)
			fmt.Printf(
				"  Machine: %+v\n",
				p.MustMachine(o, int(m.VMID)),
			)
		}

		for _, c := range p.MustContainers(o) {
			fmt.Printf("  Container list: %+v\n", c)
			fmt.Printf(
				"  Container: %+v\n",
				p.MustContainer(o, int(c.VMID)),
			)
		}
	}

	for _, d := range p.MustDomains() {
		fmt.Printf("Domain list: %+v\n", d)
		fmt.Printf("Domain: %+v\n", p.MustDomain(d.Realm))
	}

	for _, g := range p.MustGroups() {
		fmt.Printf("Group list: %+v\n", g)
		fmt.Printf("Group: %+v\n", p.MustGroup(g.GroupID))
	}

	for _, o := range p.MustPools() {
		fmt.Printf("Pool list: %+v\n", o)
		fmt.Printf("Pool: %+v\n", p.MustPool(o.PoolID))
	}

	for _, r := range p.MustRoles() {
		fmt.Printf("Role list: %+v\n", r)
		fmt.Printf("Role: %+v\n", p.MustRole(r.RoleID))
	}

	for _, u := range p.MustUsers() {
		fmt.Printf("User list: %+v\n", u)
		fmt.Printf("User: %+v\n", p.MustUser(u.UserID))
	}
}
