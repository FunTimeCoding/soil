package argument

type CreateTunnel struct {
	Name          string `json:"name"`
	Encapsulation string `json:"encapsulation"`
	Group         string `json:"group"`
}
