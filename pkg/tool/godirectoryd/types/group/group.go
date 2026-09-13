package group

type Group struct {
	Name              string   `json:"name"`
	Number            int      `json:"number"`
	Member            []string `json:"member"`
	DistinguishedName string   `json:"distinguished_name"`
}
