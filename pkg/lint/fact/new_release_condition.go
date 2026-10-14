package fact

import "github.com/funtimecoding/soil/pkg/lint/constant"

func NewReleaseCondition(
	holder string,
	field string,
) *Condition {
	return &Condition{
		Kind:   constant.ReleasesCondition,
		Holder: holder,
		Field:  field,
	}
}
