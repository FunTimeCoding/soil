package jc

type Output struct {
	Proto             string  `json:"proto"`
	RecvQ             *int    `json:"recv_q"`
	SendQ             int     `json:"send_q,omitzero"`
	LocalAddress      string  `json:"local_address,omitempty"`
	ForeignAddress    string  `json:"foreign_address,omitempty"`
	State             string  `json:"state,omitempty"`
	ProgramName       string  `json:"program_name,omitempty"`
	Kind              string  `json:"kind"`
	Pid               int     `json:"pid,omitzero"`
	LocalPort         string  `json:"local_port,omitempty"`
	ForeignPort       string  `json:"foreign_port,omitempty"`
	TransportProtocol *string `json:"transport_protocol"`
	NetworkProtocol   string  `json:"network_protocol"`
	LocalPortNum      int     `json:"local_port_num,omitzero"`
}
