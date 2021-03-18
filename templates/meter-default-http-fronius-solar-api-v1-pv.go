package templates 

import (
	"github.com/mark-sch/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "default",
		Name:   "Fronius Solar API V1 (PV Meter/ HTTP)",
		Sample: `power:
  type: http
  uri: http://192.0.2.2/solar_api/v1/GetPowerFlowRealtimeData.fcgi
  jq: if .Body.Data.Site.P_PV == null then 0 else .Body.Data.Site.P_PV end`,
	}

	registry.Add(template)
}
