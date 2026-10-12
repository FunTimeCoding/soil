package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store"
)

type Service struct {
	client face.NetboxSource
	store  *store.Store
}
