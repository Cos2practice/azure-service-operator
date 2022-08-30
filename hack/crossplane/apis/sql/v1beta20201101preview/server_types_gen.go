// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101preview

import (
	"github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:rbac:groups=sql.azure.com,resources=servers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sql.azure.com,resources={servers/status,servers/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// Generated from: https://schema.management.azure.com/schemas/2020-11-01-preview/Microsoft.Sql.json#/resourceDefinitions/servers
type Server struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Servers_Spec  `json:"spec,omitempty"`
	Status            Server_STATUS `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// Generated from: https://schema.management.azure.com/schemas/2020-11-01-preview/Microsoft.Sql.json#/resourceDefinitions/servers
type ServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Server `json:"items"`
}

// +kubebuilder:validation:Enum={"2020-11-01-preview"}
type APIVersion string

const APIVersion_Value = APIVersion("2020-11-01-preview")

type Server_STATUS struct {
	v1alpha1.ResourceStatus `json:",inline,omitempty"`
	AtProvider              ServerObservation `json:"atProvider,omitempty"`
}

type Servers_Spec struct {
	v1alpha1.ResourceSpec `json:",inline,omitempty"`
	ForProvider           ServersParameters `json:"forProvider,omitempty"`
}

type ServerObservation struct {
	// AdministratorLogin: Administrator username for the server. Once created it cannot be changed.
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	// AdministratorLoginPassword: The administrator login password (required for server creation).
	AdministratorLoginPassword *string `json:"administratorLoginPassword,omitempty"`

	// Administrators: The Azure Active Directory identity of the server.
	Administrators *ServerExternalAdministrator_STATUS `json:"administrators,omitempty"`

	// FullyQualifiedDomainName: The fully qualified domain name of the server.
	FullyQualifiedDomainName *string `json:"fullyQualifiedDomainName,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Identity: The Azure Active Directory identity of the server.
	Identity *ResourceIdentity_STATUS `json:"identity,omitempty"`

	// KeyId: A CMK URI of the key to use for encryption.
	KeyId *string `json:"keyId,omitempty"`

	// Kind: Kind of sql server. This is metadata used for the Azure portal experience.
	Kind *string `json:"kind,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// MinimalTlsVersion: Minimal TLS version. Allowed values: '1.0', '1.1', '1.2'
	MinimalTlsVersion *string `json:"minimalTlsVersion,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// PrimaryUserAssignedIdentityId: The resource id of a user assigned identity to be used by default.
	PrimaryUserAssignedIdentityId *string `json:"primaryUserAssignedIdentityId,omitempty"`

	// PrivateEndpointConnections: List of private endpoint connections on a server
	PrivateEndpointConnections []ServerPrivateEndpointConnection_STATUS `json:"privateEndpointConnections,omitempty"`

	// PublicNetworkAccess: Whether or not public endpoint access is allowed for this server.  Value is optional but if passed
	// in, must be 'Enabled' or 'Disabled'
	PublicNetworkAccess *ServerProperties_STATUS_PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// State: The state of the server.
	State *string `json:"state,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`

	// Version: The version of the server.
	Version *string `json:"version,omitempty"`

	// WorkspaceFeature: Whether or not existing server has a workspace created and if it allows connection from workspace
	WorkspaceFeature *ServerProperties_STATUS_WorkspaceFeature `json:"workspaceFeature,omitempty"`
}

