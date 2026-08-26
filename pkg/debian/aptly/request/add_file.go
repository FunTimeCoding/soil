package request

type AddFile struct {
	ForceReplace bool `json:"forceReplace,omitempty"`
}
