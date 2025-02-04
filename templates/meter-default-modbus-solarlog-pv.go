package templates 

import (
	"github.com/mark-sch/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "custom",
		Name:   "Solarlog (PV Meter)",
		Sample: `power:
  source: modbus
  uri: 192.0.2.2:502 # IP address of the SolarLog  device and ModBus port address
  id: 1
  register:
    address: 3502
    type: input
    decode: uint32s`,
	}

	registry.Add(template)
}
