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
[ -e "$GOANSIBLED_CAPTURE/fail" ] && exit 2
exit 0
`
	FixtureHold   = "hold"
	FixtureScript = `#!/bin/sh
while [ -e "$GOANSIBLED_CAPTURE/hold" ]; do sleep 0.05; done
printf '%s\n' "$@" > "$GOANSIBLED_CAPTURE/arguments"
for a in "$@"; do
	case "$a" in
	*/netbox.yaml) cat "$a" > "$GOANSIBLED_CAPTURE/inventory" ;;
	esac
done
`
)
