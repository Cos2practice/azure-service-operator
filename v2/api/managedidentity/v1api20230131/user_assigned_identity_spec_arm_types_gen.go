// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230131

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type UserAssignedIdentity_Spec_ARM struct {
	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &UserAssignedIdentity_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-01-31"
func (identity UserAssignedIdentity_Spec_ARM) GetAPIVersion() string {
	return "2023-01-31"
}

// GetName returns the Name of the resource
func (identity *UserAssignedIdentity_Spec_ARM) GetName() string {
	return identity.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ManagedIdentity/userAssignedIdentities"
func (identity *UserAssignedIdentity_Spec_ARM) GetType() string {
	return "Microsoft.ManagedIdentity/userAssignedIdentities"
}
