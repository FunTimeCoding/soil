package installed

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
)

func finding(
	key string,
	text string,
	b *Binary,
) *concern.Concern {
	result := concern.NewFile(
		key,
		fmt.Sprintf(text, b.Directory, b.Name),
		b.Path,
		false,
	)
	result.Planned = true

	return result
}
