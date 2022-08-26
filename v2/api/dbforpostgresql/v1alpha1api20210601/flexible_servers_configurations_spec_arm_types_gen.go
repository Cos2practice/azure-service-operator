// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of FlexibleServers_Configurations_Spec. Use v1beta20210601.FlexibleServers_Configurations_Spec instead
type FlexibleServers_Configurations_SpecARM struct {
	Location   *string                     `json:"location,omitempty"`
	Name       string                      `json:"name,omitempty"`
	Properties *ConfigurationPropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string           `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServers_Configurations_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (configurations FlexibleServers_Configurations_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (configurations *FlexibleServers_Configurations_SpecARM) GetName() string {
	return configurations.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers/configurations"
func (configurations *FlexibleServers_Configurations_SpecARM) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers/configurations"
}

// Deprecated version of ConfigurationProperties. Use v1beta20210601.ConfigurationProperties instead
type ConfigurationPropertiesARM struct {
	Source *string `json:"source,omitempty"`
	Value  *string `json:"value,omitempty"`
}