package response

type NodeMetrics struct {
	Name          string `json:"name"`
	CPU           string `json:"cpu"`
	CPUPercent    string `json:"cpuPercent"`
	Memory        string `json:"memory"`
	MemoryPercent string `json:"memoryPercent"`
}
