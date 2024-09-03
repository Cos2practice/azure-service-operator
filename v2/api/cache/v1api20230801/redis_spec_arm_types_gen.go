// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230801

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Redis_Spec_ARM struct {
	// Identity: The identity of the resource.
	Identity *ManagedServiceIdentity_ARM `json:"identity,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Redis cache properties.
	Properties *RedisCreateProperties_ARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Zones: A list of availability zones denoting where the resource needs to come from.
	Zones []string `json:"zones,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Redis_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-08-01"
func (redis Redis_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (redis *Redis_Spec_ARM) GetName() string {
	return redis.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis"
func (redis *Redis_Spec_ARM) GetType() string {
	return "Microsoft.Cache/redis"
}

// Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity_ARM struct {
	// Type: Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type                   *ManagedServiceIdentityType_ARM            `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails_ARM `json:"userAssignedIdentities,omitempty"`
}

// Properties supplied to Create Redis operation.
type RedisCreateProperties_ARM struct {
	// EnableNonSslPort: Specifies whether the non-ssl Redis server port (6379) is enabled.
	EnableNonSslPort *bool `json:"enableNonSslPort,omitempty"`

	// MinimumTlsVersion: Optional: requires clients to use a specified TLS version (or higher) to connect (e,g, '1.0', '1.1',
	// '1.2')
	MinimumTlsVersion *RedisCreateProperties_MinimumTlsVersion_ARM `json:"minimumTlsVersion,omitempty"`

	// PublicNetworkAccess: Whether or not public endpoint access is allowed for this cache.  Value is optional but if passed
	// in, must be 'Enabled' or 'Disabled'. If 'Disabled', private endpoints are the exclusive access method. Default value is
	// 'Enabled'
	PublicNetworkAccess *RedisCreateProperties_PublicNetworkAccess_ARM `json:"publicNetworkAccess,omitempty"`

	// RedisConfiguration: All Redis Settings. Few possible keys:
	// rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value
	// etc.
	RedisConfiguration *RedisCreateProperties_RedisConfiguration_ARM `json:"redisConfiguration,omitempty"`

	// RedisVersion: Redis version. This should be in the form 'major[.minor]' (only 'major' is required) or the value 'latest'
	// which refers to the latest stable Redis version that is available. Supported versions: 4.0, 6.0 (latest). Default value
	// is 'latest'.
	RedisVersion *string `json:"redisVersion,omitempty"`

	// ReplicasPerMaster: The number of replicas to be created per primary.
	ReplicasPerMaster *int `json:"replicasPerMaster,omitempty"`

	// ReplicasPerPrimary: The number of replicas to be created per primary.
	ReplicasPerPrimary *int `json:"replicasPerPrimary,omitempty"`

	// ShardCount: The number of shards to be created on a Premium Cluster Cache.
	ShardCount *int `json:"shardCount,omitempty"`

	// Sku: The SKU of the Redis cache to deploy.
	Sku *Sku_ARM `json:"sku,omitempty"`

	// StaticIP: Static IP address. Optionally, may be specified when deploying a Redis cache inside an existing Azure Virtual
	// Network; auto assigned by default.
	StaticIP *string `json:"staticIP,omitempty"`
	SubnetId *string `json:"subnetId,omitempty"`

	// TenantSettings: A dictionary of tenant settings
	TenantSettings map[string]string `json:"tenantSettings,omitempty"`

	// UpdateChannel: Optional: Specifies the update channel for the monthly Redis updates your Redis Cache will receive.
	// Caches using 'Preview' update channel get latest Redis updates at least 4 weeks ahead of 'Stable' channel caches.
	// Default value is 'Stable'.
	UpdateChannel *RedisCreateProperties_UpdateChannel_ARM `json:"updateChannel,omitempty"`
}

// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned, UserAssigned","UserAssigned"}
type ManagedServiceIdentityType_ARM string

const (
	ManagedServiceIdentityType_ARM_None                       = ManagedServiceIdentityType_ARM("None")
	ManagedServiceIdentityType_ARM_SystemAssigned             = ManagedServiceIdentityType_ARM("SystemAssigned")
	ManagedServiceIdentityType_ARM_SystemAssignedUserAssigned = ManagedServiceIdentityType_ARM("SystemAssigned, UserAssigned")
	ManagedServiceIdentityType_ARM_UserAssigned               = ManagedServiceIdentityType_ARM("UserAssigned")
)

// Mapping from string to ManagedServiceIdentityType_ARM
var managedServiceIdentityType_ARM_Values = map[string]ManagedServiceIdentityType_ARM{
	"none":                         ManagedServiceIdentityType_ARM_None,
	"systemassigned":               ManagedServiceIdentityType_ARM_SystemAssigned,
	"systemassigned, userassigned": ManagedServiceIdentityType_ARM_SystemAssignedUserAssigned,
	"userassigned":                 ManagedServiceIdentityType_ARM_UserAssigned,
}

// +kubebuilder:validation:Enum={"1.0","1.1","1.2"}
type RedisCreateProperties_MinimumTlsVersion_ARM string

const (
	RedisCreateProperties_MinimumTlsVersion_ARM_10 = RedisCreateProperties_MinimumTlsVersion_ARM("1.0")
	RedisCreateProperties_MinimumTlsVersion_ARM_11 = RedisCreateProperties_MinimumTlsVersion_ARM("1.1")
	RedisCreateProperties_MinimumTlsVersion_ARM_12 = RedisCreateProperties_MinimumTlsVersion_ARM("1.2")
)

// Mapping from string to RedisCreateProperties_MinimumTlsVersion_ARM
var redisCreateProperties_MinimumTlsVersion_ARM_Values = map[string]RedisCreateProperties_MinimumTlsVersion_ARM{
	"1.0": RedisCreateProperties_MinimumTlsVersion_ARM_10,
	"1.1": RedisCreateProperties_MinimumTlsVersion_ARM_11,
	"1.2": RedisCreateProperties_MinimumTlsVersion_ARM_12,
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type RedisCreateProperties_PublicNetworkAccess_ARM string

const (
	RedisCreateProperties_PublicNetworkAccess_ARM_Disabled = RedisCreateProperties_PublicNetworkAccess_ARM("Disabled")
	RedisCreateProperties_PublicNetworkAccess_ARM_Enabled  = RedisCreateProperties_PublicNetworkAccess_ARM("Enabled")
)

// Mapping from string to RedisCreateProperties_PublicNetworkAccess_ARM
var redisCreateProperties_PublicNetworkAccess_ARM_Values = map[string]RedisCreateProperties_PublicNetworkAccess_ARM{
	"disabled": RedisCreateProperties_PublicNetworkAccess_ARM_Disabled,
	"enabled":  RedisCreateProperties_PublicNetworkAccess_ARM_Enabled,
}

type RedisCreateProperties_RedisConfiguration_ARM struct {
	// AadEnabled: Specifies whether AAD based authentication has been enabled or disabled for the cache
	AadEnabled *string `json:"aad-enabled,omitempty"`

	// AofBackupEnabled: Specifies whether the aof backup is enabled
	AofBackupEnabled *string `json:"aof-backup-enabled,omitempty"`

	// AofStorageConnectionString0: First storage account connection string
	AofStorageConnectionString0 *string `json:"aof-storage-connection-string-0,omitempty"`

	// AofStorageConnectionString1: Second storage account connection string
	AofStorageConnectionString1 *string `json:"aof-storage-connection-string-1,omitempty"`

	// Authnotrequired: Specifies whether the authentication is disabled. Setting this property is highly discouraged from
	// security point of view.
	Authnotrequired *string `json:"authnotrequired,omitempty"`

	// MaxfragmentationmemoryReserved: Value in megabytes reserved for fragmentation per shard
	MaxfragmentationmemoryReserved *string `json:"maxfragmentationmemory-reserved,omitempty"`

	// MaxmemoryDelta: Value in megabytes reserved for non-cache usage per shard e.g. failover.
	MaxmemoryDelta *string `json:"maxmemory-delta,omitempty"`

	// MaxmemoryPolicy: The eviction strategy used when your data won't fit within its memory limit.
	MaxmemoryPolicy *string `json:"maxmemory-policy,omitempty"`

	// MaxmemoryReserved: Value in megabytes reserved for non-cache usage per shard e.g. failover.
	MaxmemoryReserved *string `json:"maxmemory-reserved,omitempty"`

	// NotifyKeyspaceEvents: The keyspace events which should be monitored.
	NotifyKeyspaceEvents *string `json:"notify-keyspace-events,omitempty"`

	// PreferredDataPersistenceAuthMethod: Preferred auth method to communicate to storage account used for data persistence,
	// specify SAS or ManagedIdentity, default value is SAS
	PreferredDataPersistenceAuthMethod *string `json:"preferred-data-persistence-auth-method,omitempty"`

	// RdbBackupEnabled: Specifies whether the rdb backup is enabled
	RdbBackupEnabled *string `json:"rdb-backup-enabled,omitempty"`

	// RdbBackupFrequency: Specifies the frequency for creating rdb backup in minutes. Valid values: (15, 30, 60, 360, 720,
	// 1440)
	RdbBackupFrequency *string `json:"rdb-backup-frequency,omitempty"`

	// RdbBackupMaxSnapshotCount: Specifies the maximum number of snapshots for rdb backup
	RdbBackupMaxSnapshotCount *string `json:"rdb-backup-max-snapshot-count,omitempty"`

	// RdbStorageConnectionString: The storage account connection string for storing rdb file
	RdbStorageConnectionString *string `json:"rdb-storage-connection-string,omitempty"`

	// StorageSubscriptionId: SubscriptionId of the storage account for persistence (aof/rdb) using ManagedIdentity.
	StorageSubscriptionId *string `json:"storage-subscription-id,omitempty"`
}

// +kubebuilder:validation:Enum={"Preview","Stable"}
type RedisCreateProperties_UpdateChannel_ARM string

const (
	RedisCreateProperties_UpdateChannel_ARM_Preview = RedisCreateProperties_UpdateChannel_ARM("Preview")
	RedisCreateProperties_UpdateChannel_ARM_Stable  = RedisCreateProperties_UpdateChannel_ARM("Stable")
)

// Mapping from string to RedisCreateProperties_UpdateChannel_ARM
var redisCreateProperties_UpdateChannel_ARM_Values = map[string]RedisCreateProperties_UpdateChannel_ARM{
	"preview": RedisCreateProperties_UpdateChannel_ARM_Preview,
	"stable":  RedisCreateProperties_UpdateChannel_ARM_Stable,
}

// SKU parameters supplied to the create Redis operation.
type Sku_ARM struct {
	// Capacity: The size of the Redis cache to deploy. Valid values: for C (Basic/Standard) family (0, 1, 2, 3, 4, 5, 6), for
	// P (Premium) family (1, 2, 3, 4).
	Capacity *int `json:"capacity,omitempty"`

	// Family: The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).
	Family *Sku_Family_ARM `json:"family,omitempty"`

	// Name: The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium)
	Name *Sku_Name_ARM `json:"name,omitempty"`
}

// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails_ARM struct {
}

// +kubebuilder:validation:Enum={"C","P"}
type Sku_Family_ARM string

const (
	Sku_Family_ARM_C = Sku_Family_ARM("C")
	Sku_Family_ARM_P = Sku_Family_ARM("P")
)

// Mapping from string to Sku_Family_ARM
var sku_Family_ARM_Values = map[string]Sku_Family_ARM{
	"c": Sku_Family_ARM_C,
	"p": Sku_Family_ARM_P,
}

// +kubebuilder:validation:Enum={"Basic","Premium","Standard"}
type Sku_Name_ARM string

const (
	Sku_Name_ARM_Basic    = Sku_Name_ARM("Basic")
	Sku_Name_ARM_Premium  = Sku_Name_ARM("Premium")
	Sku_Name_ARM_Standard = Sku_Name_ARM("Standard")
)

// Mapping from string to Sku_Name_ARM
var sku_Name_ARM_Values = map[string]Sku_Name_ARM{
	"basic":    Sku_Name_ARM_Basic,
	"premium":  Sku_Name_ARM_Premium,
	"standard": Sku_Name_ARM_Standard,
}
