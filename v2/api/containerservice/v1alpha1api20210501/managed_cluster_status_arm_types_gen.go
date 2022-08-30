// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210501

import "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

// Deprecated version of ManagedCluster_STATUS. Use v1beta20210501.ManagedCluster_STATUS instead
type ManagedCluster_STATUSARM struct {
	ExtendedLocation *ExtendedLocation_STATUSARM         `json:"extendedLocation,omitempty"`
	Id               *string                             `json:"id,omitempty"`
	Identity         *ManagedClusterIdentity_STATUSARM   `json:"identity,omitempty"`
	Location         *string                             `json:"location,omitempty"`
	Name             *string                             `json:"name,omitempty"`
	Properties       *ManagedClusterProperties_STATUSARM `json:"properties,omitempty"`
	Sku              *ManagedClusterSKU_STATUSARM        `json:"sku,omitempty"`
	Tags             map[string]string                   `json:"tags,omitempty"`
	Type             *string                             `json:"type,omitempty"`
}

// Deprecated version of ExtendedLocation_STATUS. Use v1beta20210501.ExtendedLocation_STATUS instead
type ExtendedLocation_STATUSARM struct {
	Name *string                      `json:"name,omitempty"`
	Type *ExtendedLocationType_STATUS `json:"type,omitempty"`
}

// Deprecated version of ManagedClusterIdentity_STATUS. Use v1beta20210501.ManagedClusterIdentity_STATUS instead
type ManagedClusterIdentity_STATUSARM struct {
	PrincipalId            *string                                                            `json:"principalId,omitempty"`
	TenantId               *string                                                            `json:"tenantId,omitempty"`
	Type                   *ManagedClusterIdentity_STATUS_Type                                `json:"type,omitempty"`
	UserAssignedIdentities map[string]ManagedClusterIdentity_STATUS_UserAssignedIdentitiesARM `json:"userAssignedIdentities,omitempty"`
}

// Deprecated version of ManagedClusterProperties_STATUS. Use v1beta20210501.ManagedClusterProperties_STATUS instead
type ManagedClusterProperties_STATUSARM struct {
	AadProfile              *ManagedClusterAADProfile_STATUSARM                   `json:"aadProfile,omitempty"`
	AddonProfiles           *v1.JSON                                              `json:"addonProfiles,omitempty"`
	AgentPoolProfiles       []ManagedClusterAgentPoolProfile_STATUSARM            `json:"agentPoolProfiles,omitempty"`
	ApiServerAccessProfile  *ManagedClusterAPIServerAccessProfile_STATUSARM       `json:"apiServerAccessProfile,omitempty"`
	AutoScalerProfile       *ManagedClusterProperties_STATUS_AutoScalerProfileARM `json:"autoScalerProfile,omitempty"`
	AutoUpgradeProfile      *ManagedClusterAutoUpgradeProfile_STATUSARM           `json:"autoUpgradeProfile,omitempty"`
	AzurePortalFQDN         *string                                               `json:"azurePortalFQDN,omitempty"`
	DisableLocalAccounts    *bool                                                 `json:"disableLocalAccounts,omitempty"`
	DiskEncryptionSetID     *string                                               `json:"diskEncryptionSetID,omitempty"`
	DnsPrefix               *string                                               `json:"dnsPrefix,omitempty"`
	EnablePodSecurityPolicy *bool                                                 `json:"enablePodSecurityPolicy,omitempty"`
	EnableRBAC              *bool                                                 `json:"enableRBAC,omitempty"`
	Fqdn                    *string                                               `json:"fqdn,omitempty"`
	FqdnSubdomain           *string                                               `json:"fqdnSubdomain,omitempty"`
	HttpProxyConfig         *ManagedClusterHTTPProxyConfig_STATUSARM              `json:"httpProxyConfig,omitempty"`
	IdentityProfile         *v1.JSON                                              `json:"identityProfile,omitempty"`
	KubernetesVersion       *string                                               `json:"kubernetesVersion,omitempty"`
	LinuxProfile            *ContainerServiceLinuxProfile_STATUSARM               `json:"linuxProfile,omitempty"`
	MaxAgentPools           *int                                                  `json:"maxAgentPools,omitempty"`
	NetworkProfile          *ContainerServiceNetworkProfile_STATUSARM             `json:"networkProfile,omitempty"`
	NodeResourceGroup       *string                                               `json:"nodeResourceGroup,omitempty"`
	PodIdentityProfile      *ManagedClusterPodIdentityProfile_STATUSARM           `json:"podIdentityProfile,omitempty"`
	PowerState              *PowerState_STATUSARM                                 `json:"powerState,omitempty"`
	PrivateFQDN             *string                                               `json:"privateFQDN,omitempty"`
	PrivateLinkResources    []PrivateLinkResource_STATUSARM                       `json:"privateLinkResources,omitempty"`
	ProvisioningState       *string                                               `json:"provisioningState,omitempty"`
	ServicePrincipalProfile *ManagedClusterServicePrincipalProfile_STATUSARM      `json:"servicePrincipalProfile,omitempty"`
	WindowsProfile          *ManagedClusterWindowsProfile_STATUSARM               `json:"windowsProfile,omitempty"`
}

