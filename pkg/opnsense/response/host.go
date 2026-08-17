package response

type Host struct {
	Identifier       string `json:"uuid"`
	Host             string `json:"host"`
	Domain           string `json:"domain"`
	Address          string `json:"ip"`
	HardwareAddress  string `json:"hwaddr"`
	ClientIdentifier string `json:"client_id"`
	Description      string `json:"descr"`
}
