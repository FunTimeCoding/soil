package tree

import "gitlab.com/gitlab-org/api/client-go/v2"

func New(v *gitlab.TreeNode) *Node {
	return &Node{
		Identifier: v.ID,
		Name:       v.Name,
		Kind:       v.Type,
		Path:       v.Path,
		Mode:       v.Mode,
		Raw:        v,
	}
}