// Deprecated version of ManagedClusterSKU_STATUS. Use v1beta20210501.ManagedClusterSKU_STATUS instead
type ManagedClusterSKU_STATUSARM struct {
	Name *ManagedClusterSKU_STATUS_Name `json:"name,omitempty"`
	Tier *ManagedClusterSKU_STATUS_Tier `json:"tier,omitempty"`
}

// Deprecated version of ContainerServiceLinuxProfile_STATUS. Use v1beta20210501.ContainerServiceLinuxProfile_STATUS instead
type ContainerServiceLinuxProfile_STATUSARM struct {
	AdminUsername *string                                     `json:"adminUsername,omitempty"`
	Ssh           *ContainerServiceSshConfiguration_STATUSARM `json:"ssh,omitempty"`
}

// Deprecated version of ContainerServiceNetworkProfile_STATUS. Use v1beta20210501.ContainerServiceNetworkProfile_STATUS instead
type ContainerServiceNetworkProfile_STATUSARM struct {
	DnsServiceIP        *string                                                `json:"dnsServiceIP,omitempty"`
	DockerBridgeCidr    *string                                                `json:"dockerBridgeCidr,omitempty"`
	LoadBalancerProfile *ManagedClusterLoadBalancerProfile_STATUSARM           `json:"loadBalancerProfile,omitempty"`
	LoadBalancerSku     *ContainerServiceNetworkProfile_STATUS_LoadBalancerSku `json:"loadBalancerSku,omitempty"`
	NetworkMode         *ContainerServiceNetworkProfile_STATUS_NetworkMode     `json:"networkMode,omitempty"`
	NetworkPlugin       *ContainerServiceNetworkProfile_STATUS_NetworkPlugin   `json:"networkPlugin,omitempty"`
	NetworkPolicy       *ContainerServiceNetworkProfile_STATUS_NetworkPolicy   `json:"networkPolicy,omitempty"`
	OutboundType        *ContainerServiceNetworkProfile_STATUS_OutboundType    `json:"outboundType,omitempty"`
	PodCidr             *string                                                `json:"podCidr,omitempty"`
	ServiceCidr         *string                                                `json:"serviceCidr,omitempty"`
}

// Deprecated version of ExtendedLocationType_STATUS. Use v1beta20210501.ExtendedLocationType_STATUS instead
type ExtendedLocationType_STATUS string

const ExtendedLocationType_STATUS_EdgeZone = ExtendedLocationType_STATUS("EdgeZone")

// Deprecated version of ManagedClusterAADProfile_STATUS. Use v1beta20210501.ManagedClusterAADProfile_STATUS instead
type ManagedClusterAADProfile_STATUSARM struct {
	AdminGroupObjectIDs []string `json:"adminGroupObjectIDs,omitempty"`
	ClientAppID         *string  `json:"clientAppID,omitempty"`
	EnableAzureRBAC     *bool    `json:"enableAzureRBAC,omitempty"`
	Managed             *bool    `json:"managed,omitempty"`
	ServerAppID         *string  `json:"serverAppID,omitempty"`
	ServerAppSecret     *string  `json:"serverAppSecret,omitempty"`
	TenantID            *string  `json:"tenantID,omitempty"`
}

