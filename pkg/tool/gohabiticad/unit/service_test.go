package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/unit/service_tester"
	"testing"
)

func TestEquipAnswersEquippedGear(t *testing.T) {
	v := service_tester.Service(t, service_tester.NewSource(t))
	g, e := v.Equip("weapon_warrior_1")
	assert.FatalOnError(t, e)
	assert.String(t, "weapon_warrior_1", g.Equipped["weapon"])
}
