// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Fleets_Member_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: The resource-specific properties for this resource.
	Properties *FleetMemberProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Fleets_Member_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-03-15-preview"
func (member Fleets_Member_Spec) GetAPIVersion() string {
	return "2023-03-15-preview"
}

// GetName returns the Name of the resource
func (member *Fleets_Member_Spec) GetName() string {
	return member.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ContainerService/fleets/members"
func (member *Fleets_Member_Spec) GetType() string {
	return "Microsoft.ContainerService/fleets/members"
}

// A member of the Fleet. It contains a reference to an existing Kubernetes cluster on Azure.
type FleetMemberProperties struct {
	ClusterResourceId *string `json:"clusterResourceId,omitempty"`

	// Group: The group this member belongs to for multi-cluster update management.
	Group *string `json:"group,omitempty"`
}
