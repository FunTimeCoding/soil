package flag

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/vulnerability"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store"
)

type Service struct {
	store         *store.Store
	vulnerability *vulnerability.Service
}
