package import_decision

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/types/import_entry"
)

type Decision struct {
	Entry     *import_entry.Entry
	Want      string
	Kind      constant.DecisionKind
	Rule      constant.DecisionRule
	Reason    string
	Keeper    string
	Subsystem string
}
