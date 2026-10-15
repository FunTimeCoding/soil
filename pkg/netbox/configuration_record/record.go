package configuration_record

type Record struct {
	Identifier  int32          `json:"id"`
	Name        string         `json:"name"`
	Weight      int32          `json:"weight"`
	Description string         `json:"description"`
	Active      bool           `json:"is_active"`
	Payload     map[string]any `json:"data"`
	Roles       []Named        `json:"roles"`
	Clusters    []Named        `json:"clusters"`
	Platforms   []Named        `json:"platforms"`
	Sites       []Named        `json:"sites"`
	Tenants     []Named        `json:"tenants"`
	Tags        []string       `json:"tags"`
}
