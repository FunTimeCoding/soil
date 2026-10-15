package configuration_request

type Request struct {
	Name      string          `json:"name,omitempty"`
	Weight    *int32          `json:"weight,omitempty"`
	Payload   *map[string]any `json:"data,omitempty"`
	Roles     *[]int32        `json:"roles,omitempty"`
	Clusters  *[]int32        `json:"clusters,omitempty"`
	Platforms *[]int32        `json:"platforms,omitempty"`
	Sites     *[]int32        `json:"sites,omitempty"`
	Tenants   *[]int32        `json:"tenants,omitempty"`
	Tags      *[]string       `json:"tags,omitempty"`
}
