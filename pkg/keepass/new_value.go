package keepass

import (
	"github.com/tobischo/gokeepasslib/v3"
	"github.com/tobischo/gokeepasslib/v3/wrappers"
)

func NewValue(
	key string,
	content string,
	protected bool,
) gokeepasslib.ValueData {
	return gokeepasslib.ValueData{
		Key: key,
		Value: gokeepasslib.V{
			Content:   content,
			Protected: wrappers.NewBoolWrapper(protected),
		},
	}
}
