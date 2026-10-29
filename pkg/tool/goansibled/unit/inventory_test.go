package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/unit/ansible_tester"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestApplyRunsAgainstStaticAndNetboxInventory(t *testing.T) {
	o := ansible_tester.New(
		t,
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				r *http.Request,
			) {
				if r.URL.Path != "/api/inventory/ansible" ||
					r.URL.Query().Get("field") != "ansible_playbook" ||
					r.Header.Get(webConstant.Authorization) !=
						join.Empty("Bearer ", constant.FixtureToken) {
					w.WriteHeader(http.StatusUnauthorized)

					return
				}

				_, e := w.Write([]byte(constant.FixtureMarkup))
				errors.PanicOnError(e)
			},
		),
	)
	r := o.WaitForRun()
	assert.String(t, "success", r.Status)
	assert.String(t, "bravo.yaml", r.Scope)
	arguments, okay := o.Captured(constant.FixtureArguments)
	assert.True(t, okay)
	a := strings.Split(strings.TrimSpace(arguments), "\n")
	assert.Count(t, 5, a)
	assert.String(t, "-i", a[0])
	assert.String(t, "/etc/ansible/hosts", a[1])
	assert.String(t, "-i", a[2])
	assert.True(t, strings.HasSuffix(a[3], "/netbox.yaml"))
	assert.String(t, "bravo.yaml", a[4])
	markup, okay := o.Captured(constant.FixtureInventory)
	assert.True(t, okay)
	assert.String(
		t,
		`all:
    children:
        bravo:
            hosts:
                charlie:
                    ansible_host: 192.168.0.5
`,
		markup,
	)
	_, e := os.Stat(a[3])
	assert.True(t, os.IsNotExist(e))
	assert.Count(t, 0, o.Reporter.Events())
}

func TestApplyFailsWhenNetboxInventoryFails(t *testing.T) {
	o := ansible_tester.New(
		t,
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				_ *http.Request,
			) {
				w.WriteHeader(http.StatusInternalServerError)
			},
		),
	)
	r := o.WaitForRun()
	assert.String(t, "error", r.Status)
	assert.String(t, "inventory", r.Scope)
	assert.String(
		t,
		"ansible inventory status: 500: 500 Internal Server Error",
		r.ErrorOutput,
	)
	_, ran := o.Captured(constant.FixtureArguments)
	assert.False(t, ran)
	assert.Count(t, 1, o.Reporter.Events())
}

func TestApplyWithoutNetboxUsesDefaultInventory(t *testing.T) {
	o := ansible_tester.New(t, nil)
	r := o.WaitForRun()
	assert.String(t, "success", r.Status)
	arguments, okay := o.Captured(constant.FixtureArguments)
	assert.True(t, okay)
	assert.String(t, "bravo.yaml\n", arguments)
}
