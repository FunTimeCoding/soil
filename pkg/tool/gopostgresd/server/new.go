package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/backup"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/catalog"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/query"
)

func New(
	c *catalog.Service,
	q *query.Service,
	b *backup.Service,
	i *instance.Inventory,
	r face.Reporter,
) *Server {
	return &Server{
		catalog:   c,
		query:     q,
		backup:    b,
		inventory: i,
		reporter:  r,
	}
}