// Deprecated version of ManagedClusterAgentPoolProfile_STATUS. Use v1beta20210501.ManagedClusterAgentPoolProfile_STATUS instead
type ManagedClusterAgentPoolProfile_STATUSARM struct {
	AvailabilityZones         []string                            `json:"availabilityZones,omitempty"`
	Count                     *int                                `json:"count,omitempty"`
	EnableAutoScaling         *bool                               `json:"enableAutoScaling,omitempty"`
	EnableEncryptionAtHost    *bool                               `json:"enableEncryptionAtHost,omitempty"`
	EnableFIPS                *bool                               `json:"enableFIPS,omitempty"`
	EnableNodePublicIP        *bool                               `json:"enableNodePublicIP,omitempty"`
	EnableUltraSSD            *bool                               `json:"enableUltraSSD,omitempty"`
	GpuInstanceProfile        *GPUInstanceProfile_STATUS          `json:"gpuInstanceProfile,omitempty"`
	KubeletConfig             *KubeletConfig_STATUSARM            `json:"kubeletConfig,omitempty"`
	KubeletDiskType           *KubeletDiskType_STATUS             `json:"kubeletDiskType,omitempty"`
	LinuxOSConfig             *LinuxOSConfig_STATUSARM            `json:"linuxOSConfig,omitempty"`
	MaxCount                  *int                                `json:"maxCount,omitempty"`
	MaxPods                   *int                                `json:"maxPods,omitempty"`
	MinCount                  *int                                `json:"minCount,omitempty"`
	Mode                      *AgentPoolMode_STATUS               `json:"mode,omitempty"`
	Name                      *string                             `json:"name,omitempty"`
	NodeImageVersion          *string                             `json:"nodeImageVersion,omitempty"`
	NodeLabels                map[string]string                   `json:"nodeLabels,omitempty"`
	NodePublicIPPrefixID      *string                             `json:"nodePublicIPPrefixID,omitempty"`
	NodeTaints                []string                            `json:"nodeTaints,omitempty"`
	OrchestratorVersion       *string                             `json:"orchestratorVersion,omitempty"`
	OsDiskSizeGB              *int                                `json:"osDiskSizeGB,omitempty"`
	OsDiskType                *OSDiskType_STATUS                  `json:"osDiskType,omitempty"`
	OsSKU                     *OSSKU_STATUS                       `json:"osSKU,omitempty"`
	OsType                    *OSType_STATUS                      `json:"osType,omitempty"`
	PodSubnetID               *string                             `json:"podSubnetID,omitempty"`
	PowerState                *PowerState_STATUSARM               `json:"powerState,omitempty"`
	ProvisioningState         *string                             `json:"provisioningState,omitempty"`
	ProximityPlacementGroupID *string                             `json:"proximityPlacementGroupID,omitempty"`
	ScaleSetEvictionPolicy    *ScaleSetEvictionPolicy_STATUS      `json:"scaleSetEvictionPolicy,omitempty"`
	ScaleSetPriority          *ScaleSetPriority_STATUS            `json:"scaleSetPriority,omitempty"`
	SpotMaxPrice              *float64                            `json:"spotMaxPrice,omitempty"`
	Tags                      map[string]string                   `json:"tags,omitempty"`
	Type                      *AgentPoolType_STATUS               `json:"type,omitempty"`
	UpgradeSettings           *AgentPoolUpgradeSettings_STATUSARM `json:"upgradeSettings,omitempty"`
	VmSize                    *string                             `json:"vmSize,omitempty"`
	VnetSubnetID              *string                             `json:"vnetSubnetID,omitempty"`
}

// Deprecated version of ManagedClusterAPIServerAccessProfile_STATUS. Use v1beta20210501.ManagedClusterAPIServerAccessProfile_STATUS instead
type ManagedClusterAPIServerAccessProfile_STATUSARM struct {
	AuthorizedIPRanges             []string `json:"authorizedIPRanges,omitempty"`
	EnablePrivateCluster           *bool    `json:"enablePrivateCluster,omitempty"`
	EnablePrivateClusterPublicFQDN *bool    `json:"enablePrivateClusterPublicFQDN,omitempty"`
	PrivateDNSZone                 *string  `json:"privateDNSZone,omitempty"`
}

// Deprecated version of ManagedClusterAutoUpgradeProfile_STATUS. Use v1beta20210501.ManagedClusterAutoUpgradeProfile_STATUS instead
type ManagedClusterAutoUpgradeProfile_STATUSARM struct {
	UpgradeChannel *ManagedClusterAutoUpgradeProfile_STATUS_UpgradeChannel `json:"upgradeChannel,omitempty"`
}

