// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Profiles_AfdEndpoints_Route_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: The JSON object that contains the properties of the Routes to create.
	Properties *RouteProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Profiles_AfdEndpoints_Route_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01"
func (route Profiles_AfdEndpoints_Route_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (route *Profiles_AfdEndpoints_Route_Spec_ARM) GetName() string {
	return route.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cdn/profiles/afdEndpoints/routes"
func (route *Profiles_AfdEndpoints_Route_Spec_ARM) GetType() string {
	return "Microsoft.Cdn/profiles/afdEndpoints/routes"
}

// The JSON object that contains the properties of the Routes to create.
type RouteProperties_ARM struct {
	// CacheConfiguration: The caching configuration for this route. To disable caching, do not provide a cacheConfiguration
	// object.
	CacheConfiguration *AfdRouteCacheConfiguration_ARM `json:"cacheConfiguration,omitempty"`

	// CustomDomains: Domains referenced by this endpoint.
	CustomDomains []ActivatedResourceReference_ARM `json:"customDomains,omitempty"`

	// EnabledState: Whether to enable use of this rule. Permitted values are 'Enabled' or 'Disabled'
	EnabledState *RouteProperties_EnabledState `json:"enabledState,omitempty"`

	// ForwardingProtocol: Protocol this rule will use when forwarding traffic to backends.
	ForwardingProtocol *RouteProperties_ForwardingProtocol `json:"forwardingProtocol,omitempty"`

	// HttpsRedirect: Whether to automatically redirect HTTP traffic to HTTPS traffic. Note that this is a easy way to set up
	// this rule and it will be the first rule that gets executed.
	HttpsRedirect *RouteProperties_HttpsRedirect `json:"httpsRedirect,omitempty"`

	// LinkToDefaultDomain: whether this route will be linked to the default endpoint domain.
	LinkToDefaultDomain *RouteProperties_LinkToDefaultDomain `json:"linkToDefaultDomain,omitempty"`

	// OriginGroup: A reference to the origin group.
	OriginGroup *ResourceReference_ARM `json:"originGroup,omitempty"`

	// OriginPath: A directory path on the origin that AzureFrontDoor can use to retrieve content from, e.g.
	// contoso.cloudapp.net/originpath.
	OriginPath *string `json:"originPath,omitempty"`

	// PatternsToMatch: The route patterns of the rule.
	PatternsToMatch []string `json:"patternsToMatch,omitempty"`

	// RuleSets: rule sets referenced by this endpoint.
	RuleSets []ResourceReference_ARM `json:"ruleSets,omitempty"`

	// SupportedProtocols: List of supported protocols for this route.
	SupportedProtocols []AFDEndpointProtocols `json:"supportedProtocols,omitempty"`
}

// Reference to another resource along with its state.
type ActivatedResourceReference_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Supported protocols for the customer's endpoint.
// +kubebuilder:validation:Enum={"Http","Https"}
type AFDEndpointProtocols string

const (
	AFDEndpointProtocols_Http  = AFDEndpointProtocols("Http")
	AFDEndpointProtocols_Https = AFDEndpointProtocols("Https")
)

// Caching settings for a caching-type route. To disable caching, do not provide a cacheConfiguration object.
type AfdRouteCacheConfiguration_ARM struct {
	// CompressionSettings: compression settings.
	CompressionSettings *CompressionSettings_ARM `json:"compressionSettings,omitempty"`

	// QueryParameters: query parameters to include or exclude (comma separated).
	QueryParameters *string `json:"queryParameters,omitempty"`

	// QueryStringCachingBehavior: Defines how Frontdoor caches requests that include query strings. You can ignore any query
	// strings when caching, ignore specific query strings, cache every request with a unique URL, or cache specific query
	// strings.
	QueryStringCachingBehavior *AfdRouteCacheConfiguration_QueryStringCachingBehavior `json:"queryStringCachingBehavior,omitempty"`
}

// Reference to another resource.
type ResourceReference_ARM struct {
	Id *string `json:"id,omitempty"`
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type RouteProperties_EnabledState string

const (
	RouteProperties_EnabledState_Disabled = RouteProperties_EnabledState("Disabled")
	RouteProperties_EnabledState_Enabled  = RouteProperties_EnabledState("Enabled")
)

// +kubebuilder:validation:Enum={"HttpOnly","HttpsOnly","MatchRequest"}
type RouteProperties_ForwardingProtocol string

const (
	RouteProperties_ForwardingProtocol_HttpOnly     = RouteProperties_ForwardingProtocol("HttpOnly")
	RouteProperties_ForwardingProtocol_HttpsOnly    = RouteProperties_ForwardingProtocol("HttpsOnly")
	RouteProperties_ForwardingProtocol_MatchRequest = RouteProperties_ForwardingProtocol("MatchRequest")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type RouteProperties_HttpsRedirect string

const (
	RouteProperties_HttpsRedirect_Disabled = RouteProperties_HttpsRedirect("Disabled")
	RouteProperties_HttpsRedirect_Enabled  = RouteProperties_HttpsRedirect("Enabled")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type RouteProperties_LinkToDefaultDomain string

const (
	RouteProperties_LinkToDefaultDomain_Disabled = RouteProperties_LinkToDefaultDomain("Disabled")
	RouteProperties_LinkToDefaultDomain_Enabled  = RouteProperties_LinkToDefaultDomain("Enabled")
)

// +kubebuilder:validation:Enum={"IgnoreQueryString","IgnoreSpecifiedQueryStrings","IncludeSpecifiedQueryStrings","UseQueryString"}
type AfdRouteCacheConfiguration_QueryStringCachingBehavior string

const (
	AfdRouteCacheConfiguration_QueryStringCachingBehavior_IgnoreQueryString            = AfdRouteCacheConfiguration_QueryStringCachingBehavior("IgnoreQueryString")
	AfdRouteCacheConfiguration_QueryStringCachingBehavior_IgnoreSpecifiedQueryStrings  = AfdRouteCacheConfiguration_QueryStringCachingBehavior("IgnoreSpecifiedQueryStrings")
	AfdRouteCacheConfiguration_QueryStringCachingBehavior_IncludeSpecifiedQueryStrings = AfdRouteCacheConfiguration_QueryStringCachingBehavior("IncludeSpecifiedQueryStrings")
	AfdRouteCacheConfiguration_QueryStringCachingBehavior_UseQueryString               = AfdRouteCacheConfiguration_QueryStringCachingBehavior("UseQueryString")
)

// settings for compression.
type CompressionSettings_ARM struct {
	// ContentTypesToCompress: List of content types on which compression applies. The value should be a valid MIME type.
	ContentTypesToCompress []string `json:"contentTypesToCompress,omitempty"`

	// IsCompressionEnabled: Indicates whether content compression is enabled on AzureFrontDoor. Default value is false. If
	// compression is enabled, content will be served as compressed if user requests for a compressed version. Content won't be
	// compressed on AzureFrontDoor when requested content is smaller than 1 byte or larger than 1 MB.
	IsCompressionEnabled *bool `json:"isCompressionEnabled,omitempty"`
}