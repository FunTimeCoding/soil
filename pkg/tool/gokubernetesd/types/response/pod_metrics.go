package response

type PodMetrics struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
	CPU       string `json:"cpu"`
	Memory    string `json:"memory"`
	CpuMillis int64  `json:"-"`
}
