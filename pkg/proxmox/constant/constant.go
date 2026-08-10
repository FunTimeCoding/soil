package constant

const (
	HostEnvironment     = "PROXMOX_HOST"
	UserEnvironment     = "PROXMOX_USER"
	PasswordEnvironment = "PROXMOX_PASSWORD"
	TokenEnvironment    = "PROXMOX_TOKEN"
	SecretEnvironment   = "PROXMOX_SECRET"
	LogEnvironment      = "PROXMOX_LOG"
	VerboseEnvironment  = "PROXMOX_VERBOSE"
	InsecureEnvironment = "PROXMOX_INSECURE"
	TimeoutEnvironment  = "PROXMOX_TIMEOUT"
)

const (
	Port int = 8006
	Base     = "/api2/json"

	Name = "name"

	ContainerShutdownForce   = false
	ContainerShutdownTimeout = 60
)
