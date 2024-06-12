// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220401

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type RoleDefinition_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: Role definition properties.
	Properties *RoleDefinitionProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &RoleDefinition_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-04-01"
func (definition RoleDefinition_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (definition *RoleDefinition_Spec_ARM) GetName() string {
	return definition.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Authorization/roleDefinitions"
func (definition *RoleDefinition_Spec_ARM) GetType() string {
	return "Microsoft.Authorization/roleDefinitions"
}

// Role definition properties.
type RoleDefinitionProperties_ARM struct {
	AssignableScopes []string `json:"assignableScopes,omitempty"`

	// Description: The role definition description.
	Description *string `json:"description,omitempty"`

	// Permissions: Role definition permissions.
	Permissions []Permission_ARM `json:"permissions,omitempty"`

	// RoleName: The role name.
	RoleName *string `json:"roleName,omitempty"`

	// Type: The role type.
	Type *string `json:"type,omitempty"`
}

// Role definition permissions.
type Permission_ARM struct {
	// Actions: Allowed actions.
	Actions []string `json:"actions,omitempty"`

	// DataActions: Allowed Data actions.
	DataActions []string `json:"dataActions,omitempty"`

	// NotActions: Denied actions.
	NotActions []string `json:"notActions,omitempty"`

	// NotDataActions: Denied Data actions.
	NotDataActions []string `json:"notDataActions,omitempty"`
}
