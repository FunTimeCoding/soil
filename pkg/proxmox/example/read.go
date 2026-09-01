package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/proxmox"
)

func Read() {
	p := proxmox.NewEnvironment()
	console.Format("Version: %+v\n", p.MustVersion())
	console.Format("Cluster: %+v\n", p.MustCluster())

	for _, n := range p.MustNodes() {
		console.Format("Node list: %+v\n", n)
		o := p.MustNode(n.Node)
		console.Format("Node: %+v\n", o)

		for _, m := range p.MustMachines(o) {
			console.Format("  Machine list: %+v\n", m)
			console.Format("  Machine: %+v\n", p.MustMachine(o, int(m.VMID)))
		}

		for _, c := range p.MustContainers(o) {
			console.Format("  Container list: %+v\n", c)
			console.Format(
				"  Container: %+v\n",
				p.MustContainer(o, int(c.VMID)),
			)
		}
	}

	for _, d := range p.MustDomains() {
		console.Format("Domain list: %+v\n", d)
		console.Format("Domain: %+v\n", p.MustDomain(d.Realm))
	}

	for _, g := range p.MustGroups() {
		console.Format("Group list: %+v\n", g)
		console.Format("Group: %+v\n", p.MustGroup(g.GroupID))
	}

	for _, o := range p.MustPools() {
		console.Format("Pool list: %+v\n", o)
		console.Format("Pool: %+v\n", p.MustPool(o.PoolID))
	}

	for _, r := range p.MustRoles() {
		console.Format("Role list: %+v\n", r)
		console.Format("Role: %+v\n", p.MustRole(r.RoleID))
	}

	for _, u := range p.MustUsers() {
		console.Format("User list: %+v\n", u)
		console.Format("User: %+v\n", p.MustUser(u.UserID))
	}
}
