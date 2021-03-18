package templates 

import (
	"github.com/mark-sch/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "tesla",
		Name:   "Tesla Powerwall (PV Meter)",
		Sample: `uri: http://192.0.2.2/
usage: pv`,
	}

	registry.Add(template)
}
