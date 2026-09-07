package consumer

// #include <stdlib.h>
import "C"

import "example/pkg/target"

func Fields() string {
	C.srand(1)

	return target.ItemFields
}
