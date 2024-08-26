// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=sql.azure.com,resources=serversconnectionpolicies,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sql.azure.com,resources={serversconnectionpolicies/status,serversconnectionpolicies/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20211101.ServersConnectionPolicy
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/ServerConnectionPolicies.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/connectionPolicies/default
type ServersConnectionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Servers_ConnectionPolicy_Spec   `json:"spec,omitempty"`
	Status            Servers_ConnectionPolicy_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ServersConnectionPolicy{}

// GetConditions returns the conditions of the resource
func (policy *ServersConnectionPolicy) GetConditions() conditions.Conditions {
	return policy.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (policy *ServersConnectionPolicy) SetConditions(conditions conditions.Conditions) {
	policy.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &ServersConnectionPolicy{}

// AzureName returns the Azure name of the resource (always "default")
func (policy *ServersConnectionPolicy) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (policy ServersConnectionPolicy) GetAPIVersion() string {
	return "2021-11-01"
}

// GetResourceScope returns the scope of the resource
func (policy *ServersConnectionPolicy) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (policy *ServersConnectionPolicy) GetSpec() genruntime.ConvertibleSpec {
	return &policy.Spec
}

// GetStatus returns the status of this resource
func (policy *ServersConnectionPolicy) GetStatus() genruntime.ConvertibleStatus {
	return &policy.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (policy *ServersConnectionPolicy) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/connectionPolicies"
func (policy *ServersConnectionPolicy) GetType() string {
	return "Microsoft.Sql/servers/connectionPolicies"
}

// NewEmptyStatus returns a new empty (blank) status
func (policy *ServersConnectionPolicy) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Servers_ConnectionPolicy_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (policy *ServersConnectionPolicy) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(policy.Spec)
	return policy.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (policy *ServersConnectionPolicy) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Servers_ConnectionPolicy_STATUS); ok {
		policy.Status = *st
		return nil
	}

	// Convert status to required version
	var st Servers_ConnectionPolicy_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	policy.Status = st
	return nil
}

// Hub marks that this ServersConnectionPolicy is the hub type for conversion
func (policy *ServersConnectionPolicy) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (policy *ServersConnectionPolicy) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: policy.Spec.OriginalVersion,
		Kind:    "ServersConnectionPolicy",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20211101.ServersConnectionPolicy
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/ServerConnectionPolicies.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/connectionPolicies/default
type ServersConnectionPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServersConnectionPolicy `json:"items"`
}

// Storage version of v1api20211101.Servers_ConnectionPolicy_Spec
type Servers_ConnectionPolicy_Spec struct {
	ConnectionType  *string `json:"connectionType,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a sql.azure.com/Server resource
	Owner       *genruntime.KnownResourceReference `group:"sql.azure.com" json:"owner,omitempty" kind:"Server"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Servers_ConnectionPolicy_Spec{}

// ConvertSpecFrom populates our Servers_ConnectionPolicy_Spec from the provided source
func (policy *Servers_ConnectionPolicy_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == policy {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(policy)
}

// ConvertSpecTo populates the provided destination from our Servers_ConnectionPolicy_Spec
func (policy *Servers_ConnectionPolicy_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == policy {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(policy)
}

// Storage version of v1api20211101.Servers_ConnectionPolicy_STATUS
type Servers_ConnectionPolicy_STATUS struct {
	Conditions     []conditions.Condition `json:"conditions,omitempty"`
	ConnectionType *string                `json:"connectionType,omitempty"`
	Id             *string                `json:"id,omitempty"`
	Kind           *string                `json:"kind,omitempty"`
	Location       *string                `json:"location,omitempty"`
	Name           *string                `json:"name,omitempty"`
	PropertyBag    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type           *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Servers_ConnectionPolicy_STATUS{}

// ConvertStatusFrom populates our Servers_ConnectionPolicy_STATUS from the provided source
func (policy *Servers_ConnectionPolicy_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == policy {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(policy)
}

// ConvertStatusTo populates the provided destination from our Servers_ConnectionPolicy_STATUS
func (policy *Servers_ConnectionPolicy_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == policy {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(policy)
}

func init() {
	SchemeBuilder.Register(&ServersConnectionPolicy{}, &ServersConnectionPolicyList{})
}
