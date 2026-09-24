package score

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"strings"
)

func ParsePolicy(s string) constant.Policy {
	switch strings.ToLower(s) {
	case constant.PolicyPessimistic, "":
		return constant.Pessimistic
	case constant.PolicyOptimistic:
		return constant.Optimistic
	case constant.PolicySkip:
		return constant.Skip
	default:
		panic(fmt.Sprintf(constant.UnknownPolicy, s))
	}
}
