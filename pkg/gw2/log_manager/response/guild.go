package response

type Guild struct {
	Identifier string `json:"Id"`
	Name       string `json:"Name"`
	Tag        string `json:"Tag"`
	Emblem     Emblem `json:"Emblem"`
}
