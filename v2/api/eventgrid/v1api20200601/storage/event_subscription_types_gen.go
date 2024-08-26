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

// +kubebuilder:rbac:groups=eventgrid.azure.com,resources=eventsubscriptions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=eventgrid.azure.com,resources={eventsubscriptions/status,eventsubscriptions/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20200601.EventSubscription
// Generator information:
// - Generated from: /eventgrid/resource-manager/Microsoft.EventGrid/stable/2020-06-01/EventGrid.json
// - ARM URI: /{scope}/providers/Microsoft.EventGrid/eventSubscriptions/{eventSubscriptionName}
type EventSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventSubscription_Spec   `json:"spec,omitempty"`
	Status            EventSubscription_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &EventSubscription{}

// GetConditions returns the conditions of the resource
func (subscription *EventSubscription) GetConditions() conditions.Conditions {
	return subscription.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (subscription *EventSubscription) SetConditions(conditions conditions.Conditions) {
	subscription.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &EventSubscription{}

// AzureName returns the Azure name of the resource
func (subscription *EventSubscription) AzureName() string {
	return subscription.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (subscription EventSubscription) GetAPIVersion() string {
	return "2020-06-01"
}

// GetResourceScope returns the scope of the resource
func (subscription *EventSubscription) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeExtension
}

// GetSpec returns the specification of this resource
func (subscription *EventSubscription) GetSpec() genruntime.ConvertibleSpec {
	return &subscription.Spec
}

// GetStatus returns the status of this resource
func (subscription *EventSubscription) GetStatus() genruntime.ConvertibleStatus {
	return &subscription.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (subscription *EventSubscription) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/eventSubscriptions"
func (subscription *EventSubscription) GetType() string {
	return "Microsoft.EventGrid/eventSubscriptions"
}

// NewEmptyStatus returns a new empty (blank) status
func (subscription *EventSubscription) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &EventSubscription_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (subscription *EventSubscription) Owner() *genruntime.ResourceReference {
	return subscription.Spec.Owner.AsResourceReference()
}

// SetStatus sets the status of this resource
func (subscription *EventSubscription) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*EventSubscription_STATUS); ok {
		subscription.Status = *st
		return nil
	}

	// Convert status to required version
	var st EventSubscription_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	subscription.Status = st
	return nil
}

// Hub marks that this EventSubscription is the hub type for conversion
func (subscription *EventSubscription) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (subscription *EventSubscription) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: subscription.Spec.OriginalVersion,
		Kind:    "EventSubscription",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20200601.EventSubscription
// Generator information:
// - Generated from: /eventgrid/resource-manager/Microsoft.EventGrid/stable/2020-06-01/EventGrid.json
// - ARM URI: /{scope}/providers/Microsoft.EventGrid/eventSubscriptions/{eventSubscriptionName}
type EventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventSubscription `json:"items"`
}

