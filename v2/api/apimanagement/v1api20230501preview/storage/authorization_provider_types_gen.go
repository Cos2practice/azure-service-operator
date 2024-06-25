// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20230501preview.AuthorizationProvider
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimauthorizationproviders.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationProviders/{authorizationProviderId}
type AuthorizationProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Service_AuthorizationProvider_Spec   `json:"spec,omitempty"`
	Status            Service_AuthorizationProvider_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &AuthorizationProvider{}

// GetConditions returns the conditions of the resource
func (provider *AuthorizationProvider) GetConditions() conditions.Conditions {
	return provider.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (provider *AuthorizationProvider) SetConditions(conditions conditions.Conditions) {
	provider.Status.Conditions = conditions
}

var _ conversion.Convertible = &AuthorizationProvider{}

// ConvertFrom populates our AuthorizationProvider from the provided hub AuthorizationProvider
func (provider *AuthorizationProvider) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.AuthorizationProvider)
	if !ok {
		return fmt.Errorf("expected apimanagement/v1api20220801/storage/AuthorizationProvider but received %T instead", hub)
	}

	return provider.AssignProperties_From_AuthorizationProvider(source)
}

// ConvertTo populates the provided hub AuthorizationProvider from our AuthorizationProvider
func (provider *AuthorizationProvider) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.AuthorizationProvider)
	if !ok {
		return fmt.Errorf("expected apimanagement/v1api20220801/storage/AuthorizationProvider but received %T instead", hub)
	}

	return provider.AssignProperties_To_AuthorizationProvider(destination)
}

var _ genruntime.KubernetesResource = &AuthorizationProvider{}

// AzureName returns the Azure name of the resource
func (provider *AuthorizationProvider) AzureName() string {
	return provider.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01-preview"
func (provider AuthorizationProvider) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (provider *AuthorizationProvider) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (provider *AuthorizationProvider) GetSpec() genruntime.ConvertibleSpec {
	return &provider.Spec
}

// GetStatus returns the status of this resource
func (provider *AuthorizationProvider) GetStatus() genruntime.ConvertibleStatus {
	return &provider.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (provider *AuthorizationProvider) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/authorizationProviders"
func (provider *AuthorizationProvider) GetType() string {
	return "Microsoft.ApiManagement/service/authorizationProviders"
}

// NewEmptyStatus returns a new empty (blank) status
func (provider *AuthorizationProvider) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Service_AuthorizationProvider_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (provider *AuthorizationProvider) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(provider.Spec)
	return provider.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (provider *AuthorizationProvider) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Service_AuthorizationProvider_STATUS); ok {
		provider.Status = *st
		return nil
	}

	// Convert status to required version
	var st Service_AuthorizationProvider_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	provider.Status = st
	return nil
}

