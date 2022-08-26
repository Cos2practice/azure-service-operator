// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of Redis_FirewallRules_Spec. Use v1beta20201201.Redis_FirewallRules_Spec instead
type Redis_FirewallRules_SpecARM struct {
	Location   *string                         `json:"location,omitempty"`
	Name       string                          `json:"name,omitempty"`
	Properties *RedisFirewallRulePropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string               `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Redis_FirewallRules_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (rules Redis_FirewallRules_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (rules *Redis_FirewallRules_SpecARM) GetName() string {
	return rules.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/firewallRules"
func (rules *Redis_FirewallRules_SpecARM) GetType() string {
	return "Microsoft.Cache/redis/firewallRules"
}

// Deprecated version of RedisFirewallRuleProperties. Use v1beta20201201.RedisFirewallRuleProperties instead
type RedisFirewallRulePropertiesARM struct {
	EndIP   *string `json:"endIP,omitempty"`
	StartIP *string `json:"startIP,omitempty"`
}