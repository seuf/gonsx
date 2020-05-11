package edgefirewall

import (
	"encoding/xml"
)

// const (
// 	Allow  Action = "allow"
// 	Block  Action = "block"
// 	Reject Action = "reject"
// )

type Source struct {
	Exclude          bool     `xml:"exclude"`
	IpAddress        string   `xml:"ipAddress,omitempty"`
	GroupingObjectId []string `xml:"groupingObjectId,omitempty"`
}

type Destination struct {
	Exclude          bool     `xml:"exclude"`
	IpAddress        string   `xml:"ipAddress,omitempty"`
	GroupingObjectId []string `xml:"groupingObjectId,omitempty"`
}

type Application struct {
	ApplicationId []string `xml:"applicationId,omitempty"`
}

type FirewallRule struct {
	XMLName        xml.Name    `xml:"firewallRule"`
	RuleId         int         `xml:"id,omitempty"`
	RuleTag        string      `xml:"ruleTag,omitempty"`
	Name           string      `xml:"name"`
	RuleType       string      `xml:"ruleType"`
	Enabled        bool        `xml:"enabled"`
	LoggingEnabled bool        `xml:"loggingEnabled"`
	Description    string      `xml:"description,omitempty"`
	Action         string      `xml:"action"`
	Source         Source      `xml:"source,omitempty"`
	Destination    Destination `xml:"destination,omitempty"`
	Application    Application `xml:"application,omitempty"`
}

type FirewallRules struct {
	XMLName      xml.Name       `xml:"firewallRules"`
	FirewallRule []FirewallRule `xml:"firewallRule"`
}

type Firewall struct {
	XMLName       xml.Name      `xml:"firewall"`
	FeatureType   string        `xml:"featureType"`
	Version       string        `xml:"version"`
	Enabled       bool          `xml:"enabled"`
	FirewallRules FirewallRules `xml:"firewallRules"`
}
