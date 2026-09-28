package constant

const (
	HostEnvironment     = "FRITZ_HOST"
	UserEnvironment     = "FRITZ_USER"
	PasswordEnvironment = "FRITZ_PASSWORD"
	DefaultHost         = "192.168.178.1"
	Port                = 49000
)

const (
	DevicePath      = "/upnp/control/deviceinfo"
	DeviceService   = "urn:dslforum-org:service:DeviceInfo:1"
	DeviceAction    = "GetInfo"
	InternetPath    = "/upnp/control/wanipconnection1"
	InternetService = "urn:dslforum-org:service:WANIPConnection:1"
	PPPPath         = "/upnp/control/wanpppconn1"
	PPPService      = "urn:dslforum-org:service:WANPPPConnection:1"
	ExternalAction  = "X_AVM-DE_GetExternalIPv6Address"
	PrefixAction    = "X_AVM-DE_GetIPv6Prefix"
)
