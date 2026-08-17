package response

type Pool struct {
	Identifier   string `json:"uuid"`
	Interface    string `json:"interface"`
	StartAddress string `json:"start_addr"`
	EndAddress   string `json:"end_addr"`
	LeaseTime    string `json:"lease_time"`
	Domain       string `json:"domain"`
	Description  string `json:"description"`
}
