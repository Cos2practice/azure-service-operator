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

// +kubebuilder:rbac:groups=insights.azure.com,resources=actiongroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=insights.azure.com,resources={actiongroups/status,actiongroups/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20230101.ActionGroup
// Generator information:
// - Generated from: /monitor/resource-manager/Microsoft.Insights/stable/2023-01-01/actionGroups_API.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/actionGroups/{actionGroupName}
type ActionGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ActionGroup_Spec           `json:"spec,omitempty"`
	Status            ActionGroupResource_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ActionGroup{}

// GetConditions returns the conditions of the resource
func (group *ActionGroup) GetConditions() conditions.Conditions {
	return group.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (group *ActionGroup) SetConditions(conditions conditions.Conditions) {
	group.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &ActionGroup{}

// AzureName returns the Azure name of the resource
func (group *ActionGroup) AzureName() string {
	return group.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-01-01"
func (group ActionGroup) GetAPIVersion() string {
	return "2023-01-01"
}

// GetResourceScope returns the scope of the resource
func (group *ActionGroup) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (group *ActionGroup) GetSpec() genruntime.ConvertibleSpec {
	return &group.Spec
}

// GetStatus returns the status of this resource
func (group *ActionGroup) GetStatus() genruntime.ConvertibleStatus {
	return &group.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (group *ActionGroup) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Insights/actionGroups"
func (group *ActionGroup) GetType() string {
	return "Microsoft.Insights/actionGroups"
}

// NewEmptyStatus returns a new empty (blank) status
func (group *ActionGroup) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ActionGroupResource_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (group *ActionGroup) Owner() *genruntime.ResourceReference {
	ownerGroup, ownerKind := genruntime.LookupOwnerGroupKind(group.Spec)
	return group.Spec.Owner.AsResourceReference(ownerGroup, ownerKind)
}

// SetStatus sets the status of this resource
func (group *ActionGroup) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ActionGroupResource_STATUS); ok {
		group.Status = *st
		return nil
	}

	// Convert status to required version
	var st ActionGroupResource_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	group.Status = st
	return nil
}

// Hub marks that this ActionGroup is the hub type for conversion
func (group *ActionGroup) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (group *ActionGroup) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: group.Spec.OriginalVersion,
		Kind:    "ActionGroup",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230101.ActionGroup
// Generator information:
// - Generated from: /monitor/resource-manager/Microsoft.Insights/stable/2023-01-01/actionGroups_API.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/actionGroups/{actionGroupName}
type ActionGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ActionGroup `json:"items"`
}

// Storage version of v1api20230101.ActionGroup_Spec
type ActionGroup_Spec struct {
	ArmRoleReceivers           []ArmRoleReceiver           `json:"armRoleReceivers,omitempty"`
	AutomationRunbookReceivers []AutomationRunbookReceiver `json:"automationRunbookReceivers,omitempty"`
	AzureAppPushReceivers      []AzureAppPushReceiver      `json:"azureAppPushReceivers,omitempty"`
	AzureFunctionReceivers     []AzureFunctionReceiver     `json:"azureFunctionReceivers,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName         string             `json:"azureName,omitempty"`
	EmailReceivers    []EmailReceiver    `json:"emailReceivers,omitempty"`
	Enabled           *bool              `json:"enabled,omitempty"`
	EventHubReceivers []EventHubReceiver `json:"eventHubReceivers,omitempty"`
	GroupShortName    *string            `json:"groupShortName,omitempty"`
	ItsmReceivers     []ItsmReceiver     `json:"itsmReceivers,omitempty"`
	Location          *string            `json:"location,omitempty"`
	LogicAppReceivers []LogicAppReceiver `json:"logicAppReceivers,omitempty"`
	OriginalVersion   string             `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner            *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag      genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	SmsReceivers     []SmsReceiver                      `json:"smsReceivers,omitempty"`
	Tags             map[string]string                  `json:"tags,omitempty"`
	VoiceReceivers   []VoiceReceiver                    `json:"voiceReceivers,omitempty"`
	WebhookReceivers []WebhookReceiver                  `json:"webhookReceivers,omitempty"`
}

var _ genruntime.ConvertibleSpec = &ActionGroup_Spec{}

// ConvertSpecFrom populates our ActionGroup_Spec from the provided source
func (group *ActionGroup_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == group {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(group)
}

// ConvertSpecTo populates the provided destination from our ActionGroup_Spec
func (group *ActionGroup_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == group {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(group)
}

// Storage version of v1api20230101.ActionGroupResource_STATUS
// An action group resource.
type ActionGroupResource_STATUS struct {
	ArmRoleReceivers           []ArmRoleReceiver_STATUS           `json:"armRoleReceivers,omitempty"`
	AutomationRunbookReceivers []AutomationRunbookReceiver_STATUS `json:"automationRunbookReceivers,omitempty"`
	AzureAppPushReceivers      []AzureAppPushReceiver_STATUS      `json:"azureAppPushReceivers,omitempty"`
	AzureFunctionReceivers     []AzureFunctionReceiver_STATUS     `json:"azureFunctionReceivers,omitempty"`
	Conditions                 []conditions.Condition             `json:"conditions,omitempty"`
	EmailReceivers             []EmailReceiver_STATUS             `json:"emailReceivers,omitempty"`
	Enabled                    *bool                              `json:"enabled,omitempty"`
	EventHubReceivers          []EventHubReceiver_STATUS          `json:"eventHubReceivers,omitempty"`
	GroupShortName             *string                            `json:"groupShortName,omitempty"`
	Id                         *string                            `json:"id,omitempty"`
	ItsmReceivers              []ItsmReceiver_STATUS              `json:"itsmReceivers,omitempty"`
	Location                   *string                            `json:"location,omitempty"`
	LogicAppReceivers          []LogicAppReceiver_STATUS          `json:"logicAppReceivers,omitempty"`
	Name                       *string                            `json:"name,omitempty"`
	PropertyBag                genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	SmsReceivers               []SmsReceiver_STATUS               `json:"smsReceivers,omitempty"`
	Tags                       map[string]string                  `json:"tags,omitempty"`
	Type                       *string                            `json:"type,omitempty"`
	VoiceReceivers             []VoiceReceiver_STATUS             `json:"voiceReceivers,omitempty"`
	WebhookReceivers           []WebhookReceiver_STATUS           `json:"webhookReceivers,omitempty"`
}

var _ genruntime.ConvertibleStatus = &ActionGroupResource_STATUS{}

// ConvertStatusFrom populates our ActionGroupResource_STATUS from the provided source
func (resource *ActionGroupResource_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == resource {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(resource)
}

// ConvertStatusTo populates the provided destination from our ActionGroupResource_STATUS
func (resource *ActionGroupResource_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == resource {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(resource)
}

// Storage version of v1api20230101.APIVersion
// +kubebuilder:validation:Enum={"2023-01-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2023-01-01")

// Storage version of v1api20230101.ArmRoleReceiver
// An arm role receiver.
type ArmRoleReceiver struct {
	Name                 *string                `json:"name,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RoleId               *string                `json:"roleId,omitempty"`
	UseCommonAlertSchema *bool                  `json:"useCommonAlertSchema,omitempty"`
}

// Storage version of v1api20230101.ArmRoleReceiver_STATUS
// An arm role receiver.
type ArmRoleReceiver_STATUS struct {
	Name                 *string                `json:"name,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RoleId               *string                `json:"roleId,omitempty"`
	UseCommonAlertSchema *bool                  `json:"useCommonAlertSchema,omitempty"`
}

// Storage version of v1api20230101.AutomationRunbookReceiver
// The Azure Automation Runbook notification receiver.
type AutomationRunbookReceiver struct {
	AutomationAccountId  *string                `json:"automationAccountId,omitempty"`
	IsGlobalRunbook      *bool                  `json:"isGlobalRunbook,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RunbookName          *string                `json:"runbookName,omitempty"`
	ServiceUri           *string                `json:"serviceUri,omitempty"`
	UseCommonAlertSchema *bool                  `json:"useCommonAlertSchema,omitempty"`

	// +kubebuilder:validation:Required
	// WebhookResourceReference: The resource id for webhook linked to this runbook.
	WebhookResourceReference *genruntime.ResourceReference `armReference:"WebhookResourceId" json:"webhookResourceReference,omitempty"`
}

// Storage version of v1api20230101.AutomationRunbookReceiver_STATUS
// The Azure Automation Runbook notification receiver.
type AutomationRunbookReceiver_STATUS struct {
	AutomationAccountId  *string                `json:"automationAccountId,omitempty"`
	IsGlobalRunbook      *bool                  `json:"isGlobalRunbook,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RunbookName          *string                `json:"runbookName,omitempty"`
	ServiceUri           *string                `json:"serviceUri,omitempty"`
	UseCommonAlertSchema *bool                  `json:"useCommonAlertSchema,omitempty"`
	WebhookResourceId    *string                `json:"webhookResourceId,omitempty"`
}

// Storage version of v1api20230101.AzureAppPushReceiver
// The Azure mobile App push notification receiver.
type AzureAppPushReceiver struct {
	EmailAddress *string                `json:"emailAddress,omitempty"`
	Name         *string                `json:"name,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230101.AzureAppPushReceiver_STATUS
// The Azure mobile App push notification receiver.
type AzureAppPushReceiver_STATUS struct {
	EmailAddress *string                `json:"emailAddress,omitempty"`
	Name         *string                `json:"name,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230101.AzureFunctionReceiver
// An azure function receiver.
type AzureFunctionReceiver struct {
	// +kubebuilder:validation:Required
	// FunctionAppResourceReference: The azure resource id of the function app.
	FunctionAppResourceReference *genruntime.ResourceReference `armReference:"FunctionAppResourceId" json:"functionAppResourceReference,omitempty"`
	FunctionName                 *string                       `json:"functionName,omitempty"`
	HttpTriggerUrl               *string                       `json:"httpTriggerUrl,omitempty"`
	Name                         *string                       `json:"name,omitempty"`
	PropertyBag                  genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	UseCommonAlertSchema         *bool                         `json:"useCommonAlertSchema,omitempty"`
}

// Storage version of v1api20230101.AzureFunctionReceiver_STATUS
// An azure function receiver.
type AzureFunctionReceiver_STATUS struct {
	FunctionAppResourceId *string                `json:"functionAppResourceId,omitempty"`
	FunctionName          *string                `json:"functionName,omitempty"`
	HttpTriggerUrl        *string                `json:"httpTriggerUrl,omitempty"`
	Name                  *string                `json:"name,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	UseCommonAlertSchema  *bool                  `json:"useCommonAlertSchema,omitempty"`
}

// Storage version of v1api20230101.EmailReceiver
// An email receiver.
type EmailReceiver struct {
	EmailAddress         *string                `json:"emailAddress,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	UseCommonAlertSchema *bool                  `json:"useCommonAlertSchema,omitempty"`
}

// Storage version of v1api20230101.EmailReceiver_STATUS
// An email receiver.
type EmailReceiver_STATUS struct {
	EmailAddress         *string                `json:"emailAddress,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status               *string                `json:"status,omitempty"`
	UseCommonAlertSchema *bool                  `json:"useCommonAlertSchema,omitempty"`
}

// Storage version of v1api20230101.EventHubReceiver
// An Event hub receiver.
type EventHubReceiver struct {
	EventHubName         *string                `json:"eventHubName,omitempty"`
	EventHubNameSpace    *string                `json:"eventHubNameSpace,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SubscriptionId       *string                `json:"subscriptionId,omitempty"`
	TenantId             *string                `json:"tenantId,omitempty"`
	UseCommonAlertSchema *bool                  `json:"useCommonAlertSchema,omitempty"`
}

// Storage version of v1api20230101.EventHubReceiver_STATUS
// An Event hub receiver.
type EventHubReceiver_STATUS struct {
	EventHubName         *string                `json:"eventHubName,omitempty"`
	EventHubNameSpace    *string                `json:"eventHubNameSpace,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SubscriptionId       *string                `json:"subscriptionId,omitempty"`
	TenantId             *string                `json:"tenantId,omitempty"`
	UseCommonAlertSchema *bool                  `json:"useCommonAlertSchema,omitempty"`
}

// Storage version of v1api20230101.ItsmReceiver
// An Itsm receiver.
type ItsmReceiver struct {
	ConnectionId        *string                `json:"connectionId,omitempty"`
	Name                *string                `json:"name,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Region              *string                `json:"region,omitempty"`
	TicketConfiguration *string                `json:"ticketConfiguration,omitempty"`
	WorkspaceId         *string                `json:"workspaceId,omitempty"`
}

// Storage version of v1api20230101.ItsmReceiver_STATUS
// An Itsm receiver.
type ItsmReceiver_STATUS struct {
	ConnectionId        *string                `json:"connectionId,omitempty"`
	Name                *string                `json:"name,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Region              *string                `json:"region,omitempty"`
	TicketConfiguration *string                `json:"ticketConfiguration,omitempty"`
	WorkspaceId         *string                `json:"workspaceId,omitempty"`
}

// Storage version of v1api20230101.LogicAppReceiver
// A logic app receiver.
type LogicAppReceiver struct {
	CallbackUrl *string                `json:"callbackUrl,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// +kubebuilder:validation:Required
	// ResourceReference: The azure resource id of the logic app receiver.
	ResourceReference    *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
	UseCommonAlertSchema *bool                         `json:"useCommonAlertSchema,omitempty"`
}

// Storage version of v1api20230101.LogicAppReceiver_STATUS
// A logic app receiver.
type LogicAppReceiver_STATUS struct {
	CallbackUrl          *string                `json:"callbackUrl,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ResourceId           *string                `json:"resourceId,omitempty"`
	UseCommonAlertSchema *bool                  `json:"useCommonAlertSchema,omitempty"`
}

// Storage version of v1api20230101.SmsReceiver
// An SMS receiver.
type SmsReceiver struct {
	CountryCode *string                `json:"countryCode,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PhoneNumber *string                `json:"phoneNumber,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230101.SmsReceiver_STATUS
// An SMS receiver.
type SmsReceiver_STATUS struct {
	CountryCode *string                `json:"countryCode,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PhoneNumber *string                `json:"phoneNumber,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status      *string                `json:"status,omitempty"`
}

// Storage version of v1api20230101.VoiceReceiver
// A voice receiver.
type VoiceReceiver struct {
	CountryCode *string                `json:"countryCode,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PhoneNumber *string                `json:"phoneNumber,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230101.VoiceReceiver_STATUS
// A voice receiver.
type VoiceReceiver_STATUS struct {
	CountryCode *string                `json:"countryCode,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PhoneNumber *string                `json:"phoneNumber,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230101.WebhookReceiver
// A webhook receiver.
type WebhookReceiver struct {
	IdentifierUri        *string                `json:"identifierUri,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	ObjectId             *string                `json:"objectId,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ServiceUri           *string                `json:"serviceUri,omitempty"`
	TenantId             *string                `json:"tenantId,omitempty"`
	UseAadAuth           *bool                  `json:"useAadAuth,omitempty"`
	UseCommonAlertSchema *bool                  `json:"useCommonAlertSchema,omitempty"`
}

// Storage version of v1api20230101.WebhookReceiver_STATUS
// A webhook receiver.
type WebhookReceiver_STATUS struct {
	IdentifierUri        *string                `json:"identifierUri,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	ObjectId             *string                `json:"objectId,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ServiceUri           *string                `json:"serviceUri,omitempty"`
	TenantId             *string                `json:"tenantId,omitempty"`
	UseAadAuth           *bool                  `json:"useAadAuth,omitempty"`
	UseCommonAlertSchema *bool                  `json:"useCommonAlertSchema,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ActionGroup{}, &ActionGroupList{})
}
