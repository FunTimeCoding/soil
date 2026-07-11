package service

import (
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/types"
)

type fieldFlip struct {
	object     types.Object
	newName    string
	references []resolve.Reference
}
