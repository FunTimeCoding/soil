package client_tester

import (
	"github.com/funtimecoding/soil/pkg/gitlab"
	"testing"
)

type Tester struct {
	Client *gitlab.Client
	t      *testing.T
}
