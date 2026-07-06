package status

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestParse(t *testing.T) {
	assert.Any(
		t,
		&Result{
			Active:      "active (running)",
			Loaded:      "loaded (/lib/systemd/system/nginx.service; enabled; vendor preset: enabled)",
			MainProcess: "10181",
			Memory:      "88.6M",
		},
		Parse(
			`● nginx.service - A high performance web server and a reverse proxy server
   Loaded: loaded (/lib/systemd/system/nginx.service; enabled; vendor preset: enabled)
   Active: active (running) since Wed 2024-11-06 12:43:45 CET; 1 months 6 days ago
     Docs: man:nginx(8)
  Process: 10179 ExecStartPre=/usr/sbin/nginx -t -q -g daemon on; master_process on; (code=exited, status=0/SUCCESS)
  Process: 10180 ExecStart=/usr/sbin/nginx -g daemon on; master_process on; (code=exited, status=0/SUCCESS)
 Main PID: 10181 (nginx)
    Tasks: 13 (limit: 4915)
   Memory: 88.6M
   CGroup: /system.slice/nginx.service
`,
		),
	)
}
