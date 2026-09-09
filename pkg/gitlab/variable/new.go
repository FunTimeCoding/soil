package variable

import "gitlab.com/gitlab-org/api/client-go/v2"

func New(v *gitlab.ProjectVariable) *Variable {
	return &Variable{
		Key:         v.Key,
		Value:       v.Value,
		Kind:        string(v.VariableType),
		Protected:   v.Protected,
		Masked:      v.Masked,
		Hidden:      v.Hidden,
		Literal:     v.Raw,
		Scope:       v.EnvironmentScope,
		Description: v.Description,
		Raw:         v,
	}
}
