package merge_request

import (
	"github.com/funtimecoding/soil/pkg/face"
	"gitlab.com/gitlab-org/api/client-go/v3"
	"time"
)

type Request struct {
	Project      int64
	Identifier   int64
	Title        string
	State        string
	Link         string
	Create       *time.Time
	Update       *time.Time
	Author       string
	Reviewers    []string
	Assignees    []string
	SourceBranch string
	TargetBranch string
	Draft        bool
	MergeStatus  string
	Notes        int64
	ageColor     face.SprintFunction
	Raw          *gitlab.BasicMergeRequest
}
