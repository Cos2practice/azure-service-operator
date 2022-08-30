// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200601storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=eventgrid.azure.com,resources=domainstopics,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=eventgrid.azure.com,resources={domainstopics/status,domainstopics/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20200601.DomainsTopic
// Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/resourceDefinitions/domains_topics
type DomainsTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Domains_Topics_Spec `json:"spec,omitempty"`
	Status            DomainTopic_STATUS  `json:"status,omitempty"`
}

var _ conditions.Conditioner = &DomainsTopic{}

// GetConditions returns the conditions of the resource
func (topic *DomainsTopic) GetConditions() conditions.Conditions {
	return topic.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (topic *DomainsTopic) SetConditions(conditions conditions.Conditions) {
	topic.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &DomainsTopic{}

// AzureName returns the Azure name of the resource
func (topic *DomainsTopic) AzureName() string {
	return topic.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (topic DomainsTopic) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (topic *DomainsTopic) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (topic *DomainsTopic) GetSpec() genruntime.ConvertibleSpec {
	return &topic.Spec
}

// GetStatus returns the status of this resource
func (topic *DomainsTopic) GetStatus() genruntime.ConvertibleStatus {
	return &topic.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/domains/topics"
func (topic *DomainsTopic) GetType() string {
	return "Microsoft.EventGrid/domains/topics"
}

// NewEmptyStatus returns a new empty (blank) status
func (topic *DomainsTopic) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DomainTopic_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (topic *DomainsTopic) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(topic.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  topic.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (topic *DomainsTopic) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DomainTopic_STATUS); ok {
		topic.Status = *st
		return nil
	}

	// Convert status to required version
	var st DomainTopic_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	topic.Status = st
	return nil
}

// Hub marks that this DomainsTopic is the hub type for conversion
func (topic *DomainsTopic) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (topic *DomainsTopic) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: topic.Spec.OriginalVersion,
		Kind:    "DomainsTopic",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20200601.DomainsTopic
// Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/resourceDefinitions/domains_topics
type DomainsTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainsTopic `json:"items"`
}

// Storage version of v1beta20200601.Domains_Topics_Spec
type Domains_Topics_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string  `json:"azureName,omitempty"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a eventgrid.azure.com/Domain resource
	Owner       *genruntime.KnownResourceReference `group:"eventgrid.azure.com" json:"owner,omitempty" kind:"Domain"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Domains_Topics_Spec{}

// ConvertSpecFrom populates our Domains_Topics_Spec from the provided source
func (topics *Domains_Topics_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == topics {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(topics)
}

// ConvertSpecTo populates the provided destination from our Domains_Topics_Spec
func (topics *Domains_Topics_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == topics {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(topics)
}

// Storage version of v1beta20200601.DomainTopic_STATUS
type DomainTopic_STATUS struct {
	Conditions        []conditions.Condition `json:"conditions,omitempty"`
	Id                *string                `json:"id,omitempty"`
	Name              *string                `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	SystemData        *SystemData_STATUS     `json:"systemData,omitempty"`
	Type              *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DomainTopic_STATUS{}

// ConvertStatusFrom populates our DomainTopic_STATUS from the provided source
func (topic *DomainTopic_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == topic {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(topic)
}

// ConvertStatusTo populates the provided destination from our DomainTopic_STATUS
func (topic *DomainTopic_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == topic {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(topic)
}

func init() {
	SchemeBuilder.Register(&DomainsTopic{}, &DomainsTopicList{})
}