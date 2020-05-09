package edgefirewall

import (
	"fmt"
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// CreateSecurityPolicyAPI api object
type DeleteEdgeFirewallRuleAPI struct {
	*api.BaseAPI
}

// NewCreate returns a new object of CreatePolicyAPI.
func NewDeleteRule(edgeId string, ruleId int) *DeleteEdgeFirewallRuleAPI {
	this := new(DeleteEdgeFirewallRuleAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodDelete, fmt.Sprintf(
		"/api/4.0/edges/%s/firewall/config/rules/%d", edgeId, ruleId,
	), nil, nil)

	return this
}

// GetResponse returns a ResponseObject of CreateServiceAPI.
func (ca DeleteEdgeFirewallRuleAPI) GetResponse() string {
	return ca.ResponseObject().(string)
}
