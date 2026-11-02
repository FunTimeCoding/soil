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
