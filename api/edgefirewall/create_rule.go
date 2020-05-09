package edgefirewall

import (
	"fmt"
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// CreateEdgeFirewallRuleAPI api object
type CreateEdgeFirewallRuleAPI struct {
	*api.BaseAPI
}

// NewCreate returns a new object of CreateEdgeFirewallRuleAPI.
func NewCreateRule(edgeId string, rules *FirewallRules) *CreateEdgeFirewallRuleAPI {
	this := new(CreateEdgeFirewallRuleAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodPost,
		fmt.Sprintf("/api/4.0/edges/%s/firewall/config/rules", edgeId),
		rules, new(string))

	return this
}

// GetResponse returns a ResponseObject of CreateServiceAPI.
func (ca CreateEdgeFirewallRuleAPI) GetResponse() string {
	return ca.ResponseObject().(string)
}
