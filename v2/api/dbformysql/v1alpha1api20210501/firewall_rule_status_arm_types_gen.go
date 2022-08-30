// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210501

// Deprecated version of FirewallRule_STATUS. Use v1beta20210501.FirewallRule_STATUS instead
type FirewallRule_STATUSARM struct {
	Id         *string                           `json:"id,omitempty"`
	Name       *string                           `json:"name,omitempty"`
	Properties *FirewallRuleProperties_STATUSARM `json:"properties,omitempty"`
	SystemData *SystemData_STATUSARM             `json:"systemData,omitempty"`
	Type       *string                           `json:"type,omitempty"`
}

// Deprecated version of FirewallRuleProperties_STATUS. Use v1beta20210501.FirewallRuleProperties_STATUS instead
type FirewallRuleProperties_STATUSARM struct {
	EndIpAddress   *string `json:"endIpAddress,omitempty"`
	StartIpAddress *string `json:"startIpAddress,omitempty"`
}