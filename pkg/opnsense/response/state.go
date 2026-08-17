package response

type State struct {
	Label              string  `json:"label"`
	Description        string  `json:"descr"`
	Interface          string  `json:"iface"`
	Protocol           string  `json:"proto"`
	Direction          string  `json:"direction"`
	SourceAddress      string  `json:"src_addr"`
	SourcePort         string  `json:"src_port"`
	DestinationAddress string  `json:"dst_addr"`
	DestinationPort    string  `json:"dst_port"`
	State              string  `json:"state"`
	Age                string  `json:"age"`
	Expires            string  `json:"expires"`
	Packets            []int64 `json:"pkts"`
	Bytes              []int64 `json:"bytes"`
	Rule               string  `json:"rule"`
}
