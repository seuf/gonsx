package edgefirewall

import (
	"fmt"
	"github.com/sky-uk/gonsx/api"
	"log"
	"net/http"
)

// UpdateFirewallRuleAPI api object
type UpdateFirewallRuleAPI struct {
	*api.BaseAPI
}

// NewCreate returns a new object of UpdateFirewallRuleAPI.
func NewUpdateRule(edgeId string, ruleId int, rule *FirewallRule) *UpdateFirewallRuleAPI {
	log.Printf(fmt.Sprintf("[DEBUG] RULE TO UPDATE : %+v", rule))
	this := new(UpdateFirewallRuleAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodPut,
		fmt.Sprintf("/api/4.0/edges/%s/firewall/config/rules/%d", edgeId, ruleId),
		rule, new(string))

	return this
}

// GetResponse returns a ResponseObject of UpdateFirewallRuleAPI.
func (ca UpdateFirewallRuleAPI) GetResponse() string {
	return ca.ResponseObject().(string)
}
