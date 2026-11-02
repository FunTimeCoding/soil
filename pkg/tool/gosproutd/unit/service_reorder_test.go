package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/types/discovered_file"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/service_tester"
	"testing"
)

func TestSetPositionNotifies(t *testing.T) {
	s := service_tester.New(t)
	s.Service.Sync(
		[]discovered_file.File{
			{
				Name:        stringsConstant.LowerAlfa,
				Path:        "alfa.md",
				ContentHash: "a",
				Content:     "a",
			},
			{
				Name:        stringsConstant.LowerBravo,
				Path:        "bravo.md",
				ContentHash: "b",
				Content:     "b",
			},
			{
				Name:        stringsConstant.LowerCharlie,
				Path:        "charlie.md",
				ContentHash: "c",
				Content:     "c",
			},
		},
	)
	s.Notifier.Reset()
	s.Service.SetPosition(s.Service.Seeds()[2].Identifier, 1)
	assert.True(t, s.Notifier.Notified > 0)
	assert.String(t, "charlie", s.Service.Seeds()[0].Name)
}

func TestReorderNotifies(t *testing.T) {
	s := service_tester.New(t)
	s.Service.Sync(
		[]discovered_file.File{
			{
				Name:        stringsConstant.LowerAlfa,
				Path:        "alfa.md",
				ContentHash: "a",
				Content:     "a",
			},
			{
				Name:        stringsConstant.LowerBravo,
				Path:        "bravo.md",
				ContentHash: "b",
				Content:     "b",
			},
		},
	)
	seeds := s.Service.Seeds()
	s.Notifier.Reset()
	s.Service.Reorder([]uint{seeds[1].Identifier, seeds[0].Identifier})
	assert.True(t, s.Notifier.Notified > 0)
	assert.String(t, "bravo", s.Service.Seeds()[0].Name)
}
