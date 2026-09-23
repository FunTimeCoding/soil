package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/types/lock_detail"
	"testing"
	"time"
)

func TestLockDetailParsesAnnotation(t *testing.T) {
	d := lock_detail.New(
		annotated(
			map[string]string{
				constant.LockAnnotation: constant.FixtureLockDetail,
			},
		),
	)
	assert.NotNil(t, d)
	assert.String(t, "11111111-2222-4333-8444-555555555555", d.Identifier)
	assert.String(t, "OperationTypePlan", d.Operation)
	assert.String(t, "user@host.example", d.Who)
	assert.Time(
		t,
		time.Date(2020, 1, 2, 3, 4, 5, 123456789, time.UTC),
		d.Created,
	)
}

func TestLockDetailMissingLease(t *testing.T) {
	assert.Nil(t, lock_detail.New(nil))
}

func TestLockDetailUnlocked(t *testing.T) {
	assert.Nil(t, lock_detail.New(annotated(map[string]string{})))
}

func TestLockDetailUnreadableAnnotation(t *testing.T) {
	assert.Nil(
		t,
		lock_detail.New(
			annotated(map[string]string{constant.LockAnnotation: "{"}),
		),
	)
}
