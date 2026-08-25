package response

type ManagerNode struct {
	SystemIdentifier string `json:"systemId"`
	IPAddress        string `json:"ipAddress"`
	Version          string `json:"version"`
	Revision         string `json:"revision"`
}
