package terraform_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/types/plan"
)

func Plan() *plan.Plan {
	result := plan.New()
	errors.PanicOnError(
		notation.DecodeBytes([]byte(constant.FixturePlan), result),
	)

	return result
}
