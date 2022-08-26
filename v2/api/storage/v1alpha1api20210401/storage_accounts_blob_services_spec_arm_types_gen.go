// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of StorageAccounts_BlobServices_Spec. Use v1beta20210401.StorageAccounts_BlobServices_Spec instead
type StorageAccounts_BlobServices_SpecARM struct {
	Location   *string                             `json:"location,omitempty"`
	Name       string                              `json:"name,omitempty"`
	Properties *BlobServicePropertiesPropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string                   `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &StorageAccounts_BlobServices_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (services StorageAccounts_BlobServices_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (services *StorageAccounts_BlobServices_SpecARM) GetName() string {
	return services.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/blobServices"
func (services *StorageAccounts_BlobServices_SpecARM) GetType() string {
	return "Microsoft.Storage/storageAccounts/blobServices"
}

// Deprecated version of BlobServicePropertiesProperties. Use v1beta20210401.BlobServicePropertiesProperties instead
type BlobServicePropertiesPropertiesARM struct {
	AutomaticSnapshotPolicyEnabled *bool                            `json:"automaticSnapshotPolicyEnabled,omitempty"`
	ChangeFeed                     *ChangeFeedARM                   `json:"changeFeed,omitempty"`
	ContainerDeleteRetentionPolicy *DeleteRetentionPolicyARM        `json:"containerDeleteRetentionPolicy,omitempty"`
	Cors                           *CorsRulesARM                    `json:"cors,omitempty"`
	DefaultServiceVersion          *string                          `json:"defaultServiceVersion,omitempty"`
	DeleteRetentionPolicy          *DeleteRetentionPolicyARM        `json:"deleteRetentionPolicy,omitempty"`
	IsVersioningEnabled            *bool                            `json:"isVersioningEnabled,omitempty"`
	LastAccessTimeTrackingPolicy   *LastAccessTimeTrackingPolicyARM `json:"lastAccessTimeTrackingPolicy,omitempty"`
	RestorePolicy                  *RestorePolicyPropertiesARM      `json:"restorePolicy,omitempty"`
}

// Deprecated version of ChangeFeed. Use v1beta20210401.ChangeFeed instead
type ChangeFeedARM struct {
	Enabled         *bool `json:"enabled,omitempty"`
	RetentionInDays *int  `json:"retentionInDays,omitempty"`
}

// Deprecated version of CorsRules. Use v1beta20210401.CorsRules instead
type CorsRulesARM struct {
	CorsRules []CorsRuleARM `json:"corsRules,omitempty"`
}

// Deprecated version of DeleteRetentionPolicy. Use v1beta20210401.DeleteRetentionPolicy instead
type DeleteRetentionPolicyARM struct {
	Days    *int  `json:"days,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// Deprecated version of LastAccessTimeTrackingPolicy. Use v1beta20210401.LastAccessTimeTrackingPolicy instead
type LastAccessTimeTrackingPolicyARM struct {
	BlobType                  []string                           `json:"blobType,omitempty"`
	Enable                    *bool                              `json:"enable,omitempty"`
	Name                      *LastAccessTimeTrackingPolicy_Name `json:"name,omitempty"`
	TrackingGranularityInDays *int                               `json:"trackingGranularityInDays,omitempty"`
}

// Deprecated version of RestorePolicyProperties. Use v1beta20210401.RestorePolicyProperties instead
type RestorePolicyPropertiesARM struct {
	Days    *int  `json:"days,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// Deprecated version of CorsRule. Use v1beta20210401.CorsRule instead
type CorsRuleARM struct {
	AllowedHeaders  []string                  `json:"allowedHeaders,omitempty"`
	AllowedMethods  []CorsRule_AllowedMethods `json:"allowedMethods,omitempty"`
	AllowedOrigins  []string                  `json:"allowedOrigins,omitempty"`
	ExposedHeaders  []string                  `json:"exposedHeaders,omitempty"`
	MaxAgeInSeconds *int                      `json:"maxAgeInSeconds,omitempty"`
}