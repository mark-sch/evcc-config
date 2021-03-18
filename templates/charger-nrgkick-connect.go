package templates 

import (
	"github.com/mark-sch/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "nrgkick-connect",
		Name:   "NRGKick Connect",
		Sample: `uri: http://192.0.2.2
mac: 00:99:22 # BT device MAC address
password: # password`,
	}

	registry.Add(template)
}
