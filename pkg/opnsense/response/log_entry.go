package response

type LogEntry struct {
	Timestamp       string `json:"__timestamp__"`
	Interface       string `json:"interface"`
	Action          string `json:"action"`
	Direction       string `json:"dir"`
	ProtocolName    string `json:"protoname"`
	Source          string `json:"src"`
	SourcePort      string `json:"srcport"`
	Destination     string `json:"dst"`
	DestinationPort string `json:"dstport"`
	RuleNumber      string `json:"rulenr"`
	RuleIdentifier  string `json:"rid"`
	Label           string `json:"label"`
}
