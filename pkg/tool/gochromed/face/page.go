package face

import (
	"github.com/funtimecoding/soil/pkg/chromium/history"
	"github.com/funtimecoding/soil/pkg/chromium/snapshot"
)

type Page interface {
	Body() string
	Navigate(l string) error
	Snapshot() ([]*snapshot.Node, error)
	Screenshot() ([]byte, error)
	Evaluate(
		expression string,
		result any,
	) error
	ClickNode(backendNodeIdentifier int64) error
	FillNode(
		backendNodeIdentifier int64,
		value string,
		direct bool,
	) error
	History() (*history.Result, error)
}
