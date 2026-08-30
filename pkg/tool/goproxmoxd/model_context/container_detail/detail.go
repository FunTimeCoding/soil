package container_detail

type Detail struct {
	Identifier  uint64 `json:"identifier"`
	Name        string `json:"name"`
	Node        string `json:"node"`
	Status      string `json:"status"`
	CPUs        int    `json:"cpus"`
	MaxMem      uint64 `json:"max_mem"`
	MaxDisk     uint64 `json:"max_disk"`
	MaxSwap     uint64 `json:"max_swap"`
	Uptime      uint64 `json:"uptime"`
	Tags        string `json:"tags,omitempty"`
	Hostname    string `json:"hostname,omitempty"`
	Description string `json:"description,omitempty"`
	Arch        string `json:"arch,omitempty"`
	Cores       int    `json:"cores,omitzero"`
	Memory      int    `json:"memory,omitzero"`
}
