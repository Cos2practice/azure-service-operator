// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220801

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Service_Policy_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: Properties of the Policy.
	Properties *PolicyContractProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Service_Policy_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-08-01"
func (policy Service_Policy_Spec_ARM) GetAPIVersion() string {
	return "2022-08-01"
}

// GetName returns the Name of the resource
func (policy *Service_Policy_Spec_ARM) GetName() string {
	return policy.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/policies"
func (policy *Service_Policy_Spec_ARM) GetType() string {
	return "Microsoft.ApiManagement/service/policies"
}

// Policy contract Properties.
type PolicyContractProperties_ARM struct {
	// Format: Format of the policyContent.
	Format *PolicyContractProperties_Format_ARM `json:"format,omitempty"`

	// Value: Contents of the Policy as defined by the format.
	Value *string `json:"value,omitempty"`
}

// +kubebuilder:validation:Enum={"rawxml","rawxml-link","xml","xml-link"}
type PolicyContractProperties_Format_ARM string

const (
	PolicyContractProperties_Format_ARM_Rawxml     = PolicyContractProperties_Format_ARM("rawxml")
	PolicyContractProperties_Format_ARM_RawxmlLink = PolicyContractProperties_Format_ARM("rawxml-link")
	PolicyContractProperties_Format_ARM_Xml        = PolicyContractProperties_Format_ARM("xml")
	PolicyContractProperties_Format_ARM_XmlLink    = PolicyContractProperties_Format_ARM("xml-link")
)

// Mapping from string to PolicyContractProperties_Format_ARM
var policyContractProperties_Format_ARM_Values = map[string]PolicyContractProperties_Format_ARM{
	"rawxml":      PolicyContractProperties_Format_ARM_Rawxml,
	"rawxml-link": PolicyContractProperties_Format_ARM_RawxmlLink,
	"xml":         PolicyContractProperties_Format_ARM_Xml,
	"xml-link":    PolicyContractProperties_Format_ARM_XmlLink,
}
