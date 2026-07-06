package node

import "github.com/funtimecoding/soil/pkg/brave/bookmark/file"

func New(v *file.Node) *Node {
	result := Stub()
	result.Type = v.Type
	result.Name = v.Name
	result.Link = v.Locator
	result.Children = NewSlice(v.Children)
	result.Raw = v

	return result
}
