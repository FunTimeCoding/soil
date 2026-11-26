package argument

type Fill struct {
	UID           string `json:"uid"`
	Value         string `json:"value"`
	TabIdentifier string `json:"tab_id"`
	Snapshot      *bool  `json:"snapshot"`
	Direct        *bool  `json:"direct"`
}
