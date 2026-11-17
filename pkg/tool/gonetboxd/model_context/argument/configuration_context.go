package argument

type ConfigurationContext struct {
	Name      string          `json:"name"`
	Payload   *map[string]any `json:"payload"`
	Weight    *float64        `json:"weight"`
	Roles     *[]string       `json:"roles"`
	Clusters  *[]string       `json:"clusters"`
	Platforms *[]string       `json:"platforms"`
	Sites     *[]string       `json:"sites"`
	Tenants   *[]string       `json:"tenants"`
	Tags      *[]string       `json:"tags"`
}
