package response

type Member struct {
	Identifier  string `json:"id"`
	DisplayName string `json:"displayName"`
	User        string `json:"userId"`
}
