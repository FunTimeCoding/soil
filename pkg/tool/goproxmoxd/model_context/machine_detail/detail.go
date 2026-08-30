package machine_detail

type Detail struct {
	Identifier          uint64  `json:"identifier"`
	Name                string  `json:"name"`
	Node                string  `json:"node"`
	Status              string  `json:"status"`
	CPU                 float64 `json:"cpu"`
	CPUs                int     `json:"cpus"`
	Mem                 uint64  `json:"mem"`
	MaxMem              uint64  `json:"max_mem"`
	Disk                uint64  `json:"disk"`
	MaxDisk             uint64  `json:"max_disk"`
	Uptime              uint64  `json:"uptime"`
	Tags                string  `json:"tags,omitempty"`
	Description         string  `json:"description,omitempty"`
	OperatingSystemType string  `json:"os_type,omitempty"`
	Sockets             int     `json:"sockets,omitzero"`
	Cores               int     `json:"cores,omitzero"`
}
