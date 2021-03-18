package templates 

import (
	"github.com/mark-sch/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "nrgkick-bluetooth",
		Name:   "NRGKick BT (Bluetooth)",
		Sample: `macaddress: 00:99:22 # BT device MAC address
pin: # App PIN number (write protection)`,
	}

	registry.Add(template)
}