type ServersParameters struct {
	// AdministratorLogin: Administrator username for the server. Once created it cannot be changed.
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	// AdministratorLoginPassword: The administrator login password (required for server creation).
	AdministratorLoginPassword *string `json:"administratorLoginPassword,omitempty"`

	// Administrators: Properties of a active directory administrator.
	Administrators *ServerExternalAdministrator `json:"administrators,omitempty"`

	// Identity: Azure Active Directory identity configuration for a resource.
	Identity *ResourceIdentity `json:"identity,omitempty"`

	// KeyId: A CMK URI of the key to use for encryption.
	KeyId *string `json:"keyId,omitempty"`

	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// MinimalTlsVersion: Minimal TLS version. Allowed values: '1.0', '1.1', '1.2'
	MinimalTlsVersion *string `json:"minimalTlsVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Name: The name of the server.
	Name string `json:"name,omitempty"`

	// PrimaryUserAssignedIdentityId: The resource id of a user assigned identity to be used by default.
	PrimaryUserAssignedIdentityId *string `json:"primaryUserAssignedIdentityId,omitempty"`

	// PublicNetworkAccess: Whether or not public endpoint access is allowed for this server.  Value is optional but if passed
	// in, must be 'Enabled' or 'Disabled'.
	PublicNetworkAccess       *ServerProperties_PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`
	ResourceGroupName         string                                `json:"resourceGroupName,omitempty"`
	ResourceGroupNameRef      *v1alpha1.Reference                   `json:"resourceGroupNameRef,omitempty"`
	ResourceGroupNameSelector *v1alpha1.Selector                    `json:"resourceGroupNameSelector,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`

	// Version: The version of the server.
	Version *string `json:"version,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2020-11-01-preview/Microsoft.Sql.json#/definitions/ResourceIdentity
type ResourceIdentity struct {
	// Type: The identity type. Set this to 'SystemAssigned' in order to automatically create and assign an Azure Active
	// Directory principal for the resource.
	Type *ResourceIdentity_Type `json:"type,omitempty"`
}

type ResourceIdentity_STATUS struct {
	// PrincipalId: The Azure Active Directory principal id.
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: The Azure Active Directory tenant id.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: The identity type. Set this to 'SystemAssigned' in order to automatically create and assign an Azure Active
	// Directory principal for the resource.
	Type *ResourceIdentity_STATUS_Type `json:"type,omitempty"`

	// UserAssignedIdentities: The resource ids of the user assigned identities to use
	UserAssignedIdentities map[string]UserIdentity_STATUS `json:"userAssignedIdentities,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2020-11-01-preview/Microsoft.Sql.json#/definitions/ServerExternalAdministrator
type ServerExternalAdministrator struct {
	// AdministratorType: Type of the sever administrator.
	AdministratorType *ServerExternalAdministrator_AdministratorType `json:"administratorType,omitempty"`

	// AzureADOnlyAuthentication: Azure Active Directory only Authentication enabled.
	AzureADOnlyAuthentication *bool `json:"azureADOnlyAuthentication,omitempty"`

	// Login: Login name of the server administrator.
	Login *string `json:"login,omitempty"`

	// PrincipalType: Principal Type of the sever administrator.
	PrincipalType *ServerExternalAdministrator_PrincipalType `json:"principalType,omitempty"`

	// +kubebuilder:validation:Pattern="^[0-9a-fA-F]{8}(-[0-9a-fA-F]{4}){3}-[0-9a-fA-F]{12}$"
	// Sid: SID (object ID) of the server administrator.
	Sid *string `json:"sid,omitempty"`

	// +kubebuilder:validation:Pattern="^[0-9a-fA-F]{8}(-[0-9a-fA-F]{4}){3}-[0-9a-fA-F]{12}$"
	// TenantId: Tenant ID of the administrator.
	TenantId *string `json:"tenantId,omitempty"`
}

type ServerExternalAdministrator_STATUS struct {
	// AdministratorType: Type of the sever administrator.
	AdministratorType *ServerExternalAdministrator_STATUS_AdministratorType `json:"administratorType,omitempty"`

	// AzureADOnlyAuthentication: Azure Active Directory only Authentication enabled.
	AzureADOnlyAuthentication *bool `json:"azureADOnlyAuthentication,omitempty"`

	// Login: Login name of the server administrator.
	Login *string `json:"login,omitempty"`

	// PrincipalType: Principal Type of the sever administrator.
	PrincipalType *ServerExternalAdministrator_STATUS_PrincipalType `json:"principalType,omitempty"`

	// Sid: SID (object ID) of the server administrator.
	Sid *string `json:"sid,omitempty"`

	// TenantId: Tenant ID of the administrator.
	TenantId *string `json:"tenantId,omitempty"`
}

type ServerPrivateEndpointConnection_STATUS struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Properties: Private endpoint connection properties
	Properties *PrivateEndpointConnectionProperties_STATUS `json:"properties,omitempty"`
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ServerProperties_PublicNetworkAccess string

const (
	ServerProperties_PublicNetworkAccess_Disabled = ServerProperties_PublicNetworkAccess("Disabled")
	ServerProperties_PublicNetworkAccess_Enabled  = ServerProperties_PublicNetworkAccess("Enabled")
)

type ServerProperties_STATUS_PublicNetworkAccess string

const (
	ServerProperties_STATUS_PublicNetworkAccess_Disabled = ServerProperties_STATUS_PublicNetworkAccess("Disabled")
	ServerProperties_STATUS_PublicNetworkAccess_Enabled  = ServerProperties_STATUS_PublicNetworkAccess("Enabled")
)

type ServerProperties_STATUS_WorkspaceFeature string

const (
	ServerProperties_STATUS_WorkspaceFeature_Connected    = ServerProperties_STATUS_WorkspaceFeature("Connected")
	ServerProperties_STATUS_WorkspaceFeature_Disconnected = ServerProperties_STATUS_WorkspaceFeature("Disconnected")
)

type PrivateEndpointConnectionProperties_STATUS struct {
	// PrivateEndpoint: Private endpoint which the connection belongs to.
	PrivateEndpoint *PrivateEndpointProperty_STATUS `json:"privateEndpoint,omitempty"`

	// PrivateLinkServiceConnectionState: Connection state of the private endpoint connection.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateProperty_STATUS `json:"privateLinkServiceConnectionState,omitempty"`

	// ProvisioningState: State of the private endpoint connection.
	ProvisioningState *PrivateEndpointConnectionProperties_STATUS_ProvisioningState `json:"provisioningState,omitempty"`
}

type ResourceIdentity_STATUS_Type string

const (
	ResourceIdentity_STATUS_Type_None                       = ResourceIdentity_STATUS_Type("None")
	ResourceIdentity_STATUS_Type_SystemAssigned             = ResourceIdentity_STATUS_Type("SystemAssigned")
	ResourceIdentity_STATUS_Type_SystemAssignedUserAssigned = ResourceIdentity_STATUS_Type("SystemAssigned,UserAssigned")
	ResourceIdentity_STATUS_Type_UserAssigned               = ResourceIdentity_STATUS_Type("UserAssigned")
)

// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned,UserAssigned","UserAssigned"}
type ResourceIdentity_Type string

const (
	ResourceIdentity_Type_None                       = ResourceIdentity_Type("None")
	ResourceIdentity_Type_SystemAssigned             = ResourceIdentity_Type("SystemAssigned")
	ResourceIdentity_Type_SystemAssignedUserAssigned = ResourceIdentity_Type("SystemAssigned,UserAssigned")
	ResourceIdentity_Type_UserAssigned               = ResourceIdentity_Type("UserAssigned")
)

// +kubebuilder:validation:Enum={"ActiveDirectory"}
type ServerExternalAdministrator_AdministratorType string

const ServerExternalAdministrator_AdministratorType_ActiveDirectory = ServerExternalAdministrator_AdministratorType("ActiveDirectory")

// +kubebuilder:validation:Enum={"Application","Group","User"}
type ServerExternalAdministrator_PrincipalType string

const (
	ServerExternalAdministrator_PrincipalType_Application = ServerExternalAdministrator_PrincipalType("Application")
	ServerExternalAdministrator_PrincipalType_Group       = ServerExternalAdministrator_PrincipalType("Group")
	ServerExternalAdministrator_PrincipalType_User        = ServerExternalAdministrator_PrincipalType("User")
)

type ServerExternalAdministrator_STATUS_AdministratorType string

const ServerExternalAdministrator_STATUS_AdministratorType_ActiveDirectory = ServerExternalAdministrator_STATUS_AdministratorType("ActiveDirectory")

type ServerExternalAdministrator_STATUS_PrincipalType string

const (
	ServerExternalAdministrator_STATUS_PrincipalType_Application = ServerExternalAdministrator_STATUS_PrincipalType("Application")
	ServerExternalAdministrator_STATUS_PrincipalType_Group       = ServerExternalAdministrator_STATUS_PrincipalType("Group")
	ServerExternalAdministrator_STATUS_PrincipalType_User        = ServerExternalAdministrator_STATUS_PrincipalType("User")
)

type UserIdentity_STATUS struct {
	// ClientId: The Azure Active Directory client id.
	ClientId *string `json:"clientId,omitempty"`

	// PrincipalId: The Azure Active Directory principal id.
	PrincipalId *string `json:"principalId,omitempty"`
}

type PrivateEndpointConnectionProperties_STATUS_ProvisioningState string

const (
	PrivateEndpointConnectionProperties_STATUS_ProvisioningState_Approving = PrivateEndpointConnectionProperties_STATUS_ProvisioningState("Approving")
	PrivateEndpointConnectionProperties_STATUS_ProvisioningState_Dropping  = PrivateEndpointConnectionProperties_STATUS_ProvisioningState("Dropping")
	PrivateEndpointConnectionProperties_STATUS_ProvisioningState_Failed    = PrivateEndpointConnectionProperties_STATUS_ProvisioningState("Failed")
	PrivateEndpointConnectionProperties_STATUS_ProvisioningState_Ready     = PrivateEndpointConnectionProperties_STATUS_ProvisioningState("Ready")
	PrivateEndpointConnectionProperties_STATUS_ProvisioningState_Rejecting = PrivateEndpointConnectionProperties_STATUS_ProvisioningState("Rejecting")
)

type PrivateEndpointProperty_STATUS struct {
	// Id: Resource id of the private endpoint.
	Id *string `json:"id,omitempty"`
}

type PrivateLinkServiceConnectionStateProperty_STATUS struct {
	// ActionsRequired: The actions required for private link service connection.
	ActionsRequired *PrivateLinkServiceConnectionStateProperty_STATUS_ActionsRequired `json:"actionsRequired,omitempty"`

	// Description: The private link service connection description.
	Description *string `json:"description,omitempty"`

	// Status: The private link service connection status.
	Status *PrivateLinkServiceConnectionStateProperty_STATUS_Status `json:"status,omitempty"`
}

type PrivateLinkServiceConnectionStateProperty_STATUS_ActionsRequired string

const PrivateLinkServiceConnectionStateProperty_STATUS_ActionsRequired_None = PrivateLinkServiceConnectionStateProperty_STATUS_ActionsRequired("None")

type PrivateLinkServiceConnectionStateProperty_STATUS_Status string

const (
	PrivateLinkServiceConnectionStateProperty_STATUS_Status_Approved     = PrivateLinkServiceConnectionStateProperty_STATUS_Status("Approved")
	PrivateLinkServiceConnectionStateProperty_STATUS_Status_Disconnected = PrivateLinkServiceConnectionStateProperty_STATUS_Status("Disconnected")
	PrivateLinkServiceConnectionStateProperty_STATUS_Status_Pending      = PrivateLinkServiceConnectionStateProperty_STATUS_Status("Pending")
	PrivateLinkServiceConnectionStateProperty_STATUS_Status_Rejected     = PrivateLinkServiceConnectionStateProperty_STATUS_Status("Rejected")
)

func init() {
	SchemeBuilder.Register(&Server{}, &ServerList{})
}