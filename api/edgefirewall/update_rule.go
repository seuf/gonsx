package edgefirewall

import (
	"fmt"
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// CreateSecurityPolicyAPI api object
type UpdateFirewallRuleAPI struct {
	*api.BaseAPI
}

// NewCreate returns a new object of CreatePolicyAPI.
func NewUpdateRule(edgeId string, ruleId int, rules *FirewallRules) *UpdateFirewallRuleAPI {
	this := new(UpdateFirewallRuleAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodPut,
		fmt.Sprintf("/api/4.0/edges/%s/firewall/config/rules/%d", edgeId, ruleId),
		rules, new(FirewallRules))
	//this.SetRequestHeader("If-Match", etag)

	return this
}

// GetResponse returns a ResponseObject of CreateServiceAPI.
func (ca UpdateFirewallRuleAPI) GetResponse() string {
	return ca.ResponseObject().(string)
}
