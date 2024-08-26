// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210701

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Workspaces_Connection_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: Properties of workspace connection.
	Properties *WorkspaceConnectionProps_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Workspaces_Connection_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-07-01"
func (connection Workspaces_Connection_Spec_ARM) GetAPIVersion() string {
	return "2021-07-01"
}

// GetName returns the Name of the resource
func (connection *Workspaces_Connection_Spec_ARM) GetName() string {
	return connection.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.MachineLearningServices/workspaces/connections"
func (connection *Workspaces_Connection_Spec_ARM) GetType() string {
	return "Microsoft.MachineLearningServices/workspaces/connections"
}

// Workspace Connection specific properties.
type WorkspaceConnectionProps_ARM struct {
	// AuthType: Authorization type of the workspace connection.
	AuthType *string `json:"authType,omitempty"`

	// Category: Category of the workspace connection.
	Category *string `json:"category,omitempty"`

	// Target: Target of the workspace connection.
	Target *string `json:"target,omitempty"`

	// Value: Value details of the workspace connection.
	Value *string `json:"value,omitempty"`

	// ValueFormat: format for the workspace connection value
	ValueFormat *WorkspaceConnectionProps_ValueFormat_ARM `json:"valueFormat,omitempty"`
}

// +kubebuilder:validation:Enum={"JSON"}
type WorkspaceConnectionProps_ValueFormat_ARM string

const WorkspaceConnectionProps_ValueFormat_ARM_JSON = WorkspaceConnectionProps_ValueFormat_ARM("JSON")

// Mapping from string to WorkspaceConnectionProps_ValueFormat_ARM
var workspaceConnectionProps_ValueFormat_ARM_Values = map[string]WorkspaceConnectionProps_ValueFormat_ARM{
	"json": WorkspaceConnectionProps_ValueFormat_ARM_JSON,
}
