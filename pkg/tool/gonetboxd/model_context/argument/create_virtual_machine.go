package argument

type CreateVirtualMachine struct {
	Name    string `json:"name"`
	Cluster string `json:"cluster"`
}
