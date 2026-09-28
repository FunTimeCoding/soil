package library

type Library struct {
	Identifier     string `json:"id"`
	Name           string `json:"name"`
	CollectionType string `json:"collection_type,omitempty"`
}
