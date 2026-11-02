package node_capacity

func New(
	cpuMillis int64,
	memoryBytes int64,
) *NodeCapacity {
	return &NodeCapacity{CpuMillis: cpuMillis, MemoryBytes: memoryBytes}
}