// Storage version of v1api20200601.EventSubscription_Spec
type EventSubscription_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName             string                        `json:"azureName,omitempty"`
	DeadLetterDestination *DeadLetterDestination        `json:"deadLetterDestination,omitempty"`
	Destination           *EventSubscriptionDestination `json:"destination,omitempty"`
	EventDeliverySchema   *string                       `json:"eventDeliverySchema,omitempty"`
	ExpirationTimeUtc     *string                       `json:"expirationTimeUtc,omitempty"`
	Filter                *EventSubscriptionFilter      `json:"filter,omitempty"`
	Labels                []string                      `json:"labels,omitempty"`
	OriginalVersion       string                        `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. This resource is an
	// extension resource, which means that any other Azure resource can be its owner.
	Owner       *genruntime.ArbitraryOwnerReference `json:"owner,omitempty"`
	PropertyBag genruntime.PropertyBag              `json:"$propertyBag,omitempty"`
	RetryPolicy *RetryPolicy                        `json:"retryPolicy,omitempty"`
}

var _ genruntime.ConvertibleSpec = &EventSubscription_Spec{}

// ConvertSpecFrom populates our EventSubscription_Spec from the provided source
func (subscription *EventSubscription_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == subscription {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(subscription)
}

// ConvertSpecTo populates the provided destination from our EventSubscription_Spec
func (subscription *EventSubscription_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == subscription {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(subscription)
}

// Storage version of v1api20200601.EventSubscription_STATUS
// Event Subscription
type EventSubscription_STATUS struct {
	Conditions            []conditions.Condition               `json:"conditions,omitempty"`
	DeadLetterDestination *DeadLetterDestination_STATUS        `json:"deadLetterDestination,omitempty"`
	Destination           *EventSubscriptionDestination_STATUS `json:"destination,omitempty"`
	EventDeliverySchema   *string                              `json:"eventDeliverySchema,omitempty"`
	ExpirationTimeUtc     *string                              `json:"expirationTimeUtc,omitempty"`
	Filter                *EventSubscriptionFilter_STATUS      `json:"filter,omitempty"`
	Id                    *string                              `json:"id,omitempty"`
	Labels                []string                             `json:"labels,omitempty"`
	Name                  *string                              `json:"name,omitempty"`
	PropertyBag           genruntime.PropertyBag               `json:"$propertyBag,omitempty"`
	ProvisioningState     *string                              `json:"provisioningState,omitempty"`
	RetryPolicy           *RetryPolicy_STATUS                  `json:"retryPolicy,omitempty"`
	SystemData            *SystemData_STATUS                   `json:"systemData,omitempty"`
	Topic                 *string                              `json:"topic,omitempty"`
	Type                  *string                              `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &EventSubscription_STATUS{}

