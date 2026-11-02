package sample

func New(hypervisor string) *Sample {
	return &Sample{Hypervisor: hypervisor}
}
