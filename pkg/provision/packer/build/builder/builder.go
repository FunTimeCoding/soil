package builder

type Builder struct {
	ImageLink       string     `json:"iso_url"`
	ImageChecksum   string     `json:"iso_checksum"`
	WebDirectory    string     `json:"http_directory"`
	OutputDirectory string     `json:"output_directory"`
	QemuBinary      string     `json:"qemu_binary"`
	Type            string     `json:"type"`
	Name            string     `json:"vm_name"`
	Accelerator     string     `json:"accelerator"`
	Cores           int        `json:"cpus"`
	Threads         int        `json:"threads"`
	Memory          int        `json:"memory"`
	DiskInterface   string     `json:"disk_interface,omitempty"`
	DiskSize        int        `json:"disk_size"`
	DiskFormat      string     `json:"format"`
	NetDevice       string     `json:"net_device,omitempty"`
	Headless        bool       `json:"headless,omitzero"`
	Display         string     `json:"display,omitempty"`
	QemuArguments   [][]string `json:"qemuargs"`
	BootWait        string     `json:"boot_wait"`
	BootCommand     []string   `json:"boot_command"`
	SshUsername     string     `json:"ssh_username"`
	SshPassword     string     `json:"ssh_password"`
	SshWaitTimeout  string     `json:"ssh_wait_timeout"`
	ShutdownCommand string     `json:"shutdown_command"`
}
