package policy

type Policy struct {
	KeyName string   `json:"tsigKeyName"`
	Domain  string   `json:"domain"`
	Types   []string `json:"allowedTypes"`
}
