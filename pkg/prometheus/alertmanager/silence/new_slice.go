package silence

import (
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/prometheus/alertmanager/api/v2/models"
)

func NewSlice(
	v models.GettableSilences,
	link *locator.Locator,
) []*Silence {
	var result []*Silence

	for _, s := range v {
		result = append(result, New(s, link))
	}

	return result
}
