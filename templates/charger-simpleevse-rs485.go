package templates 

import (
	"github.com/mark-sch/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "simpleevse",
		Name:   "EVSE DIN (Modbus RTU)",
		Sample: `device: /dev/ttyUSB0 # serial RS485 interface`,
	}

	registry.Add(template)
}
