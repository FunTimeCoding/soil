package lookup

import (
	"fmt"
	"go/token"
)

func TakenAt(
	set *token.FileSet,
	name string,
	at token.Pos,
) string {
	position := set.Position(at)

	return fmt.Sprintf(
		"%s is taken at %s:%d",
		name,
		position.Filename,
		position.Line,
	)
}
