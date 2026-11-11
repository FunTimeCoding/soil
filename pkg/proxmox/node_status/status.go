package node_status

type Status struct {
	BootDetail      BootDetail      `json:"boot-info"`
	Idle            float64         `json:"idle"`
	Kernel          Kernel          `json:"current-kernel"`
	KernelVersion   string          `json:"kversion"`
	LoadAverage     []string        `json:"loadavg"`
	Memory          Memory          `json:"memory"`
	Processor       float64         `json:"cpu"`
	ProcessorDetail ProcessorDetail `json:"cpuinfo"`
	ProxmoxVersion  string          `json:"pveversion"`
	RootFilesystem  RootFilesystem  `json:"rootfs"`
	SamePageMerge   SamePageMerge   `json:"ksm"`
	Swap            Swap            `json:"swap"`
	Uptime          int             `json:"uptime"`
	Wait            float64         `json:"wait"`
}
