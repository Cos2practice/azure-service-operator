// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARM struct {
	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties to create and update Azure Cosmos DB trigger.
	Properties *SqlTriggerCreateUpdateProperties_ARM `json:"properties,omitempty"`
	Tags       map[string]string                     `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (trigger DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARM) GetAPIVersion() string {
	return "2021-05-15"
}

// GetName returns the Name of the resource
func (trigger *DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARM) GetName() string {
	return trigger.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/triggers"
func (trigger *DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARM) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/triggers"
}

// Properties to create and update Azure Cosmos DB trigger.
type SqlTriggerCreateUpdateProperties_ARM struct {
	// Options: A key-value pair of options to be applied for the request. This corresponds to the headers sent with the
	// request.
	Options *CreateUpdateOptions_ARM `json:"options,omitempty"`

	// Resource: The standard JSON format of a trigger
	Resource *SqlTriggerResource_ARM `json:"resource,omitempty"`
}

// Cosmos DB SQL trigger resource object
type SqlTriggerResource_ARM struct {
	// Body: Body of the Trigger
	Body *string `json:"body,omitempty"`

	// Id: Name of the Cosmos DB SQL trigger
	Id *string `json:"id,omitempty"`

	// TriggerOperation: The operation the trigger is associated with
	TriggerOperation *SqlTriggerResource_TriggerOperation_ARM `json:"triggerOperation,omitempty"`

	// TriggerType: Type of the Trigger
	TriggerType *SqlTriggerResource_TriggerType_ARM `json:"triggerType,omitempty"`
}

// +kubebuilder:validation:Enum={"All","Create","Delete","Replace","Update"}
type SqlTriggerResource_TriggerOperation_ARM string

const (
	SqlTriggerResource_TriggerOperation_ARM_All     = SqlTriggerResource_TriggerOperation_ARM("All")
	SqlTriggerResource_TriggerOperation_ARM_Create  = SqlTriggerResource_TriggerOperation_ARM("Create")
	SqlTriggerResource_TriggerOperation_ARM_Delete  = SqlTriggerResource_TriggerOperation_ARM("Delete")
	SqlTriggerResource_TriggerOperation_ARM_Replace = SqlTriggerResource_TriggerOperation_ARM("Replace")
	SqlTriggerResource_TriggerOperation_ARM_Update  = SqlTriggerResource_TriggerOperation_ARM("Update")
)

// Mapping from string to SqlTriggerResource_TriggerOperation_ARM
var sqlTriggerResource_TriggerOperation_ARM_Values = map[string]SqlTriggerResource_TriggerOperation_ARM{
	"all":     SqlTriggerResource_TriggerOperation_ARM_All,
	"create":  SqlTriggerResource_TriggerOperation_ARM_Create,
	"delete":  SqlTriggerResource_TriggerOperation_ARM_Delete,
	"replace": SqlTriggerResource_TriggerOperation_ARM_Replace,
	"update":  SqlTriggerResource_TriggerOperation_ARM_Update,
}

// +kubebuilder:validation:Enum={"Post","Pre"}
type SqlTriggerResource_TriggerType_ARM string

const (
	SqlTriggerResource_TriggerType_ARM_Post = SqlTriggerResource_TriggerType_ARM("Post")
	SqlTriggerResource_TriggerType_ARM_Pre  = SqlTriggerResource_TriggerType_ARM("Pre")
)

// Mapping from string to SqlTriggerResource_TriggerType_ARM
var sqlTriggerResource_TriggerType_ARM_Values = map[string]SqlTriggerResource_TriggerType_ARM{
	"post": SqlTriggerResource_TriggerType_ARM_Post,
	"pre":  SqlTriggerResource_TriggerType_ARM_Pre,
}
