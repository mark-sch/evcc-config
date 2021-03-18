package templates 

import (
	"github.com/mark-sch/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "evsewifi",
		Name:   "EVSE-Wifi",
		Sample: `uri: http://192.0.2.2`,
	}

	registry.Add(template)
}
