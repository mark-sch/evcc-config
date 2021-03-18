package templates 

import (
	"github.com/mark-sch/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "vehicle",
		Type:   "porsche",
		Name:   "Porsche",
		Sample: `title: Taycan # display name for UI
capacity: 83 # kWh
user: # user
password: # password
vin: WP...`,
	}

	registry.Add(template)
}
