// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20201201

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Redis_FirewallRule_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: redis cache firewall rule properties
	Properties *RedisFirewallRuleProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Redis_FirewallRule_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (rule Redis_FirewallRule_Spec_ARM) GetAPIVersion() string {
	return "2020-12-01"
}

// GetName returns the Name of the resource
func (rule *Redis_FirewallRule_Spec_ARM) GetName() string {
	return rule.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/firewallRules"
func (rule *Redis_FirewallRule_Spec_ARM) GetType() string {
	return "Microsoft.Cache/redis/firewallRules"
}

// Specifies a range of IP addresses permitted to connect to the cache
type RedisFirewallRuleProperties_ARM struct {
	// EndIP: highest IP address included in the range
	EndIP *string `json:"endIP,omitempty"`

	// StartIP: lowest IP address included in the range
	StartIP *string `json:"startIP,omitempty"`
}
