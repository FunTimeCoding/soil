package request

type PortForward struct {
	ResourceType string
	Name         string
	Namespace    string
	LocalPort    int
	TargetPort   int
}
