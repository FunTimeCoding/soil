package store_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
)

func FilterOutcome() *result.Outcome {
	return result.NewOutcome(
		[]result.Search{
			{
				Title: "alfa",
				Metadata: map[string][]string{
					constant.FixtureTagKey:    {constant.FixtureBuildValue},
					constant.FixtureAuthorKey: {"one"},
				},
			},
			{
				Title: "bravo",
				Metadata: map[string][]string{
					constant.FixtureTagKey:    {"groom"},
					constant.FixtureAuthorKey: {"one"},
				},
			},
			{
				Title: "charlie",
				Metadata: map[string][]string{
					constant.FixtureTagKey:    {constant.FixtureBuildValue},
					constant.FixtureAuthorKey: {"two"},
				},
			},
		},
	)
}
