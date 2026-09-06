package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"testing"
)

func TestClientConstantsMissingSiblingsFlagged(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/gotest/constant/constant.go",
		`package constant

const TokenEnvironment = "GOTEST_TOKEN"
const HostEnvironment = "GOTEST_HOST"
`,
	)
	c := scan.ClientConstants(v)
	assert.Integer(t, 2, len(c))
}

func TestClientConstantsSplitAcrossPackagesFlagged(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/gotest/constant/constant.go",
		`package constant

const TokenEnvironment = "GOTEST_TOKEN"
`,
	)
	v.WriteString(
		"pkg/tool/gotestd/constant/constant.go",
		`package constant

const HostEnvironment = "GOTEST_HOST"
const PortEnvironment = "GOTEST_PORT"
const InsecureEnvironment = "GOTEST_INSECURE"
`,
	)
	c := scan.ClientConstants(v)
	assert.Integer(t, 3, len(c))
}

func TestClientConstantsCompleteClean(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/gotest/constant/constant.go",
		`package constant

const (
	HostEnvironment     = "GOTEST_HOST"
	PortEnvironment     = "GOTEST_PORT"
	InsecureEnvironment = "GOTEST_INSECURE"
	TokenEnvironment    = "GOTEST_TOKEN"
)
`,
	)
	c := scan.ClientConstants(v)
	assert.Integer(t, 0, len(c))
}

func TestClientConstantsNoTokenIgnored(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/gotest/constant/constant.go",
		`package constant

const (
	HostEnvironment     = "GOTEST_HOST"
	PortEnvironment     = "GOTEST_PORT"
	InsecureEnvironment = "GOTEST_INSECURE"
)
`,
	)
	c := scan.ClientConstants(v)
	assert.Integer(t, 0, len(c))
}

func TestClientConstantsUpstreamTokenIgnored(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/gotest/constant/constant.go",
		`package constant

const TokenEnvironment = "RUNDECK_TOKEN"
`,
	)
	c := scan.ClientConstants(v)
	assert.Integer(t, 0, len(c))
}