// AssignProperties_From_AuthorizationProvider populates our AuthorizationProvider from the provided source AuthorizationProvider
func (provider *AuthorizationProvider) AssignProperties_From_AuthorizationProvider(source *storage.AuthorizationProvider) error {

	// ObjectMeta
	provider.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Service_AuthorizationProvider_Spec
	err := spec.AssignProperties_From_Service_AuthorizationProvider_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Service_AuthorizationProvider_Spec() to populate field Spec")
	}
	provider.Spec = spec

	// Status
	var status Service_AuthorizationProvider_STATUS
	err = status.AssignProperties_From_Service_AuthorizationProvider_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Service_AuthorizationProvider_STATUS() to populate field Status")
	}
	provider.Status = status

	// Invoke the augmentConversionForAuthorizationProvider interface (if implemented) to customize the conversion
	var providerAsAny any = provider
	if augmentedProvider, ok := providerAsAny.(augmentConversionForAuthorizationProvider); ok {
		err := augmentedProvider.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_AuthorizationProvider populates the provided destination AuthorizationProvider from our AuthorizationProvider
func (provider *AuthorizationProvider) AssignProperties_To_AuthorizationProvider(destination *storage.AuthorizationProvider) error {

	// ObjectMeta
	destination.ObjectMeta = *provider.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.Service_AuthorizationProvider_Spec
	err := provider.Spec.AssignProperties_To_Service_AuthorizationProvider_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Service_AuthorizationProvider_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.Service_AuthorizationProvider_STATUS
	err = provider.Status.AssignProperties_To_Service_AuthorizationProvider_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Service_AuthorizationProvider_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForAuthorizationProvider interface (if implemented) to customize the conversion
	var providerAsAny any = provider
	if augmentedProvider, ok := providerAsAny.(augmentConversionForAuthorizationProvider); ok {
		err := augmentedProvider.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (provider *AuthorizationProvider) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: provider.Spec.OriginalVersion,
		Kind:    "AuthorizationProvider",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230501preview.AuthorizationProvider
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimauthorizationproviders.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationProviders/{authorizationProviderId}
type AuthorizationProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthorizationProvider `json:"items"`
}

type augmentConversionForAuthorizationProvider interface {
	AssignPropertiesFrom(src *storage.AuthorizationProvider) error
	AssignPropertiesTo(dst *storage.AuthorizationProvider) error
}

// Storage version of v1api20230501preview.Service_AuthorizationProvider_Spec
type Service_AuthorizationProvider_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName        string                               `json:"azureName,omitempty"`
	DisplayName      *string                              `json:"displayName,omitempty"`
	IdentityProvider *string                              `json:"identityProvider,omitempty"`
	Oauth2           *AuthorizationProviderOAuth2Settings `json:"oauth2,omitempty"`
	OriginalVersion  string                               `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a apimanagement.azure.com/Service resource
	Owner       *genruntime.KnownResourceReference `group:"apimanagement.azure.com" json:"owner,omitempty" kind:"Service"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Service_AuthorizationProvider_Spec{}

// ConvertSpecFrom populates our Service_AuthorizationProvider_Spec from the provided source
func (provider *Service_AuthorizationProvider_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.Service_AuthorizationProvider_Spec)
	if ok {
		// Populate our instance from source
		return provider.AssignProperties_From_Service_AuthorizationProvider_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.Service_AuthorizationProvider_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = provider.AssignProperties_From_Service_AuthorizationProvider_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Service_AuthorizationProvider_Spec
func (provider *Service_AuthorizationProvider_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.Service_AuthorizationProvider_Spec)
	if ok {
		// Populate destination from our instance
		return provider.AssignProperties_To_Service_AuthorizationProvider_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.Service_AuthorizationProvider_Spec{}
	err := provider.AssignProperties_To_Service_AuthorizationProvider_Spec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_Service_AuthorizationProvider_Spec populates our Service_AuthorizationProvider_Spec from the provided source Service_AuthorizationProvider_Spec
func (provider *Service_AuthorizationProvider_Spec) AssignProperties_From_Service_AuthorizationProvider_Spec(source *storage.Service_AuthorizationProvider_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	provider.AzureName = source.AzureName

	// DisplayName
	provider.DisplayName = genruntime.ClonePointerToString(source.DisplayName)

	// IdentityProvider
	provider.IdentityProvider = genruntime.ClonePointerToString(source.IdentityProvider)

	// Oauth2
	if source.Oauth2 != nil {
		var oauth2 AuthorizationProviderOAuth2Settings
		err := oauth2.AssignProperties_From_AuthorizationProviderOAuth2Settings(source.Oauth2)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_AuthorizationProviderOAuth2Settings() to populate field Oauth2")
		}
		provider.Oauth2 = &oauth2
	} else {
		provider.Oauth2 = nil
	}

	// OriginalVersion
	provider.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		provider.Owner = &owner
	} else {
		provider.Owner = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		provider.PropertyBag = propertyBag
	} else {
		provider.PropertyBag = nil
	}

	// Invoke the augmentConversionForService_AuthorizationProvider_Spec interface (if implemented) to customize the conversion
	var providerAsAny any = provider
	if augmentedProvider, ok := providerAsAny.(augmentConversionForService_AuthorizationProvider_Spec); ok {
		err := augmentedProvider.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Service_AuthorizationProvider_Spec populates the provided destination Service_AuthorizationProvider_Spec from our Service_AuthorizationProvider_Spec
func (provider *Service_AuthorizationProvider_Spec) AssignProperties_To_Service_AuthorizationProvider_Spec(destination *storage.Service_AuthorizationProvider_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(provider.PropertyBag)

	// AzureName
	destination.AzureName = provider.AzureName

	// DisplayName
	destination.DisplayName = genruntime.ClonePointerToString(provider.DisplayName)

	// IdentityProvider
	destination.IdentityProvider = genruntime.ClonePointerToString(provider.IdentityProvider)

	// Oauth2
	if provider.Oauth2 != nil {
		var oauth2 storage.AuthorizationProviderOAuth2Settings
		err := provider.Oauth2.AssignProperties_To_AuthorizationProviderOAuth2Settings(&oauth2)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_AuthorizationProviderOAuth2Settings() to populate field Oauth2")
		}
		destination.Oauth2 = &oauth2
	} else {
		destination.Oauth2 = nil
	}

	// OriginalVersion
	destination.OriginalVersion = provider.OriginalVersion

	// Owner
	if provider.Owner != nil {
		owner := provider.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForService_AuthorizationProvider_Spec interface (if implemented) to customize the conversion
	var providerAsAny any = provider
	if augmentedProvider, ok := providerAsAny.(augmentConversionForService_AuthorizationProvider_Spec); ok {
		err := augmentedProvider.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20230501preview.Service_AuthorizationProvider_STATUS
type Service_AuthorizationProvider_STATUS struct {
	Conditions       []conditions.Condition                      `json:"conditions,omitempty"`
	DisplayName      *string                                     `json:"displayName,omitempty"`
	Id               *string                                     `json:"id,omitempty"`
	IdentityProvider *string                                     `json:"identityProvider,omitempty"`
	Name             *string                                     `json:"name,omitempty"`
	Oauth2           *AuthorizationProviderOAuth2Settings_STATUS `json:"oauth2,omitempty"`
	PropertyBag      genruntime.PropertyBag                      `json:"$propertyBag,omitempty"`
	Type             *string                                     `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Service_AuthorizationProvider_STATUS{}

// ConvertStatusFrom populates our Service_AuthorizationProvider_STATUS from the provided source
func (provider *Service_AuthorizationProvider_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.Service_AuthorizationProvider_STATUS)
	if ok {
		// Populate our instance from source
		return provider.AssignProperties_From_Service_AuthorizationProvider_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.Service_AuthorizationProvider_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = provider.AssignProperties_From_Service_AuthorizationProvider_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Service_AuthorizationProvider_STATUS
func (provider *Service_AuthorizationProvider_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.Service_AuthorizationProvider_STATUS)
	if ok {
		// Populate destination from our instance
		return provider.AssignProperties_To_Service_AuthorizationProvider_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.Service_AuthorizationProvider_STATUS{}
	err := provider.AssignProperties_To_Service_AuthorizationProvider_STATUS(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

// AssignProperties_From_Service_AuthorizationProvider_STATUS populates our Service_AuthorizationProvider_STATUS from the provided source Service_AuthorizationProvider_STATUS
func (provider *Service_AuthorizationProvider_STATUS) AssignProperties_From_Service_AuthorizationProvider_STATUS(source *storage.Service_AuthorizationProvider_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	provider.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// DisplayName
	provider.DisplayName = genruntime.ClonePointerToString(source.DisplayName)

	// Id
	provider.Id = genruntime.ClonePointerToString(source.Id)

	// IdentityProvider
	provider.IdentityProvider = genruntime.ClonePointerToString(source.IdentityProvider)

	// Name
	provider.Name = genruntime.ClonePointerToString(source.Name)

	// Oauth2
	if source.Oauth2 != nil {
		var oauth2 AuthorizationProviderOAuth2Settings_STATUS
		err := oauth2.AssignProperties_From_AuthorizationProviderOAuth2Settings_STATUS(source.Oauth2)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_AuthorizationProviderOAuth2Settings_STATUS() to populate field Oauth2")
		}
		provider.Oauth2 = &oauth2
	} else {
		provider.Oauth2 = nil
	}

	// Type
	provider.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		provider.PropertyBag = propertyBag
	} else {
		provider.PropertyBag = nil
	}

	// Invoke the augmentConversionForService_AuthorizationProvider_STATUS interface (if implemented) to customize the conversion
	var providerAsAny any = provider
	if augmentedProvider, ok := providerAsAny.(augmentConversionForService_AuthorizationProvider_STATUS); ok {
		err := augmentedProvider.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Service_AuthorizationProvider_STATUS populates the provided destination Service_AuthorizationProvider_STATUS from our Service_AuthorizationProvider_STATUS
func (provider *Service_AuthorizationProvider_STATUS) AssignProperties_To_Service_AuthorizationProvider_STATUS(destination *storage.Service_AuthorizationProvider_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(provider.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(provider.Conditions)

	// DisplayName
	destination.DisplayName = genruntime.ClonePointerToString(provider.DisplayName)

	// Id
	destination.Id = genruntime.ClonePointerToString(provider.Id)

	// IdentityProvider
	destination.IdentityProvider = genruntime.ClonePointerToString(provider.IdentityProvider)

	// Name
	destination.Name = genruntime.ClonePointerToString(provider.Name)

	// Oauth2
	if provider.Oauth2 != nil {
		var oauth2 storage.AuthorizationProviderOAuth2Settings_STATUS
		err := provider.Oauth2.AssignProperties_To_AuthorizationProviderOAuth2Settings_STATUS(&oauth2)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_AuthorizationProviderOAuth2Settings_STATUS() to populate field Oauth2")
		}
		destination.Oauth2 = &oauth2
	} else {
		destination.Oauth2 = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(provider.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForService_AuthorizationProvider_STATUS interface (if implemented) to customize the conversion
	var providerAsAny any = provider
	if augmentedProvider, ok := providerAsAny.(augmentConversionForService_AuthorizationProvider_STATUS); ok {
		err := augmentedProvider.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForService_AuthorizationProvider_Spec interface {
	AssignPropertiesFrom(src *storage.Service_AuthorizationProvider_Spec) error
	AssignPropertiesTo(dst *storage.Service_AuthorizationProvider_Spec) error
}

type augmentConversionForService_AuthorizationProvider_STATUS interface {
	AssignPropertiesFrom(src *storage.Service_AuthorizationProvider_STATUS) error
	AssignPropertiesTo(dst *storage.Service_AuthorizationProvider_STATUS) error
}

// Storage version of v1api20230501preview.AuthorizationProviderOAuth2Settings
// OAuth2 settings details
type AuthorizationProviderOAuth2Settings struct {
	GrantTypes  *AuthorizationProviderOAuth2GrantTypes `json:"grantTypes,omitempty"`
	PropertyBag genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
	RedirectUrl *string                                `json:"redirectUrl,omitempty"`
}

// AssignProperties_From_AuthorizationProviderOAuth2Settings populates our AuthorizationProviderOAuth2Settings from the provided source AuthorizationProviderOAuth2Settings
func (settings *AuthorizationProviderOAuth2Settings) AssignProperties_From_AuthorizationProviderOAuth2Settings(source *storage.AuthorizationProviderOAuth2Settings) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// GrantTypes
	if source.GrantTypes != nil {
		var grantType AuthorizationProviderOAuth2GrantTypes
		err := grantType.AssignProperties_From_AuthorizationProviderOAuth2GrantTypes(source.GrantTypes)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_AuthorizationProviderOAuth2GrantTypes() to populate field GrantTypes")
		}
		settings.GrantTypes = &grantType
	} else {
		settings.GrantTypes = nil
	}

	// RedirectUrl
	settings.RedirectUrl = genruntime.ClonePointerToString(source.RedirectUrl)

	// Update the property bag
	if len(propertyBag) > 0 {
		settings.PropertyBag = propertyBag
	} else {
		settings.PropertyBag = nil
	}

	// Invoke the augmentConversionForAuthorizationProviderOAuth2Settings interface (if implemented) to customize the conversion
	var settingsAsAny any = settings
	if augmentedSettings, ok := settingsAsAny.(augmentConversionForAuthorizationProviderOAuth2Settings); ok {
		err := augmentedSettings.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_AuthorizationProviderOAuth2Settings populates the provided destination AuthorizationProviderOAuth2Settings from our AuthorizationProviderOAuth2Settings
func (settings *AuthorizationProviderOAuth2Settings) AssignProperties_To_AuthorizationProviderOAuth2Settings(destination *storage.AuthorizationProviderOAuth2Settings) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(settings.PropertyBag)

	// GrantTypes
	if settings.GrantTypes != nil {
		var grantType storage.AuthorizationProviderOAuth2GrantTypes
		err := settings.GrantTypes.AssignProperties_To_AuthorizationProviderOAuth2GrantTypes(&grantType)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_AuthorizationProviderOAuth2GrantTypes() to populate field GrantTypes")
		}
		destination.GrantTypes = &grantType
	} else {
		destination.GrantTypes = nil
	}

	// RedirectUrl
	destination.RedirectUrl = genruntime.ClonePointerToString(settings.RedirectUrl)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForAuthorizationProviderOAuth2Settings interface (if implemented) to customize the conversion
	var settingsAsAny any = settings
	if augmentedSettings, ok := settingsAsAny.(augmentConversionForAuthorizationProviderOAuth2Settings); ok {
		err := augmentedSettings.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20230501preview.AuthorizationProviderOAuth2Settings_STATUS
// OAuth2 settings details
type AuthorizationProviderOAuth2Settings_STATUS struct {
	GrantTypes  *AuthorizationProviderOAuth2GrantTypes_STATUS `json:"grantTypes,omitempty"`
	PropertyBag genruntime.PropertyBag                        `json:"$propertyBag,omitempty"`
	RedirectUrl *string                                       `json:"redirectUrl,omitempty"`
}

// AssignProperties_From_AuthorizationProviderOAuth2Settings_STATUS populates our AuthorizationProviderOAuth2Settings_STATUS from the provided source AuthorizationProviderOAuth2Settings_STATUS
func (settings *AuthorizationProviderOAuth2Settings_STATUS) AssignProperties_From_AuthorizationProviderOAuth2Settings_STATUS(source *storage.AuthorizationProviderOAuth2Settings_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// GrantTypes
	if source.GrantTypes != nil {
		var grantType AuthorizationProviderOAuth2GrantTypes_STATUS
		err := grantType.AssignProperties_From_AuthorizationProviderOAuth2GrantTypes_STATUS(source.GrantTypes)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_AuthorizationProviderOAuth2GrantTypes_STATUS() to populate field GrantTypes")
		}
		settings.GrantTypes = &grantType
	} else {
		settings.GrantTypes = nil
	}

	// RedirectUrl
	settings.RedirectUrl = genruntime.ClonePointerToString(source.RedirectUrl)

	// Update the property bag
	if len(propertyBag) > 0 {
		settings.PropertyBag = propertyBag
	} else {
		settings.PropertyBag = nil
	}

	// Invoke the augmentConversionForAuthorizationProviderOAuth2Settings_STATUS interface (if implemented) to customize the conversion
	var settingsAsAny any = settings
	if augmentedSettings, ok := settingsAsAny.(augmentConversionForAuthorizationProviderOAuth2Settings_STATUS); ok {
		err := augmentedSettings.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_AuthorizationProviderOAuth2Settings_STATUS populates the provided destination AuthorizationProviderOAuth2Settings_STATUS from our AuthorizationProviderOAuth2Settings_STATUS
func (settings *AuthorizationProviderOAuth2Settings_STATUS) AssignProperties_To_AuthorizationProviderOAuth2Settings_STATUS(destination *storage.AuthorizationProviderOAuth2Settings_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(settings.PropertyBag)

	// GrantTypes
	if settings.GrantTypes != nil {
		var grantType storage.AuthorizationProviderOAuth2GrantTypes_STATUS
		err := settings.GrantTypes.AssignProperties_To_AuthorizationProviderOAuth2GrantTypes_STATUS(&grantType)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_AuthorizationProviderOAuth2GrantTypes_STATUS() to populate field GrantTypes")
		}
		destination.GrantTypes = &grantType
	} else {
		destination.GrantTypes = nil
	}

	// RedirectUrl
	destination.RedirectUrl = genruntime.ClonePointerToString(settings.RedirectUrl)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForAuthorizationProviderOAuth2Settings_STATUS interface (if implemented) to customize the conversion
	var settingsAsAny any = settings
	if augmentedSettings, ok := settingsAsAny.(augmentConversionForAuthorizationProviderOAuth2Settings_STATUS); ok {
		err := augmentedSettings.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForAuthorizationProviderOAuth2Settings interface {
	AssignPropertiesFrom(src *storage.AuthorizationProviderOAuth2Settings) error
	AssignPropertiesTo(dst *storage.AuthorizationProviderOAuth2Settings) error
}

type augmentConversionForAuthorizationProviderOAuth2Settings_STATUS interface {
	AssignPropertiesFrom(src *storage.AuthorizationProviderOAuth2Settings_STATUS) error
	AssignPropertiesTo(dst *storage.AuthorizationProviderOAuth2Settings_STATUS) error
}

// Storage version of v1api20230501preview.AuthorizationProviderOAuth2GrantTypes
// Authorization Provider oauth2 grant types settings
type AuthorizationProviderOAuth2GrantTypes struct {
	AuthorizationCode *genruntime.SecretMapReference `json:"authorizationCode,omitempty"`
	ClientCredentials *genruntime.SecretMapReference `json:"clientCredentials,omitempty"`
	PropertyBag       genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_AuthorizationProviderOAuth2GrantTypes populates our AuthorizationProviderOAuth2GrantTypes from the provided source AuthorizationProviderOAuth2GrantTypes
func (types *AuthorizationProviderOAuth2GrantTypes) AssignProperties_From_AuthorizationProviderOAuth2GrantTypes(source *storage.AuthorizationProviderOAuth2GrantTypes) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AuthorizationCode
	if source.AuthorizationCode != nil {
		authorizationCode := source.AuthorizationCode.Copy()
		types.AuthorizationCode = &authorizationCode
	} else {
		types.AuthorizationCode = nil
	}

	// ClientCredentials
	if source.ClientCredentials != nil {
		clientCredential := source.ClientCredentials.Copy()
		types.ClientCredentials = &clientCredential
	} else {
		types.ClientCredentials = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		types.PropertyBag = propertyBag
	} else {
		types.PropertyBag = nil
	}

	// Invoke the augmentConversionForAuthorizationProviderOAuth2GrantTypes interface (if implemented) to customize the conversion
	var typesAsAny any = types
	if augmentedTypes, ok := typesAsAny.(augmentConversionForAuthorizationProviderOAuth2GrantTypes); ok {
		err := augmentedTypes.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_AuthorizationProviderOAuth2GrantTypes populates the provided destination AuthorizationProviderOAuth2GrantTypes from our AuthorizationProviderOAuth2GrantTypes
func (types *AuthorizationProviderOAuth2GrantTypes) AssignProperties_To_AuthorizationProviderOAuth2GrantTypes(destination *storage.AuthorizationProviderOAuth2GrantTypes) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(types.PropertyBag)

	// AuthorizationCode
	if types.AuthorizationCode != nil {
		authorizationCode := types.AuthorizationCode.Copy()
		destination.AuthorizationCode = &authorizationCode
	} else {
		destination.AuthorizationCode = nil
	}

	// ClientCredentials
	if types.ClientCredentials != nil {
		clientCredential := types.ClientCredentials.Copy()
		destination.ClientCredentials = &clientCredential
	} else {
		destination.ClientCredentials = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForAuthorizationProviderOAuth2GrantTypes interface (if implemented) to customize the conversion
	var typesAsAny any = types
	if augmentedTypes, ok := typesAsAny.(augmentConversionForAuthorizationProviderOAuth2GrantTypes); ok {
		err := augmentedTypes.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20230501preview.AuthorizationProviderOAuth2GrantTypes_STATUS
// Authorization Provider oauth2 grant types settings
type AuthorizationProviderOAuth2GrantTypes_STATUS struct {
	AuthorizationCode map[string]string      `json:"authorizationCode,omitempty"`
	ClientCredentials map[string]string      `json:"clientCredentials,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_AuthorizationProviderOAuth2GrantTypes_STATUS populates our AuthorizationProviderOAuth2GrantTypes_STATUS from the provided source AuthorizationProviderOAuth2GrantTypes_STATUS
func (types *AuthorizationProviderOAuth2GrantTypes_STATUS) AssignProperties_From_AuthorizationProviderOAuth2GrantTypes_STATUS(source *storage.AuthorizationProviderOAuth2GrantTypes_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AuthorizationCode
	types.AuthorizationCode = genruntime.CloneMapOfStringToString(source.AuthorizationCode)

	// ClientCredentials
	types.ClientCredentials = genruntime.CloneMapOfStringToString(source.ClientCredentials)

	// Update the property bag
	if len(propertyBag) > 0 {
		types.PropertyBag = propertyBag
	} else {
		types.PropertyBag = nil
	}

	// Invoke the augmentConversionForAuthorizationProviderOAuth2GrantTypes_STATUS interface (if implemented) to customize the conversion
	var typesAsAny any = types
	if augmentedTypes, ok := typesAsAny.(augmentConversionForAuthorizationProviderOAuth2GrantTypes_STATUS); ok {
		err := augmentedTypes.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_AuthorizationProviderOAuth2GrantTypes_STATUS populates the provided destination AuthorizationProviderOAuth2GrantTypes_STATUS from our AuthorizationProviderOAuth2GrantTypes_STATUS
func (types *AuthorizationProviderOAuth2GrantTypes_STATUS) AssignProperties_To_AuthorizationProviderOAuth2GrantTypes_STATUS(destination *storage.AuthorizationProviderOAuth2GrantTypes_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(types.PropertyBag)

	// AuthorizationCode
	destination.AuthorizationCode = genruntime.CloneMapOfStringToString(types.AuthorizationCode)

	// ClientCredentials
	destination.ClientCredentials = genruntime.CloneMapOfStringToString(types.ClientCredentials)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForAuthorizationProviderOAuth2GrantTypes_STATUS interface (if implemented) to customize the conversion
	var typesAsAny any = types
	if augmentedTypes, ok := typesAsAny.(augmentConversionForAuthorizationProviderOAuth2GrantTypes_STATUS); ok {
		err := augmentedTypes.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForAuthorizationProviderOAuth2GrantTypes interface {
	AssignPropertiesFrom(src *storage.AuthorizationProviderOAuth2GrantTypes) error
	AssignPropertiesTo(dst *storage.AuthorizationProviderOAuth2GrantTypes) error
}

type augmentConversionForAuthorizationProviderOAuth2GrantTypes_STATUS interface {
	AssignPropertiesFrom(src *storage.AuthorizationProviderOAuth2GrantTypes_STATUS) error
	AssignPropertiesTo(dst *storage.AuthorizationProviderOAuth2GrantTypes_STATUS) error
}

func init() {
	SchemeBuilder.Register(&AuthorizationProvider{}, &AuthorizationProviderList{})
}