// ConvertStatusFrom populates our EventSubscription_STATUS from the provided source
func (subscription *EventSubscription_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == subscription {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(subscription)
}

// ConvertStatusTo populates the provided destination from our EventSubscription_STATUS
func (subscription *EventSubscription_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == subscription {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(subscription)
}

// Storage version of v1api20200601.DeadLetterDestination
type DeadLetterDestination struct {
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	StorageBlob *StorageBlobDeadLetterDestination `json:"storageBlob,omitempty"`
}

// Storage version of v1api20200601.DeadLetterDestination_STATUS
type DeadLetterDestination_STATUS struct {
	PropertyBag genruntime.PropertyBag                   `json:"$propertyBag,omitempty"`
	StorageBlob *StorageBlobDeadLetterDestination_STATUS `json:"storageBlob,omitempty"`
}

// Storage version of v1api20200601.EventSubscriptionDestination
type EventSubscriptionDestination struct {
	AzureFunction    *AzureFunctionEventSubscriptionDestination    `json:"azureFunction,omitempty"`
	EventHub         *EventHubEventSubscriptionDestination         `json:"eventHub,omitempty"`
	HybridConnection *HybridConnectionEventSubscriptionDestination `json:"hybridConnection,omitempty"`
	PropertyBag      genruntime.PropertyBag                        `json:"$propertyBag,omitempty"`
	ServiceBusQueue  *ServiceBusQueueEventSubscriptionDestination  `json:"serviceBusQueue,omitempty"`
	ServiceBusTopic  *ServiceBusTopicEventSubscriptionDestination  `json:"serviceBusTopic,omitempty"`
	StorageQueue     *StorageQueueEventSubscriptionDestination     `json:"storageQueue,omitempty"`
	WebHook          *WebHookEventSubscriptionDestination          `json:"webHook,omitempty"`
}

// Storage version of v1api20200601.EventSubscriptionDestination_STATUS
type EventSubscriptionDestination_STATUS struct {
	AzureFunction    *AzureFunctionEventSubscriptionDestination_STATUS    `json:"azureFunction,omitempty"`
	EventHub         *EventHubEventSubscriptionDestination_STATUS         `json:"eventHub,omitempty"`
	HybridConnection *HybridConnectionEventSubscriptionDestination_STATUS `json:"hybridConnection,omitempty"`
	PropertyBag      genruntime.PropertyBag                               `json:"$propertyBag,omitempty"`
	ServiceBusQueue  *ServiceBusQueueEventSubscriptionDestination_STATUS  `json:"serviceBusQueue,omitempty"`
	ServiceBusTopic  *ServiceBusTopicEventSubscriptionDestination_STATUS  `json:"serviceBusTopic,omitempty"`
	StorageQueue     *StorageQueueEventSubscriptionDestination_STATUS     `json:"storageQueue,omitempty"`
	WebHook          *WebHookEventSubscriptionDestination_STATUS          `json:"webHook,omitempty"`
}

// Storage version of v1api20200601.EventSubscriptionFilter
// Filter for the Event Subscription.
type EventSubscriptionFilter struct {
	AdvancedFilters        []AdvancedFilter       `json:"advancedFilters,omitempty"`
	IncludedEventTypes     []string               `json:"includedEventTypes,omitempty"`
	IsSubjectCaseSensitive *bool                  `json:"isSubjectCaseSensitive,omitempty"`
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SubjectBeginsWith      *string                `json:"subjectBeginsWith,omitempty"`
	SubjectEndsWith        *string                `json:"subjectEndsWith,omitempty"`
}

// Storage version of v1api20200601.EventSubscriptionFilter_STATUS
// Filter for the Event Subscription.
type EventSubscriptionFilter_STATUS struct {
	AdvancedFilters        []AdvancedFilter_STATUS `json:"advancedFilters,omitempty"`
	IncludedEventTypes     []string                `json:"includedEventTypes,omitempty"`
	IsSubjectCaseSensitive *bool                   `json:"isSubjectCaseSensitive,omitempty"`
	PropertyBag            genruntime.PropertyBag  `json:"$propertyBag,omitempty"`
	SubjectBeginsWith      *string                 `json:"subjectBeginsWith,omitempty"`
	SubjectEndsWith        *string                 `json:"subjectEndsWith,omitempty"`
}

// Storage version of v1api20200601.RetryPolicy
// Information about the retry policy for an event subscription.
type RetryPolicy struct {
	EventTimeToLiveInMinutes *int                   `json:"eventTimeToLiveInMinutes,omitempty"`
	MaxDeliveryAttempts      *int                   `json:"maxDeliveryAttempts,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20200601.RetryPolicy_STATUS
// Information about the retry policy for an event subscription.
type RetryPolicy_STATUS struct {
	EventTimeToLiveInMinutes *int                   `json:"eventTimeToLiveInMinutes,omitempty"`
	MaxDeliveryAttempts      *int                   `json:"maxDeliveryAttempts,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20200601.AdvancedFilter
type AdvancedFilter struct {
	BoolEquals                *BoolEqualsAdvancedFilter                `json:"boolEquals,omitempty"`
	NumberGreaterThan         *NumberGreaterThanAdvancedFilter         `json:"numberGreaterThan,omitempty"`
	NumberGreaterThanOrEquals *NumberGreaterThanOrEqualsAdvancedFilter `json:"numberGreaterThanOrEquals,omitempty"`
	NumberIn                  *NumberInAdvancedFilter                  `json:"numberIn,omitempty"`
	NumberLessThan            *NumberLessThanAdvancedFilter            `json:"numberLessThan,omitempty"`
	NumberLessThanOrEquals    *NumberLessThanOrEqualsAdvancedFilter    `json:"numberLessThanOrEquals,omitempty"`
	NumberNotIn               *NumberNotInAdvancedFilter               `json:"numberNotIn,omitempty"`
	PropertyBag               genruntime.PropertyBag                   `json:"$propertyBag,omitempty"`
	StringBeginsWith          *StringBeginsWithAdvancedFilter          `json:"stringBeginsWith,omitempty"`
	StringContains            *StringContainsAdvancedFilter            `json:"stringContains,omitempty"`
	StringEndsWith            *StringEndsWithAdvancedFilter            `json:"stringEndsWith,omitempty"`
	StringIn                  *StringInAdvancedFilter                  `json:"stringIn,omitempty"`
	StringNotIn               *StringNotInAdvancedFilter               `json:"stringNotIn,omitempty"`
}

// Storage version of v1api20200601.AdvancedFilter_STATUS
type AdvancedFilter_STATUS struct {
	BoolEquals                *BoolEqualsAdvancedFilter_STATUS                `json:"boolEquals,omitempty"`
	NumberGreaterThan         *NumberGreaterThanAdvancedFilter_STATUS         `json:"numberGreaterThan,omitempty"`
	NumberGreaterThanOrEquals *NumberGreaterThanOrEqualsAdvancedFilter_STATUS `json:"numberGreaterThanOrEquals,omitempty"`
	NumberIn                  *NumberInAdvancedFilter_STATUS                  `json:"numberIn,omitempty"`
	NumberLessThan            *NumberLessThanAdvancedFilter_STATUS            `json:"numberLessThan,omitempty"`
	NumberLessThanOrEquals    *NumberLessThanOrEqualsAdvancedFilter_STATUS    `json:"numberLessThanOrEquals,omitempty"`
	NumberNotIn               *NumberNotInAdvancedFilter_STATUS               `json:"numberNotIn,omitempty"`
	PropertyBag               genruntime.PropertyBag                          `json:"$propertyBag,omitempty"`
	StringBeginsWith          *StringBeginsWithAdvancedFilter_STATUS          `json:"stringBeginsWith,omitempty"`
	StringContains            *StringContainsAdvancedFilter_STATUS            `json:"stringContains,omitempty"`
	StringEndsWith            *StringEndsWithAdvancedFilter_STATUS            `json:"stringEndsWith,omitempty"`
	StringIn                  *StringInAdvancedFilter_STATUS                  `json:"stringIn,omitempty"`
	StringNotIn               *StringNotInAdvancedFilter_STATUS               `json:"stringNotIn,omitempty"`
}

// Storage version of v1api20200601.AzureFunctionEventSubscriptionDestination
type AzureFunctionEventSubscriptionDestination struct {
	EndpointType                  *string                `json:"endpointType,omitempty"`
	MaxEventsPerBatch             *int                   `json:"maxEventsPerBatch,omitempty"`
	PreferredBatchSizeInKilobytes *int                   `json:"preferredBatchSizeInKilobytes,omitempty"`
	PropertyBag                   genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// ResourceReference: The Azure Resource Id that represents the endpoint of the Azure Function destination of an event
	// subscription.
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
}

// Storage version of v1api20200601.AzureFunctionEventSubscriptionDestination_STATUS
type AzureFunctionEventSubscriptionDestination_STATUS struct {
	EndpointType                  *string                `json:"endpointType,omitempty"`
	MaxEventsPerBatch             *int                   `json:"maxEventsPerBatch,omitempty"`
	PreferredBatchSizeInKilobytes *int                   `json:"preferredBatchSizeInKilobytes,omitempty"`
	PropertyBag                   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ResourceId                    *string                `json:"resourceId,omitempty"`
}

// Storage version of v1api20200601.EventHubEventSubscriptionDestination
type EventHubEventSubscriptionDestination struct {
	EndpointType *string                `json:"endpointType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// ResourceReference: The Azure Resource Id that represents the endpoint of an Event Hub destination of an event
	// subscription.
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
}

// Storage version of v1api20200601.EventHubEventSubscriptionDestination_STATUS
type EventHubEventSubscriptionDestination_STATUS struct {
	EndpointType *string                `json:"endpointType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ResourceId   *string                `json:"resourceId,omitempty"`
}

// Storage version of v1api20200601.HybridConnectionEventSubscriptionDestination
type HybridConnectionEventSubscriptionDestination struct {
	EndpointType *string                `json:"endpointType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// ResourceReference: The Azure Resource ID of an hybrid connection that is the destination of an event subscription.
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
}

// Storage version of v1api20200601.HybridConnectionEventSubscriptionDestination_STATUS
type HybridConnectionEventSubscriptionDestination_STATUS struct {
	EndpointType *string                `json:"endpointType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ResourceId   *string                `json:"resourceId,omitempty"`
}

// Storage version of v1api20200601.ServiceBusQueueEventSubscriptionDestination
type ServiceBusQueueEventSubscriptionDestination struct {
	EndpointType *string                `json:"endpointType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// ResourceReference: The Azure Resource Id that represents the endpoint of the Service Bus destination of an event
	// subscription.
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
}

// Storage version of v1api20200601.ServiceBusQueueEventSubscriptionDestination_STATUS
type ServiceBusQueueEventSubscriptionDestination_STATUS struct {
	EndpointType *string                `json:"endpointType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ResourceId   *string                `json:"resourceId,omitempty"`
}

// Storage version of v1api20200601.ServiceBusTopicEventSubscriptionDestination
type ServiceBusTopicEventSubscriptionDestination struct {
	EndpointType *string                `json:"endpointType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// ResourceReference: The Azure Resource Id that represents the endpoint of the Service Bus Topic destination of an event
	// subscription.
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
}

// Storage version of v1api20200601.ServiceBusTopicEventSubscriptionDestination_STATUS
type ServiceBusTopicEventSubscriptionDestination_STATUS struct {
	EndpointType *string                `json:"endpointType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ResourceId   *string                `json:"resourceId,omitempty"`
}

// Storage version of v1api20200601.StorageBlobDeadLetterDestination
type StorageBlobDeadLetterDestination struct {
	BlobContainerName *string                `json:"blobContainerName,omitempty"`
	EndpointType      *string                `json:"endpointType,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// ResourceReference: The Azure Resource ID of the storage account that is the destination of the deadletter events
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
}

// Storage version of v1api20200601.StorageBlobDeadLetterDestination_STATUS
type StorageBlobDeadLetterDestination_STATUS struct {
	BlobContainerName *string                `json:"blobContainerName,omitempty"`
	EndpointType      *string                `json:"endpointType,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ResourceId        *string                `json:"resourceId,omitempty"`
}

// Storage version of v1api20200601.StorageQueueEventSubscriptionDestination
type StorageQueueEventSubscriptionDestination struct {
	EndpointType *string                `json:"endpointType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	QueueName    *string                `json:"queueName,omitempty"`

	// ResourceReference: The Azure Resource ID of the storage account that contains the queue that is the destination of an
	// event subscription.
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
}

// Storage version of v1api20200601.StorageQueueEventSubscriptionDestination_STATUS
type StorageQueueEventSubscriptionDestination_STATUS struct {
	EndpointType *string                `json:"endpointType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	QueueName    *string                `json:"queueName,omitempty"`
	ResourceId   *string                `json:"resourceId,omitempty"`
}

// Storage version of v1api20200601.WebHookEventSubscriptionDestination
type WebHookEventSubscriptionDestination struct {
	AzureActiveDirectoryApplicationIdOrUri *string                     `json:"azureActiveDirectoryApplicationIdOrUri,omitempty"`
	AzureActiveDirectoryTenantId           *string                     `json:"azureActiveDirectoryTenantId,omitempty"`
	EndpointType                           *string                     `json:"endpointType,omitempty"`
	EndpointUrl                            *genruntime.SecretReference `json:"endpointUrl,omitempty"`
	MaxEventsPerBatch                      *int                        `json:"maxEventsPerBatch,omitempty"`
	PreferredBatchSizeInKilobytes          *int                        `json:"preferredBatchSizeInKilobytes,omitempty"`
	PropertyBag                            genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20200601.WebHookEventSubscriptionDestination_STATUS
type WebHookEventSubscriptionDestination_STATUS struct {
	AzureActiveDirectoryApplicationIdOrUri *string                `json:"azureActiveDirectoryApplicationIdOrUri,omitempty"`
	AzureActiveDirectoryTenantId           *string                `json:"azureActiveDirectoryTenantId,omitempty"`
	EndpointBaseUrl                        *string                `json:"endpointBaseUrl,omitempty"`
	EndpointType                           *string                `json:"endpointType,omitempty"`
	MaxEventsPerBatch                      *int                   `json:"maxEventsPerBatch,omitempty"`
	PreferredBatchSizeInKilobytes          *int                   `json:"preferredBatchSizeInKilobytes,omitempty"`
	PropertyBag                            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20200601.BoolEqualsAdvancedFilter
type BoolEqualsAdvancedFilter struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value        *bool                  `json:"value,omitempty"`
}

// Storage version of v1api20200601.BoolEqualsAdvancedFilter_STATUS
type BoolEqualsAdvancedFilter_STATUS struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value        *bool                  `json:"value,omitempty"`
}

// Storage version of v1api20200601.NumberGreaterThanAdvancedFilter
type NumberGreaterThanAdvancedFilter struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value        *float64               `json:"value,omitempty"`
}

// Storage version of v1api20200601.NumberGreaterThanAdvancedFilter_STATUS
type NumberGreaterThanAdvancedFilter_STATUS struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value        *float64               `json:"value,omitempty"`
}

// Storage version of v1api20200601.NumberGreaterThanOrEqualsAdvancedFilter
type NumberGreaterThanOrEqualsAdvancedFilter struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value        *float64               `json:"value,omitempty"`
}

