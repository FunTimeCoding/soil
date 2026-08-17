package state

type State struct {
	Label              string
	Description        string
	Interface          string
	Protocol           string
	Direction          string
	SourceAddress      string
	SourcePort         string
	DestinationAddress string
	DestinationPort    string
	State              string
	Age                string
	Expires            string
	Packets            []int64
	Bytes              []int64
	Rule               string
}
