package templates 

import (
	"github.com/mark-sch/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "wallbe",
		Name:   "Wallbe (Eco, Pro)",
		Sample: `uri: 192.168.0.8:502 # TCP ModBus address`,
	}

	registry.Add(template)
}