// Storage version of v1api20200601.NumberGreaterThanOrEqualsAdvancedFilter_STATUS
type NumberGreaterThanOrEqualsAdvancedFilter_STATUS struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value        *float64               `json:"value,omitempty"`
}

// Storage version of v1api20200601.NumberInAdvancedFilter
type NumberInAdvancedFilter struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Values       []float64              `json:"values,omitempty"`
}

// Storage version of v1api20200601.NumberInAdvancedFilter_STATUS
type NumberInAdvancedFilter_STATUS struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Values       []float64              `json:"values,omitempty"`
}

// Storage version of v1api20200601.NumberLessThanAdvancedFilter
type NumberLessThanAdvancedFilter struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value        *float64               `json:"value,omitempty"`
}

// Storage version of v1api20200601.NumberLessThanAdvancedFilter_STATUS
type NumberLessThanAdvancedFilter_STATUS struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value        *float64               `json:"value,omitempty"`
}

// Storage version of v1api20200601.NumberLessThanOrEqualsAdvancedFilter
type NumberLessThanOrEqualsAdvancedFilter struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value        *float64               `json:"value,omitempty"`
}

// Storage version of v1api20200601.NumberLessThanOrEqualsAdvancedFilter_STATUS
type NumberLessThanOrEqualsAdvancedFilter_STATUS struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value        *float64               `json:"value,omitempty"`
}

