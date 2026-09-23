package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/choice"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/cruise"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/frame"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/seed"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/turn"
	"gorm.io/gorm"
	"time"
)

func New(
	d *gorm.DB,
	clock func() time.Time,
) *Store {
	errors.PanicOnError(
		d.AutoMigrate(
			seed.Stub(),
			frame.Stub(),
			decision.Stub(),
			choice.Stub(),
			turn.Stub(),
			cruise.Stub(),
		),
	)

	return &Store{mapper: d, clock: clock}
}
