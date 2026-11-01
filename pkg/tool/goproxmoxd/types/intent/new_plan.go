package intent

func NewPlan(hypervisor string) *Plan {
	return &Plan{Hypervisor: hypervisor}
}
