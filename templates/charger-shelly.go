package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "shelly",
		Name:   "Shelly",
		Sample: `uri: http://192.168.xxx.xxx  # shelly device ip address (local)
channel: 0  # shelly device relay channel 
standbypower: 15  # treat as charging above this power`,
	}

	registry.Add(template)
}