// Deprecated version of ManagedClusterHTTPProxyConfig_STATUS. Use v1beta20210501.ManagedClusterHTTPProxyConfig_STATUS instead
type ManagedClusterHTTPProxyConfig_STATUSARM struct {
	HttpProxy  *string  `json:"httpProxy,omitempty"`
	HttpsProxy *string  `json:"httpsProxy,omitempty"`
	NoProxy    []string `json:"noProxy,omitempty"`
	TrustedCa  *string  `json:"trustedCa,omitempty"`
}

// Deprecated version of ManagedClusterIdentity_STATUS_Type. Use v1beta20210501.ManagedClusterIdentity_STATUS_Type instead
type ManagedClusterIdentity_STATUS_Type string

const (
	ManagedClusterIdentity_STATUS_Type_None           = ManagedClusterIdentity_STATUS_Type("None")
	ManagedClusterIdentity_STATUS_Type_SystemAssigned = ManagedClusterIdentity_STATUS_Type("SystemAssigned")
	ManagedClusterIdentity_STATUS_Type_UserAssigned   = ManagedClusterIdentity_STATUS_Type("UserAssigned")
)

// Deprecated version of ManagedClusterIdentity_STATUS_UserAssignedIdentities. Use v1beta20210501.ManagedClusterIdentity_STATUS_UserAssignedIdentities instead
type ManagedClusterIdentity_STATUS_UserAssignedIdentitiesARM struct {
	ClientId    *string `json:"clientId,omitempty"`
	PrincipalId *string `json:"principalId,omitempty"`
}

// Deprecated version of ManagedClusterPodIdentityProfile_STATUS. Use v1beta20210501.ManagedClusterPodIdentityProfile_STATUS instead
type ManagedClusterPodIdentityProfile_STATUSARM struct {
	AllowNetworkPluginKubenet      *bool                                          `json:"allowNetworkPluginKubenet,omitempty"`
	Enabled                        *bool                                          `json:"enabled,omitempty"`
	UserAssignedIdentities         []ManagedClusterPodIdentity_STATUSARM          `json:"userAssignedIdentities,omitempty"`
	UserAssignedIdentityExceptions []ManagedClusterPodIdentityException_STATUSARM `json:"userAssignedIdentityExceptions,omitempty"`
}

// Deprecated version of ManagedClusterProperties_STATUS_AutoScalerProfile. Use v1beta20210501.ManagedClusterProperties_STATUS_AutoScalerProfile instead
type ManagedClusterProperties_STATUS_AutoScalerProfileARM struct {
	BalanceSimilarNodeGroups      *string                                                     `json:"balance-similar-node-groups,omitempty"`
	Expander                      *ManagedClusterProperties_STATUS_AutoScalerProfile_Expander `json:"expander,omitempty"`
	MaxEmptyBulkDelete            *string                                                     `json:"max-empty-bulk-delete,omitempty"`
	MaxGracefulTerminationSec     *string                                                     `json:"max-graceful-termination-sec,omitempty"`
	MaxNodeProvisionTime          *string                                                     `json:"max-node-provision-time,omitempty"`
	MaxTotalUnreadyPercentage     *string                                                     `json:"max-total-unready-percentage,omitempty"`
	NewPodScaleUpDelay            *string                                                     `json:"new-pod-scale-up-delay,omitempty"`
	OkTotalUnreadyCount           *string                                                     `json:"ok-total-unready-count,omitempty"`
	ScaleDownDelayAfterAdd        *string                                                     `json:"scale-down-delay-after-add,omitempty"`
	ScaleDownDelayAfterDelete     *string                                                     `json:"scale-down-delay-after-delete,omitempty"`
	ScaleDownDelayAfterFailure    *string                                                     `json:"scale-down-delay-after-failure,omitempty"`
	ScaleDownUnneededTime         *string                                                     `json:"scale-down-unneeded-time,omitempty"`
	ScaleDownUnreadyTime          *string                                                     `json:"scale-down-unready-time,omitempty"`
	ScaleDownUtilizationThreshold *string                                                     `json:"scale-down-utilization-threshold,omitempty"`
	ScanInterval                  *string                                                     `json:"scan-interval,omitempty"`
	SkipNodesWithLocalStorage     *string                                                     `json:"skip-nodes-with-local-storage,omitempty"`
	SkipNodesWithSystemPods       *string                                                     `json:"skip-nodes-with-system-pods,omitempty"`
}

