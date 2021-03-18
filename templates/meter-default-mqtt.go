package templates 

import (
	"github.com/mark-sch/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "default",
		Name:   "Generic (MQTT)",
		Sample: `power: # power reading
  type: mqtt # use mqtt plugin
  topic: mbmd/sdm1-1/Power # mqtt topic
  timeout: 10s # don't use older values`,
	}

	registry.Add(template)
}
