package argument

type CreateCluster struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Site string `json:"site"`
}
