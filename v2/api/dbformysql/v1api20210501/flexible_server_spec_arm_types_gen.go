// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210501

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type FlexibleServer_Spec_ARM struct {
	// Identity: The cmk identity for the server.
	Identity *Identity_ARM `json:"identity,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties of the server.
	Properties *ServerProperties_ARM `json:"properties,omitempty"`

	// Sku: The SKU (pricing tier) of the server.
	Sku *Sku_ARM `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServer_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-01"
func (server FlexibleServer_Spec_ARM) GetAPIVersion() string {
	return "2021-05-01"
}

// GetName returns the Name of the resource
func (server *FlexibleServer_Spec_ARM) GetName() string {
	return server.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMySQL/flexibleServers"
func (server *FlexibleServer_Spec_ARM) GetType() string {
	return "Microsoft.DBforMySQL/flexibleServers"
}

// Properties to configure Identity for Bring your Own Keys
type Identity_ARM struct {
	// Type: Type of managed service identity.
	Type                   *Identity_Type_ARM                         `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails_ARM `json:"userAssignedIdentities,omitempty"`
}

// The properties of a server.
type ServerProperties_ARM struct {
	// AdministratorLogin: The administrator's login name of a server. Can only be specified when the server is being created
	// (and is required for creation).
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	// AdministratorLoginPassword: The password of the administrator login (required for server creation).
	AdministratorLoginPassword *string `json:"administratorLoginPassword,omitempty"`

	// AvailabilityZone: availability Zone information of the server.
	AvailabilityZone *string `json:"availabilityZone,omitempty"`

	// Backup: Backup related properties of a server.
	Backup *Backup_ARM `json:"backup,omitempty"`

	// CreateMode: The mode to create a new MySQL server.
	CreateMode *ServerProperties_CreateMode_ARM `json:"createMode,omitempty"`

	// DataEncryption: The Data Encryption for CMK.
	DataEncryption *DataEncryption_ARM `json:"dataEncryption,omitempty"`

	// HighAvailability: High availability related properties of a server.
	HighAvailability *HighAvailability_ARM `json:"highAvailability,omitempty"`

	// MaintenanceWindow: Maintenance window of a server.
	MaintenanceWindow *MaintenanceWindow_ARM `json:"maintenanceWindow,omitempty"`

	// Network: Network related properties of a server.
	Network *Network_ARM `json:"network,omitempty"`

	// ReplicationRole: The replication role.
	ReplicationRole *ReplicationRole_ARM `json:"replicationRole,omitempty"`

	// RestorePointInTime: Restore point creation time (ISO8601 format), specifying the time to restore from.
	RestorePointInTime *string `json:"restorePointInTime,omitempty"`

	// SourceServerResourceId: The source MySQL server id.
	SourceServerResourceId *string `json:"sourceServerResourceId,omitempty"`

	// Storage: Storage related properties of a server.
	Storage *Storage_ARM `json:"storage,omitempty"`

	// Version: Server version.
	Version *ServerVersion_ARM `json:"version,omitempty"`
}

// Billing information related properties of a server.
type Sku_ARM struct {
	// Name: The name of the sku, e.g. Standard_D32s_v3.
	Name *string `json:"name,omitempty"`

	// Tier: The tier of the particular SKU, e.g. GeneralPurpose.
	Tier *Sku_Tier_ARM `json:"tier,omitempty"`
}

// Storage Profile properties of a server
type Backup_ARM struct {
	// BackupRetentionDays: Backup retention days for the server.
	BackupRetentionDays *int `json:"backupRetentionDays,omitempty"`

	// GeoRedundantBackup: Whether or not geo redundant backup is enabled.
	GeoRedundantBackup *EnableStatusEnum_ARM `json:"geoRedundantBackup,omitempty"`
}

// The date encryption for cmk.
type DataEncryption_ARM struct {
	// GeoBackupKeyURI: Geo backup key uri as key vault can't cross region, need cmk in same region as geo backup
	GeoBackupKeyURI                 *string `json:"geoBackupKeyURI,omitempty"`
	GeoBackupUserAssignedIdentityId *string `json:"geoBackupUserAssignedIdentityId,omitempty"`

	// PrimaryKeyURI: Primary key uri
	PrimaryKeyURI                 *string `json:"primaryKeyURI,omitempty"`
	PrimaryUserAssignedIdentityId *string `json:"primaryUserAssignedIdentityId,omitempty"`

	// Type: The key type, AzureKeyVault for enable cmk, SystemManaged for disable cmk.
	Type *DataEncryption_Type_ARM `json:"type,omitempty"`
}

// Network related properties of a server
type HighAvailability_ARM struct {
	// Mode: High availability mode for a server.
	Mode *HighAvailability_Mode_ARM `json:"mode,omitempty"`

	// StandbyAvailabilityZone: Availability zone of the standby server.
	StandbyAvailabilityZone *string `json:"standbyAvailabilityZone,omitempty"`
}

// +kubebuilder:validation:Enum={"UserAssigned"}
type Identity_Type_ARM string

const Identity_Type_ARM_UserAssigned = Identity_Type_ARM("UserAssigned")

// Mapping from string to Identity_Type_ARM
var identity_Type_ARM_Values = map[string]Identity_Type_ARM{
	"userassigned": Identity_Type_ARM_UserAssigned,
}

// Maintenance window of a server.
type MaintenanceWindow_ARM struct {
	// CustomWindow: indicates whether custom window is enabled or disabled
	CustomWindow *string `json:"customWindow,omitempty"`

	// DayOfWeek: day of week for maintenance window
	DayOfWeek *int `json:"dayOfWeek,omitempty"`

	// StartHour: start hour for maintenance window
	StartHour *int `json:"startHour,omitempty"`

	// StartMinute: start minute for maintenance window
	StartMinute *int `json:"startMinute,omitempty"`
}

// Network related properties of a server
type Network_ARM struct {
	DelegatedSubnetResourceId *string `json:"delegatedSubnetResourceId,omitempty"`
	PrivateDnsZoneResourceId  *string `json:"privateDnsZoneResourceId,omitempty"`
}

// The replication role.
// +kubebuilder:validation:Enum={"None","Replica","Source"}
type ReplicationRole_ARM string

const (
	ReplicationRole_ARM_None    = ReplicationRole_ARM("None")
	ReplicationRole_ARM_Replica = ReplicationRole_ARM("Replica")
	ReplicationRole_ARM_Source  = ReplicationRole_ARM("Source")
)

// Mapping from string to ReplicationRole_ARM
var replicationRole_ARM_Values = map[string]ReplicationRole_ARM{
	"none":    ReplicationRole_ARM_None,
	"replica": ReplicationRole_ARM_Replica,
	"source":  ReplicationRole_ARM_Source,
}

// +kubebuilder:validation:Enum={"Default","GeoRestore","PointInTimeRestore","Replica"}
type ServerProperties_CreateMode_ARM string

const (
	ServerProperties_CreateMode_ARM_Default            = ServerProperties_CreateMode_ARM("Default")
	ServerProperties_CreateMode_ARM_GeoRestore         = ServerProperties_CreateMode_ARM("GeoRestore")
	ServerProperties_CreateMode_ARM_PointInTimeRestore = ServerProperties_CreateMode_ARM("PointInTimeRestore")
	ServerProperties_CreateMode_ARM_Replica            = ServerProperties_CreateMode_ARM("Replica")
)

// Mapping from string to ServerProperties_CreateMode_ARM
var serverProperties_CreateMode_ARM_Values = map[string]ServerProperties_CreateMode_ARM{
	"default":            ServerProperties_CreateMode_ARM_Default,
	"georestore":         ServerProperties_CreateMode_ARM_GeoRestore,
	"pointintimerestore": ServerProperties_CreateMode_ARM_PointInTimeRestore,
	"replica":            ServerProperties_CreateMode_ARM_Replica,
}

// The version of a server.
// +kubebuilder:validation:Enum={"5.7","8.0.21"}
type ServerVersion_ARM string

const (
	ServerVersion_ARM_57   = ServerVersion_ARM("5.7")
	ServerVersion_ARM_8021 = ServerVersion_ARM("8.0.21")
)

// Mapping from string to ServerVersion_ARM
var serverVersion_ARM_Values = map[string]ServerVersion_ARM{
	"5.7":    ServerVersion_ARM_57,
	"8.0.21": ServerVersion_ARM_8021,
}

// +kubebuilder:validation:Enum={"Burstable","GeneralPurpose","MemoryOptimized"}
type Sku_Tier_ARM string

const (
	Sku_Tier_ARM_Burstable       = Sku_Tier_ARM("Burstable")
	Sku_Tier_ARM_GeneralPurpose  = Sku_Tier_ARM("GeneralPurpose")
	Sku_Tier_ARM_MemoryOptimized = Sku_Tier_ARM("MemoryOptimized")
)

// Mapping from string to Sku_Tier_ARM
var sku_Tier_ARM_Values = map[string]Sku_Tier_ARM{
	"burstable":       Sku_Tier_ARM_Burstable,
	"generalpurpose":  Sku_Tier_ARM_GeneralPurpose,
	"memoryoptimized": Sku_Tier_ARM_MemoryOptimized,
}

// Storage Profile properties of a server
type Storage_ARM struct {
	// AutoGrow: Enable Storage Auto Grow or not.
	AutoGrow *EnableStatusEnum_ARM `json:"autoGrow,omitempty"`

	// Iops: Storage IOPS for a server.
	Iops *int `json:"iops,omitempty"`

	// StorageSizeGB: Max storage size allowed for a server.
	StorageSizeGB *int `json:"storageSizeGB,omitempty"`
}

// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails_ARM struct {
}

// +kubebuilder:validation:Enum={"AzureKeyVault","SystemManaged"}
type DataEncryption_Type_ARM string

const (
	DataEncryption_Type_ARM_AzureKeyVault = DataEncryption_Type_ARM("AzureKeyVault")
	DataEncryption_Type_ARM_SystemManaged = DataEncryption_Type_ARM("SystemManaged")
)

// Mapping from string to DataEncryption_Type_ARM
var dataEncryption_Type_ARM_Values = map[string]DataEncryption_Type_ARM{
	"azurekeyvault": DataEncryption_Type_ARM_AzureKeyVault,
	"systemmanaged": DataEncryption_Type_ARM_SystemManaged,
}

// Enum to indicate whether value is 'Enabled' or 'Disabled'
// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type EnableStatusEnum_ARM string

const (
	EnableStatusEnum_ARM_Disabled = EnableStatusEnum_ARM("Disabled")
	EnableStatusEnum_ARM_Enabled  = EnableStatusEnum_ARM("Enabled")
)

// Mapping from string to EnableStatusEnum_ARM
var enableStatusEnum_ARM_Values = map[string]EnableStatusEnum_ARM{
	"disabled": EnableStatusEnum_ARM_Disabled,
	"enabled":  EnableStatusEnum_ARM_Enabled,
}

// +kubebuilder:validation:Enum={"Disabled","SameZone","ZoneRedundant"}
type HighAvailability_Mode_ARM string

const (
	HighAvailability_Mode_ARM_Disabled      = HighAvailability_Mode_ARM("Disabled")
	HighAvailability_Mode_ARM_SameZone      = HighAvailability_Mode_ARM("SameZone")
	HighAvailability_Mode_ARM_ZoneRedundant = HighAvailability_Mode_ARM("ZoneRedundant")
)

// Mapping from string to HighAvailability_Mode_ARM
var highAvailability_Mode_ARM_Values = map[string]HighAvailability_Mode_ARM{
	"disabled":      HighAvailability_Mode_ARM_Disabled,
	"samezone":      HighAvailability_Mode_ARM_SameZone,
	"zoneredundant": HighAvailability_Mode_ARM_ZoneRedundant,
}
