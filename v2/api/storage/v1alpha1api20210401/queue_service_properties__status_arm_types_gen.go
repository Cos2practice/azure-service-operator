// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

// Deprecated version of QueueServiceProperties_Status. Use v1beta20210401.QueueServiceProperties_Status instead
type QueueServiceProperties_StatusARM struct {
	Id         *string                                      `json:"id,omitempty"`
	Name       *string                                      `json:"name,omitempty"`
	Properties *QueueServiceProperties_Status_PropertiesARM `json:"properties,omitempty"`
	Type       *string                                      `json:"type,omitempty"`
}

// Deprecated version of QueueServiceProperties_Status_Properties. Use v1beta20210401.QueueServiceProperties_Status_Properties instead
type QueueServiceProperties_Status_PropertiesARM struct {
	Cors *CorsRules_StatusARM `json:"cors,omitempty"`
}