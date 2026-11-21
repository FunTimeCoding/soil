package remediation_item

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/flag"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/affected_count"
)

type Item struct {
	Subject         string
	Kind            string
	Release         string
	Source          string
	Kernel          bool
	Reason          string
	Hosts           []string
	Vulnerabilities []string
	TopScore        float64
	Unknown         bool
	Count           *affected_count.Count
	Decision        *flag.Flag
	Addressed       bool
}
