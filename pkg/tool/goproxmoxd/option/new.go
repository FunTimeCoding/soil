package option

func New() *Proxmox {
	return &Proxmox{Hypervisor: BuildHypervisor, Shell: BuildShell}
}
