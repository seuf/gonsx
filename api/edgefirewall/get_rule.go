package edgefirewall

import (
	"fmt"
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// CreateSecurityPolicyAPI api object
type GetFirewallRuleAPI struct {
	*api.BaseAPI
}

// NewCreate returns a new object of CreatePolicyAPI.
func NewGetRule(edgeId string, ruleId int) *GetFirewallRuleAPI {
	this := new(GetFirewallRuleAPI)

	this.BaseAPI = api.NewBaseAPI(http.MethodGet,
		fmt.Sprintf("/api/4.0/edges/%s/firewall/config/rules/%d", edgeId, ruleId),
		nil, new(FirewallRule))

	return this
}

// GetResponse returns a ResponseObject of CreateServiceAPI.
func (ca GetFirewallRuleAPI) GetResponse() *FirewallRule {
	return ca.ResponseObject().(*FirewallRule)
}
