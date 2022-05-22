// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

// Deprecated version of RedisFirewallRule_Status. Use v1beta20201201.RedisFirewallRule_Status instead
type RedisFirewallRule_StatusARM struct {
	Id         *string                                `json:"id,omitempty"`
	Name       *string                                `json:"name,omitempty"`
	Properties *RedisFirewallRuleProperties_StatusARM `json:"properties,omitempty"`
	Type       *string                                `json:"type,omitempty"`
}

// Deprecated version of RedisFirewallRuleProperties_Status. Use v1beta20201201.RedisFirewallRuleProperties_Status instead
type RedisFirewallRuleProperties_StatusARM struct {
	EndIP   *string `json:"endIP,omitempty"`
	StartIP *string `json:"startIP,omitempty"`
}