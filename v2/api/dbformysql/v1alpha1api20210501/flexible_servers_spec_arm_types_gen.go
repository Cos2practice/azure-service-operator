// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210501

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of FlexibleServers_Spec. Use v1beta20210501.FlexibleServers_Spec instead
type FlexibleServers_SpecARM struct {
	Identity   *IdentityARM         `json:"identity,omitempty"`
	Location   *string              `json:"location,omitempty"`
	Name       string               `json:"name,omitempty"`
	Properties *ServerPropertiesARM `json:"properties,omitempty"`
	Sku        *SkuARM              `json:"sku,omitempty"`
	Tags       map[string]string    `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServers_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-01"
func (servers FlexibleServers_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (servers *FlexibleServers_SpecARM) GetName() string {
	return servers.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMySQL/flexibleServers"
func (servers *FlexibleServers_SpecARM) GetType() string {
	return "Microsoft.DBforMySQL/flexibleServers"
}

// Deprecated version of Identity. Use v1beta20210501.Identity instead
type IdentityARM struct {
	Type *Identity_Type `json:"type,omitempty"`
}

// Deprecated version of ServerProperties. Use v1beta20210501.ServerProperties instead
type ServerPropertiesARM struct {
	AdministratorLogin         *string                           `json:"administratorLogin,omitempty"`
	AdministratorLoginPassword *string                           `json:"administratorLoginPassword,omitempty"`
	AvailabilityZone           *string                           `json:"availabilityZone,omitempty"`
	Backup                     *BackupARM                        `json:"backup,omitempty"`
	CreateMode                 *ServerProperties_CreateMode      `json:"createMode,omitempty"`
	DataEncryption             *DataEncryptionARM                `json:"dataEncryption,omitempty"`
	HighAvailability           *HighAvailabilityARM              `json:"highAvailability,omitempty"`
	MaintenanceWindow          *MaintenanceWindowARM             `json:"maintenanceWindow,omitempty"`
	Network                    *NetworkARM                       `json:"network,omitempty"`
	ReplicationRole            *ServerProperties_ReplicationRole `json:"replicationRole,omitempty"`
	RestorePointInTime         *string                           `json:"restorePointInTime,omitempty"`
	SourceServerResourceId     *string                           `json:"sourceServerResourceId,omitempty"`
	Storage                    *StorageARM                       `json:"storage,omitempty"`
	Version                    *ServerProperties_Version         `json:"version,omitempty"`
}

// Deprecated version of Sku. Use v1beta20210501.Sku instead
type SkuARM struct {
	Name *string   `json:"name,omitempty"`
	Tier *Sku_Tier `json:"tier,omitempty"`
}

// Deprecated version of Backup. Use v1beta20210501.Backup instead
type BackupARM struct {
	BackupRetentionDays *int                       `json:"backupRetentionDays,omitempty"`
	GeoRedundantBackup  *Backup_GeoRedundantBackup `json:"geoRedundantBackup,omitempty"`
}

// Deprecated version of DataEncryption. Use v1beta20210501.DataEncryption instead
type DataEncryptionARM struct {
	GeoBackupKeyURI                 *string              `json:"geoBackupKeyURI,omitempty"`
	GeoBackupUserAssignedIdentityId *string              `json:"geoBackupUserAssignedIdentityId,omitempty"`
	PrimaryKeyURI                   *string              `json:"primaryKeyURI,omitempty"`
	PrimaryUserAssignedIdentityId   *string              `json:"primaryUserAssignedIdentityId,omitempty"`
	Type                            *DataEncryption_Type `json:"type,omitempty"`
}

// Deprecated version of HighAvailability. Use v1beta20210501.HighAvailability instead
type HighAvailabilityARM struct {
	Mode                    *HighAvailability_Mode `json:"mode,omitempty"`
	StandbyAvailabilityZone *string                `json:"standbyAvailabilityZone,omitempty"`
}

// Deprecated version of Identity_Type. Use v1beta20210501.Identity_Type instead
// +kubebuilder:validation:Enum={"UserAssigned"}
type Identity_Type string

const Identity_Type_UserAssigned = Identity_Type("UserAssigned")

// Deprecated version of MaintenanceWindow. Use v1beta20210501.MaintenanceWindow instead
type MaintenanceWindowARM struct {
	CustomWindow *string `json:"customWindow,omitempty"`
	DayOfWeek    *int    `json:"dayOfWeek,omitempty"`
	StartHour    *int    `json:"startHour,omitempty"`
	StartMinute  *int    `json:"startMinute,omitempty"`
}

// Deprecated version of Network. Use v1beta20210501.Network instead
type NetworkARM struct {
	DelegatedSubnetResourceId *string `json:"delegatedSubnetResourceId,omitempty"`
	PrivateDnsZoneResourceId  *string `json:"privateDnsZoneResourceId,omitempty"`
}

// Deprecated version of Sku_Tier. Use v1beta20210501.Sku_Tier instead
// +kubebuilder:validation:Enum={"Burstable","GeneralPurpose","MemoryOptimized"}
type Sku_Tier string

const (
	Sku_Tier_Burstable       = Sku_Tier("Burstable")
	Sku_Tier_GeneralPurpose  = Sku_Tier("GeneralPurpose")
	Sku_Tier_MemoryOptimized = Sku_Tier("MemoryOptimized")
)

// Deprecated version of Storage. Use v1beta20210501.Storage instead
type StorageARM struct {
	AutoGrow      *Storage_AutoGrow `json:"autoGrow,omitempty"`
	Iops          *int              `json:"iops,omitempty"`
	StorageSizeGB *int              `json:"storageSizeGB,omitempty"`
}