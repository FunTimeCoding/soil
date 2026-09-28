package band

type SetupResponse struct {
	ProvisioningMode  int `xml:"Body>AMT_SetupAndConfigurationService>ProvisioningMode"`
	ProvisioningState int `xml:"Body>AMT_SetupAndConfigurationService>ProvisioningState"`
}
