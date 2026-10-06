package query

import "k8s.io/apimachinery/pkg/types"

type Patch struct {
	ResourceType string
	Name         string
	Namespace    string
	Patch        string
	Type         types.PatchType
	DryRun       bool
}
