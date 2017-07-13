package fwrules

import (
	"fmt"
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// CreateFWRulesAPI default struct
type UpdateFWRulesAPI struct {
	*api.BaseAPI
}

// NewCreate - Returns all the rules in the specified context
func NewUpdate(updateRule Rule) *UpdateFWRulesAPI {
	this := new(CreateFWRulesAPI)
	var endpoint string
	switch updateRule.RuleType {
	case "LAYER3":
		endpoint = fmt.Sprintf("/api/4.0/firewall/globalroot-0/config/layer3sections/%d/rules", updateRule.SectionID)

	case "LAYER2":
		endpoint = fmt.Sprintf("/api/4.0/firewall/globalroot-0/config/layer2sections/%d/rules", updateRule.SectionID)

	}

	this.BaseAPI = api.NewBaseAPI(http.MethodPut, endpoint, updateRule, new(Rule))
	return this
}

// GetResponse - Returns ResponseObject from GetAllFirewallRulesAPI of Rule type.
func (updateAPI UpdateFWRulesAPI) GetResponse() Rule {
	return *updateAPI.ResponseObject().(*Rule)
}