// Storage version of v1api20200601.NumberNotInAdvancedFilter
type NumberNotInAdvancedFilter struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Values       []float64              `json:"values,omitempty"`
}

// Storage version of v1api20200601.NumberNotInAdvancedFilter_STATUS
type NumberNotInAdvancedFilter_STATUS struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Values       []float64              `json:"values,omitempty"`
}

// Storage version of v1api20200601.StringBeginsWithAdvancedFilter
type StringBeginsWithAdvancedFilter struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Values       []string               `json:"values,omitempty"`
}

// Storage version of v1api20200601.StringBeginsWithAdvancedFilter_STATUS
type StringBeginsWithAdvancedFilter_STATUS struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Values       []string               `json:"values,omitempty"`
}

// Storage version of v1api20200601.StringContainsAdvancedFilter
type StringContainsAdvancedFilter struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Values       []string               `json:"values,omitempty"`
}

// Storage version of v1api20200601.StringContainsAdvancedFilter_STATUS
type StringContainsAdvancedFilter_STATUS struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Values       []string               `json:"values,omitempty"`
}

// Storage version of v1api20200601.StringEndsWithAdvancedFilter
type StringEndsWithAdvancedFilter struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Values       []string               `json:"values,omitempty"`
}

// Storage version of v1api20200601.StringEndsWithAdvancedFilter_STATUS
type StringEndsWithAdvancedFilter_STATUS struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Values       []string               `json:"values,omitempty"`
}

// Storage version of v1api20200601.StringInAdvancedFilter
type StringInAdvancedFilter struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Values       []string               `json:"values,omitempty"`
}

// Storage version of v1api20200601.StringInAdvancedFilter_STATUS
type StringInAdvancedFilter_STATUS struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Values       []string               `json:"values,omitempty"`
}

// Storage version of v1api20200601.StringNotInAdvancedFilter
type StringNotInAdvancedFilter struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Values       []string               `json:"values,omitempty"`
}

// Storage version of v1api20200601.StringNotInAdvancedFilter_STATUS
type StringNotInAdvancedFilter_STATUS struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Values       []string               `json:"values,omitempty"`
}

func init() {
	SchemeBuilder.Register(&EventSubscription{}, &EventSubscriptionList{})
}
