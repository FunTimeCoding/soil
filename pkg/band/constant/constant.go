package constant

const (
	DefaultUser = "admin"

	SecurePort = 16993

	Path        = "/wsman"
	ContentType = "application/soap+xml;charset=UTF-8"

	GetAction        = "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get"
	EnumerateAction  = "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate"
	PullAction       = "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull"
	PowerStateAction = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService/RequestPowerStateChange"
	AnonymousRole    = "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous"

	EnumerationNamespace = "http://schemas.xmlsoap.org/ws/2004/09/enumeration"

	PowerOn              = 2
	PowerCycle           = 5
	PowerOffSoft         = 8
	PowerOffSoftGraceful = 12
	PowerResetGraceful   = 14

	AssociatedPowerResource = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_AssociatedPowerManagementService"
	PowerServiceResource    = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService"
	ComputerSystemResource  = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ComputerSystem"
	GeneralSettingsResource = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_GeneralSettings"
	SetupResource           = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService"
	RedirectionResource     = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService"
	ScreenSettingsResource  = "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_KVMRedirectionSettingData"
	ConsentResource         = "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService"
)
