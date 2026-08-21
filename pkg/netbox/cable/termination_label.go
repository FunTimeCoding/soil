package cable

import (
	"fmt"
	"github.com/netbox-community/go-netbox/v4"
)

func terminationLabel(terminations []netbox.GenericObject) string {
	if len(terminations) == 0 {
		return ""
	}

	t := terminations[0]
	o, okay := t.Object.(map[string]any)

	if !okay {
		return fmt.Sprintf("%s #%d", t.ObjectType, t.ObjectId)
	}

	name, _ := o["name"].(string)
	d, okay1 := o["device"].(map[string]any)

	if !okay1 {
		return name
	}

	deviceName, _ := d["name"].(string)

	return fmt.Sprintf("%s %s", deviceName, name)
}
