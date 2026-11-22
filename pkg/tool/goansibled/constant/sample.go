package constant

const (
	FixtureCaptureEnvironment = "GOANSIBLED_CAPTURE"
	FixtureToken              = "alfa-token"
	FixturePlaybook           = "bravo.yaml"
	FixtureArguments          = "arguments"
	FixtureInventory          = "inventory"
	FixtureMarkup             = `all:
    children:
        bravo:
            hosts:
                charlie:
                    ansible_host: 192.168.0.5
`
	FixtureHostCommand = `alfa:
  group: bravo
  module: ansible.builtin.systemd_service
  arguments: name=echo state=restarted
`
	FixtureAdHoc       = "ad-hoc"
	FixtureOutside     = "outside"
	FixtureFail        = "fail"
	FixtureCallback    = "callback"
	FixtureAdHocScript = `#!/bin/sh
case "$*" in
*--list-hosts*)
	echo "  hosts (0):"
	[ -e "$GOANSIBLED_CAPTURE/outside" ] && exit 0
	for a in "$@"; do
		case "$a" in
		*:\&*) echo "    ${a#*&}" ;;
		esac
	done
	exit 0
	;;
esac
printf '%s\n' "$@" > "$GOANSIBLED_CAPTURE/ad-hoc"
printf '%s %s\n' "$ANSIBLE_STDOUT_CALLBACK" "$ANSIBLE_LOAD_CALLBACK_PLUGINS" > "$GOANSIBLED_CAPTURE/callback"
host=
for a in "$@"; do
	case "$a" in
	*:\&*) host="${a#*&}" ;;
	esac
done
if [ -e "$GOANSIBLED_CAPTURE/fail" ]; then
	printf '{"plays":[{"tasks":[{"hosts":{"%s":{"rc":2,"stdout":"","stderr":"echo failed","failed":true}}}]}]}\n' "$host"
	exit 2
fi
if [ -n "$host" ]; then
	printf '{"plays":[{"tasks":[{"hosts":{"%s":{"rc":0,"stdout":"echo restarted","stderr":""}}}]}]}\n' "$host"
	exit 0
fi
printf '{"plays":[{"tasks":[{"hosts":{"delta":{"unreachable":true,"msg":"delta down"},"charlie":{"rc":0,"stdout":"echo restarted","stderr":""}}}]}]}\n'
exit 4
`
	FixtureCallbackReport = `{"custom_stats":{},"global_custom_stats":{},"plays":[{"play":{"duration":{"end":"2026-11-22T20:34:35.374576Z","start":"2026-11-22T20:34:25.324912Z"},"id":"02000000-0000-580f-f01a-000000000004","name":"Ansible Ad-Hoc","path":""},"tasks":[{"hosts":{"bravo":{"action":"ansible.builtin.shell","changed":false,"exception":"(traceback unavailable)","msg":"Task failed: Failed to connect to the host via ssh: ssh: connect to host 192.0.2.1 port 22: Operation timed out","unreachable":true},"alfa":{"action":"ansible.builtin.shell","ansible_facts":{"discovered_interpreter_python":"/usr/bin/python3.13"},"changed":true,"cmd":"echo hello; echo oops >&2; exit 3","delta":"0:00:00.002959","end":"2026-11-22 21:34:26.028255","exception":"(traceback unavailable)","failed":true,"msg":"The command exited with a non-zero return code.","rc":3,"start":"2026-11-22 21:34:26.025296","stderr":"oops","stderr_lines":["oops"],"stdout":"hello","stdout_lines":["hello"],"warnings":["Host 'alfa' is using the discovered Python interpreter at '/usr/bin/python3.13', but future installation of another Python interpreter could cause a different interpreter to be discovered."]}},"task":{"duration":{"end":"2026-11-22T20:34:35.374576Z","start":"2026-11-22T20:34:25.331110Z"},"id":"02000000-0000-580f-f01a-000000000006","name":"ansible.builtin.shell","path":""}}]}],"stats":{"bravo":{"changed":0,"failures":0,"ignored":0,"ok":0,"rescued":0,"skipped":0,"unreachable":1},"alfa":{"changed":0,"failures":1,"ignored":0,"ok":0,"rescued":0,"skipped":0,"unreachable":0}}}`
	FixtureHold           = "hold"
	FixtureScript         = `#!/bin/sh
while [ -e "$GOANSIBLED_CAPTURE/hold" ]; do sleep 0.05; done
printf '%s\n' "$@" > "$GOANSIBLED_CAPTURE/arguments"
for a in "$@"; do
	case "$a" in
	*/netbox.yaml) cat "$a" > "$GOANSIBLED_CAPTURE/inventory" ;;
	esac
done
`
)
