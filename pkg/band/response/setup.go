package response

type Setup struct {
	ProvisioningMode  int `xml:"Body>AMT_SetupAndConfigurationService>ProvisioningMode"`
	ProvisioningState int `xml:"Body>AMT_SetupAndConfigurationService>ProvisioningState"`
}
