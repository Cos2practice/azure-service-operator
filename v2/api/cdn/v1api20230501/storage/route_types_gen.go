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

// +kubebuilder:rbac:groups=cdn.azure.com,resources=routes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cdn.azure.com,resources={routes/status,routes/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20230501.Route
// Generator information:
// - Generated from: /cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/afdx.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/afdEndpoints/{endpointName}/routes/{routeName}
type Route struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Profiles_AfdEndpoints_Route_Spec   `json:"spec,omitempty"`
	Status            Profiles_AfdEndpoints_Route_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Route{}

// GetConditions returns the conditions of the resource
func (route *Route) GetConditions() conditions.Conditions {
	return route.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (route *Route) SetConditions(conditions conditions.Conditions) {
	route.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &Route{}

// AzureName returns the Azure name of the resource
func (route *Route) AzureName() string {
	return route.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01"
func (route Route) GetAPIVersion() string {
	return "2023-05-01"
}

// GetResourceScope returns the scope of the resource
func (route *Route) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (route *Route) GetSpec() genruntime.ConvertibleSpec {
	return &route.Spec
}

// GetStatus returns the status of this resource
func (route *Route) GetStatus() genruntime.ConvertibleStatus {
	return &route.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (route *Route) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cdn/profiles/afdEndpoints/routes"
func (route *Route) GetType() string {
	return "Microsoft.Cdn/profiles/afdEndpoints/routes"
}

// NewEmptyStatus returns a new empty (blank) status
func (route *Route) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Profiles_AfdEndpoints_Route_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (route *Route) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(route.Spec)
	return route.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (route *Route) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Profiles_AfdEndpoints_Route_STATUS); ok {
		route.Status = *st
		return nil
	}

	// Convert status to required version
	var st Profiles_AfdEndpoints_Route_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	route.Status = st
	return nil
}

// Hub marks that this Route is the hub type for conversion
func (route *Route) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (route *Route) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: route.Spec.OriginalVersion,
		Kind:    "Route",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230501.Route
// Generator information:
// - Generated from: /cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/afdx.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/afdEndpoints/{endpointName}/routes/{routeName}
type RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route `json:"items"`
}

// Storage version of v1api20230501.Profiles_AfdEndpoints_Route_Spec
type Profiles_AfdEndpoints_Route_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName           string                       `json:"azureName,omitempty"`
	CacheConfiguration  *AfdRouteCacheConfiguration  `json:"cacheConfiguration,omitempty"`
	CustomDomains       []ActivatedResourceReference `json:"customDomains,omitempty"`
	EnabledState        *string                      `json:"enabledState,omitempty"`
	ForwardingProtocol  *string                      `json:"forwardingProtocol,omitempty"`
	HttpsRedirect       *string                      `json:"httpsRedirect,omitempty"`
	LinkToDefaultDomain *string                      `json:"linkToDefaultDomain,omitempty"`
	OriginGroup         *ResourceReference           `json:"originGroup,omitempty"`
	OriginPath          *string                      `json:"originPath,omitempty"`
	OriginalVersion     string                       `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a cdn.azure.com/AfdEndpoint resource
	Owner              *genruntime.KnownResourceReference `group:"cdn.azure.com" json:"owner,omitempty" kind:"AfdEndpoint"`
	PatternsToMatch    []string                           `json:"patternsToMatch,omitempty"`
	PropertyBag        genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	RuleSets           []ResourceReference                `json:"ruleSets,omitempty"`
	SupportedProtocols []string                           `json:"supportedProtocols,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Profiles_AfdEndpoints_Route_Spec{}

// ConvertSpecFrom populates our Profiles_AfdEndpoints_Route_Spec from the provided source
func (route *Profiles_AfdEndpoints_Route_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == route {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(route)
}

// ConvertSpecTo populates the provided destination from our Profiles_AfdEndpoints_Route_Spec
func (route *Profiles_AfdEndpoints_Route_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == route {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(route)
}

// Storage version of v1api20230501.Profiles_AfdEndpoints_Route_STATUS
type Profiles_AfdEndpoints_Route_STATUS struct {
	CacheConfiguration  *AfdRouteCacheConfiguration_STATUS                                                  `json:"cacheConfiguration,omitempty"`
	Conditions          []conditions.Condition                                                              `json:"conditions,omitempty"`
	CustomDomains       []ActivatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbedded `json:"customDomains,omitempty"`
	DeploymentStatus    *string                                                                             `json:"deploymentStatus,omitempty"`
	EnabledState        *string                                                                             `json:"enabledState,omitempty"`
	EndpointName        *string                                                                             `json:"endpointName,omitempty"`
	ForwardingProtocol  *string                                                                             `json:"forwardingProtocol,omitempty"`
	HttpsRedirect       *string                                                                             `json:"httpsRedirect,omitempty"`
	Id                  *string                                                                             `json:"id,omitempty"`
	LinkToDefaultDomain *string                                                                             `json:"linkToDefaultDomain,omitempty"`
	Name                *string                                                                             `json:"name,omitempty"`
	OriginGroup         *ResourceReference_STATUS                                                           `json:"originGroup,omitempty"`
	OriginPath          *string                                                                             `json:"originPath,omitempty"`
	PatternsToMatch     []string                                                                            `json:"patternsToMatch,omitempty"`
	PropertyBag         genruntime.PropertyBag                                                              `json:"$propertyBag,omitempty"`
	ProvisioningState   *string                                                                             `json:"provisioningState,omitempty"`
	RuleSets            []ResourceReference_STATUS                                                          `json:"ruleSets,omitempty"`
	SupportedProtocols  []string                                                                            `json:"supportedProtocols,omitempty"`
	SystemData          *SystemData_STATUS                                                                  `json:"systemData,omitempty"`
	Type                *string                                                                             `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Profiles_AfdEndpoints_Route_STATUS{}

// ConvertStatusFrom populates our Profiles_AfdEndpoints_Route_STATUS from the provided source
func (route *Profiles_AfdEndpoints_Route_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == route {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(route)
}

// ConvertStatusTo populates the provided destination from our Profiles_AfdEndpoints_Route_STATUS
func (route *Profiles_AfdEndpoints_Route_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == route {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(route)
}

// Storage version of v1api20230501.ActivatedResourceReference
// Reference to another resource along with its state.
type ActivatedResourceReference struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// Reference: Resource ID.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1api20230501.ActivatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbedded
// Reference to another resource along with its state.
type ActivatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230501.AfdRouteCacheConfiguration
// Caching settings for a caching-type route. To disable caching, do not provide a cacheConfiguration object.
type AfdRouteCacheConfiguration struct {
	CompressionSettings        *CompressionSettings   `json:"compressionSettings,omitempty"`
	PropertyBag                genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	QueryParameters            *string                `json:"queryParameters,omitempty"`
	QueryStringCachingBehavior *string                `json:"queryStringCachingBehavior,omitempty"`
}

// Storage version of v1api20230501.AfdRouteCacheConfiguration_STATUS
// Caching settings for a caching-type route. To disable caching, do not provide a cacheConfiguration object.
type AfdRouteCacheConfiguration_STATUS struct {
	CompressionSettings        *CompressionSettings_STATUS `json:"compressionSettings,omitempty"`
	PropertyBag                genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	QueryParameters            *string                     `json:"queryParameters,omitempty"`
	QueryStringCachingBehavior *string                     `json:"queryStringCachingBehavior,omitempty"`
}

// Storage version of v1api20230501.CompressionSettings
// settings for compression.
type CompressionSettings struct {
	ContentTypesToCompress []string               `json:"contentTypesToCompress,omitempty"`
	IsCompressionEnabled   *bool                  `json:"isCompressionEnabled,omitempty"`
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230501.CompressionSettings_STATUS
// settings for compression.
type CompressionSettings_STATUS struct {
	ContentTypesToCompress []string               `json:"contentTypesToCompress,omitempty"`
	IsCompressionEnabled   *bool                  `json:"isCompressionEnabled,omitempty"`
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Route{}, &RouteList{})
}
