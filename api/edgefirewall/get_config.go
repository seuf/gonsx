package edgefirewall

import (
	"fmt"
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// CreateSecurityPolicyAPI api object
type GetEdgeFirewallConfig struct {
	*api.BaseAPI
}

// NewCreate returns a new object of CreatePolicyAPI.
func NewGetEdgeFirewallConfig(edgeId string) *GetEdgeFirewallConfig {
	this := new(GetEdgeFirewallConfig)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet,
		fmt.Sprintf("/api/4.0/edges/%s/firewall/config", edgeId),
		nil, new(Firewall))

	return this
}

// GetResponse returns a ResponseObject of CreateServiceAPI.
func (ca GetEdgeFirewallConfig) GetResponse() *Firewall {
	return ca.ResponseObject().(*Firewall)
}
