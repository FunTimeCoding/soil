package tag

func New(v *Response) *Tag {
	return &Tag{Name: v.Name, LastUpdated: v.LastUpdated}
}
