package response

type ScreenService struct {
	EnabledState int `xml:"Body>PullResponse>Items>CIM_KVMRedirectionSAP>EnabledState"`
}
