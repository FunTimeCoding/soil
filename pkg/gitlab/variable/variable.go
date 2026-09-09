package variable

import "gitlab.com/gitlab-org/api/client-go/v2"

type Variable struct {
	Key         string
	Value       string
	Kind        string
	Protected   bool
	Masked      bool
	Hidden      bool
	Literal     bool // upstream raw: the value is not expanded
	Scope       string
	Description string
	Raw         *gitlab.ProjectVariable
}
