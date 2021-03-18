package templates 

import (
	"github.com/mark-sch/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "Generic SunSpec PV inverter (PV Meter)",
		Sample: `uri: 192.0.2.2:502
id: 1`,
	}

	registry.Add(template)
}
