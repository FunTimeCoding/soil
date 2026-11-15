package pointer

import "github.com/funtimecoding/soil/pkg/lint/constant"

type Resolution struct {
	Concern string
	Reason  constant.Reason
	Target  string
	Hint    string
}