// Deprecated version of ManagedClusterServicePrincipalProfile_STATUS. Use v1beta20210501.ManagedClusterServicePrincipalProfile_STATUS instead
type ManagedClusterServicePrincipalProfile_STATUSARM struct {
	ClientId *string `json:"clientId,omitempty"`
	Secret   *string `json:"secret,omitempty"`
}

// Deprecated version of ManagedClusterSKU_STATUS_Name. Use v1beta20210501.ManagedClusterSKU_STATUS_Name instead
type ManagedClusterSKU_STATUS_Name string

const ManagedClusterSKU_STATUS_Name_Basic = ManagedClusterSKU_STATUS_Name("Basic")

// Deprecated version of ManagedClusterSKU_STATUS_Tier. Use v1beta20210501.ManagedClusterSKU_STATUS_Tier instead
type ManagedClusterSKU_STATUS_Tier string

const (
	ManagedClusterSKU_STATUS_Tier_Free = ManagedClusterSKU_STATUS_Tier("Free")
	ManagedClusterSKU_STATUS_Tier_Paid = ManagedClusterSKU_STATUS_Tier("Paid")
)

// Deprecated version of ManagedClusterWindowsProfile_STATUS. Use v1beta20210501.ManagedClusterWindowsProfile_STATUS instead
type ManagedClusterWindowsProfile_STATUSARM struct {
	AdminPassword  *string                                          `json:"adminPassword,omitempty"`
	AdminUsername  *string                                          `json:"adminUsername,omitempty"`
	EnableCSIProxy *bool                                            `json:"enableCSIProxy,omitempty"`
	LicenseType    *ManagedClusterWindowsProfile_STATUS_LicenseType `json:"licenseType,omitempty"`
}

// Deprecated version of PrivateLinkResource_STATUS. Use v1beta20210501.PrivateLinkResource_STATUS instead
type PrivateLinkResource_STATUSARM struct {
	GroupId              *string  `json:"groupId,omitempty"`
	Id                   *string  `json:"id,omitempty"`
	Name                 *string  `json:"name,omitempty"`
	PrivateLinkServiceID *string  `json:"privateLinkServiceID,omitempty"`
	RequiredMembers      []string `json:"requiredMembers,omitempty"`
	Type                 *string  `json:"type,omitempty"`
}

// Deprecated version of ContainerServiceSshConfiguration_STATUS. Use v1beta20210501.ContainerServiceSshConfiguration_STATUS instead
type ContainerServiceSshConfiguration_STATUSARM struct {
	PublicKeys []ContainerServiceSshPublicKey_STATUSARM `json:"publicKeys,omitempty"`
}

// Deprecated version of ManagedClusterLoadBalancerProfile_STATUS. Use v1beta20210501.ManagedClusterLoadBalancerProfile_STATUS instead
type ManagedClusterLoadBalancerProfile_STATUSARM struct {
	AllocatedOutboundPorts *int                                                            `json:"allocatedOutboundPorts,omitempty"`
	EffectiveOutboundIPs   []ResourceReference_STATUSARM                                   `json:"effectiveOutboundIPs,omitempty"`
	IdleTimeoutInMinutes   *int                                                            `json:"idleTimeoutInMinutes,omitempty"`
	ManagedOutboundIPs     *ManagedClusterLoadBalancerProfile_STATUS_ManagedOutboundIPsARM `json:"managedOutboundIPs,omitempty"`
	OutboundIPPrefixes     *ManagedClusterLoadBalancerProfile_STATUS_OutboundIPPrefixesARM `json:"outboundIPPrefixes,omitempty"`
	OutboundIPs            *ManagedClusterLoadBalancerProfile_STATUS_OutboundIPsARM        `json:"outboundIPs,omitempty"`
}

// Deprecated version of ManagedClusterPodIdentity_STATUS. Use v1beta20210501.ManagedClusterPodIdentity_STATUS instead
type ManagedClusterPodIdentity_STATUSARM struct {
	BindingSelector   *string                                               `json:"bindingSelector,omitempty"`
	Identity          *UserAssignedIdentity_STATUSARM                       `json:"identity,omitempty"`
	Name              *string                                               `json:"name,omitempty"`
	Namespace         *string                                               `json:"namespace,omitempty"`
	ProvisioningInfo  *ManagedClusterPodIdentity_STATUS_ProvisioningInfoARM `json:"provisioningInfo,omitempty"`
	ProvisioningState *ManagedClusterPodIdentity_STATUS_ProvisioningState   `json:"provisioningState,omitempty"`
}

