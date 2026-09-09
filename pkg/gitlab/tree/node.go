package tree

import "gitlab.com/gitlab-org/api/client-go/v2"

type Node struct {
	Identifier string
	Name       string
	Kind       string
	Path       string
	Mode       string
	Raw        *gitlab.TreeNode
}
