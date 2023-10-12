// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20230101 "github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20230101"
	v20230101s "github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20230101storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type BackupVaultExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *BackupVaultExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20230101.BackupVault{},
		&v20230101s.BackupVault{}}
}