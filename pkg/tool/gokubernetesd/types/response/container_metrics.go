package response

type ContainerMetrics struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
	Container string `json:"container"`
	CPU       string `json:"cpu"`
	Memory    string `json:"memory"`
	CpuMillis int64  `json:"-"`
}
