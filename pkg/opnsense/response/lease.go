package response

type Lease struct {
	Address              string   `json:"address"`
	HardwareAddress      string   `json:"hwaddr"`
	Hostname             string   `json:"hostname"`
	ClientIdentifier     string   `json:"client_id"`
	Expire               int64    `json:"expire"`
	InterfaceName        string   `json:"if_name"`
	InterfaceDescription string   `json:"if_descr"`
	Reserved             []string `json:"is_reserved"`
}