// Deprecated version of ManagedClusterPodIdentityException_STATUS. Use v1beta20210501.ManagedClusterPodIdentityException_STATUS instead
type ManagedClusterPodIdentityException_STATUSARM struct {
	Name      *string           `json:"name,omitempty"`
	Namespace *string           `json:"namespace,omitempty"`
	PodLabels map[string]string `json:"podLabels,omitempty"`
}

// Deprecated version of ContainerServiceSshPublicKey_STATUS. Use v1beta20210501.ContainerServiceSshPublicKey_STATUS instead
type ContainerServiceSshPublicKey_STATUSARM struct {
	KeyData *string `json:"keyData,omitempty"`
}

// Deprecated version of ManagedClusterLoadBalancerProfile_STATUS_ManagedOutboundIPs. Use v1beta20210501.ManagedClusterLoadBalancerProfile_STATUS_ManagedOutboundIPs instead
type ManagedClusterLoadBalancerProfile_STATUS_ManagedOutboundIPsARM struct {
	Count *int `json:"count,omitempty"`
}

// Deprecated version of ManagedClusterLoadBalancerProfile_STATUS_OutboundIPPrefixes. Use v1beta20210501.ManagedClusterLoadBalancerProfile_STATUS_OutboundIPPrefixes instead
type ManagedClusterLoadBalancerProfile_STATUS_OutboundIPPrefixesARM struct {
	PublicIPPrefixes []ResourceReference_STATUSARM `json:"publicIPPrefixes,omitempty"`
}

// Deprecated version of ManagedClusterLoadBalancerProfile_STATUS_OutboundIPs. Use v1beta20210501.ManagedClusterLoadBalancerProfile_STATUS_OutboundIPs instead
type ManagedClusterLoadBalancerProfile_STATUS_OutboundIPsARM struct {
	PublicIPs []ResourceReference_STATUSARM `json:"publicIPs,omitempty"`
}

// Deprecated version of ManagedClusterPodIdentity_STATUS_ProvisioningInfo. Use v1beta20210501.ManagedClusterPodIdentity_STATUS_ProvisioningInfo instead
type ManagedClusterPodIdentity_STATUS_ProvisioningInfoARM struct {
	Error *ManagedClusterPodIdentityProvisioningError_STATUSARM `json:"error,omitempty"`
}

// Deprecated version of ResourceReference_STATUS. Use v1beta20210501.ResourceReference_STATUS instead
type ResourceReference_STATUSARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of UserAssignedIdentity_STATUS. Use v1beta20210501.UserAssignedIdentity_STATUS instead
type UserAssignedIdentity_STATUSARM struct {
	ClientId   *string `json:"clientId,omitempty"`
	ObjectId   *string `json:"objectId,omitempty"`
	ResourceId *string `json:"resourceId,omitempty"`
}

// Deprecated version of ManagedClusterPodIdentityProvisioningError_STATUS. Use v1beta20210501.ManagedClusterPodIdentityProvisioningError_STATUS instead
type ManagedClusterPodIdentityProvisioningError_STATUSARM struct {
	Error *ManagedClusterPodIdentityProvisioningErrorBody_STATUSARM `json:"error,omitempty"`
}

// Deprecated version of ManagedClusterPodIdentityProvisioningErrorBody_STATUS. Use v1beta20210501.ManagedClusterPodIdentityProvisioningErrorBody_STATUS instead
type ManagedClusterPodIdentityProvisioningErrorBody_STATUSARM struct {
	Code    *string                                                             `json:"code,omitempty"`
	Details []ManagedClusterPodIdentityProvisioningErrorBody_STATUS_UnrolledARM `json:"details,omitempty"`
	Message *string                                                             `json:"message,omitempty"`
	Target  *string                                                             `json:"target,omitempty"`
}

// Deprecated version of ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled. Use v1beta20210501.ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled instead
type ManagedClusterPodIdentityProvisioningErrorBody_STATUS_UnrolledARM struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Target  *string `json:"target,omitempty"`
}