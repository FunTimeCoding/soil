package response

type Minion struct {
	Identifier        string              `json:"id"`
	Host              string              `json:"host"`
	AbsoluteDomainName              string              `json:"fqdn"`
	Domain            string              `json:"domain"`
	Localhost         string              `json:"localhost"`
	Master            string              `json:"master"`
	OperatingSystem                string              `json:"os"`
	OperatingSystemFullName        string              `json:"osfullname"`
	OperatingSystemRelease         string              `json:"osrelease"`
	OperatingSystemFamily          string              `json:"os_family"`
	Kernel            string              `json:"kernel"`
	KernelRelease     string              `json:"kernelrelease"`
	CPUArch           string              `json:"cpuarch"`
	NumCPUS           int                 `json:"num_cpus"`
	MemTotal          int                 `json:"mem_total"`
	Virtual           string              `json:"virtual"`
	IPv4              []string            `json:"ipv4"`
	IPv6              []string            `json:"ipv6"`
	IP4Interfaces     map[string][]string `json:"ip4_interfaces"`
	SaltVersion       string              `json:"saltversion"`
	MachineIdentifier string              `json:"machine_id"`
	Manufacturer      string              `json:"manufacturer"`
	ProductName       string              `json:"productname"`
	SerialNumber      string              `json:"serialnumber"`
}
