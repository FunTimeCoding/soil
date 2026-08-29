package request

type Host struct {
	Host             string `json:"host,omitempty"`
	Domain           string `json:"domain,omitempty"`
	Address          string `json:"ip,omitempty"`
	HardwareAddress  string `json:"hwaddr,omitempty"`
	ClientIdentifier string `json:"client_id,omitempty"`
	Description      string `json:"descr,omitempty"`
}
