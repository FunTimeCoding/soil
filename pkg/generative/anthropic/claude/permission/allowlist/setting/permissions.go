package setting

type Permissions struct {
	Allow                 []string `json:"allow"`
	AdditionalDirectories []string `json:"additionalDirectories"`
}
