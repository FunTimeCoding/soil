package argument

type Click struct {
	UID           string `json:"uid"`
	TabIdentifier string `json:"tab_id"`
	Snapshot      *bool  `json:"snapshot"`
}
