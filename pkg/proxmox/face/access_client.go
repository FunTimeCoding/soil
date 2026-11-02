package face

import "github.com/luthermonson/go-proxmox"

type AccessClient interface {
	Domains() (proxmox.Domains, error)
	Domain(name string) (*proxmox.Domain, error)
	Groups() (proxmox.Groups, error)
	Group(name string) (*proxmox.Group, error)
	Pools() ([]*proxmox.Pool, error)
	Pool(name string) (*proxmox.Pool, error)
	Roles() ([]*proxmox.Role, error)
	Role(name string) (*proxmox.Permission, error)
	Users() ([]*proxmox.User, error)
	User(name string) (*proxmox.User, error)
	MustDomains() proxmox.Domains
	MustDomain(name string) *proxmox.Domain
	MustGroups() proxmox.Groups
	MustGroup(name string) *proxmox.Group
	MustPools() []*proxmox.Pool
	MustPool(name string) *proxmox.Pool
	MustRoles() []*proxmox.Role
	MustRole(name string) *proxmox.Permission
	MustUsers() []*proxmox.User
	MustUser(name string) *proxmox.User
}
