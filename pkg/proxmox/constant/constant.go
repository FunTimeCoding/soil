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

const LocalPrefix byte = 0x02
const (
	MaximumInstance = 255
	MaximumMachine  = 0xFFFFFFFF
)

const (
	BridgeKey          = "bridge"
	TagKey             = "tag"
	HardwareAddressKey = "hwaddr"
)

const (
	NodeType      = "node"
	MachineType   = "qemu"
	ContainerType = "lxc"
	StorageType   = "storage"
)

const (
	RunningStatus   = "running"
	OnlineStatus    = "online"
	AvailableStatus = "available"
)
