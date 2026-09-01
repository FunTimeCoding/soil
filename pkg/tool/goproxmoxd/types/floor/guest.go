package floor

type Guest struct {
	Hypervisor  string
	Node        string
	Kind        string
	Identifier  uint64
	Name        string
	Status      string
	Unbacked    bool
	Processor   float64
	Memory      uint64
	MemoryTotal uint64
}
