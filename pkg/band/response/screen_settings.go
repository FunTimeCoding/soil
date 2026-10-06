package response

type ScreenSettings struct {
	EnabledInFirmware bool `xml:"Body>IPS_KVMRedirectionSettingData>EnabledByMEBx"`
	SessionTimeout    int  `xml:"Body>IPS_KVMRedirectionSettingData>SessionTimeout"`
}
