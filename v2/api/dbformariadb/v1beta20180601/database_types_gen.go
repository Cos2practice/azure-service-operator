// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180601

import (
	"fmt"
	v20180601s "github.com/Azure/azure-service-operator/v2/api/dbformariadb/v1beta20180601storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Generated from: https://schema.management.azure.com/schemas/2018-06-01/Microsoft.DBforMariaDB.json#/resourceDefinitions/servers_databases
type Database struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Servers_Database_Spec `json:"spec,omitempty"`
	Status            Database_STATUS       `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Database{}

// GetConditions returns the conditions of the resource
func (database *Database) GetConditions() conditions.Conditions {
	return database.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (database *Database) SetConditions(conditions conditions.Conditions) {
	database.Status.Conditions = conditions
}

var _ conversion.Convertible = &Database{}

// ConvertFrom populates our Database from the provided hub Database
func (database *Database) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20180601s.Database)
	if !ok {
		return fmt.Errorf("expected dbformariadb/v1beta20180601storage/Database but received %T instead", hub)
	}

	return database.AssignProperties_From_Database(source)
}

// ConvertTo populates the provided hub Database from our Database
func (database *Database) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20180601s.Database)
	if !ok {
		return fmt.Errorf("expected dbformariadb/v1beta20180601storage/Database but received %T instead", hub)
	}

	return database.AssignProperties_To_Database(destination)
}

// +kubebuilder:webhook:path=/mutate-dbformariadb-azure-com-v1beta20180601-database,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbformariadb.azure.com,resources=databases,verbs=create;update,versions=v1beta20180601,name=default.v1beta20180601.databases.dbformariadb.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &Database{}

// Default applies defaults to the Database resource
func (database *Database) Default() {
	database.defaultImpl()
	var temp interface{} = database
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (database *Database) defaultAzureName() {
	if database.Spec.AzureName == "" {
		database.Spec.AzureName = database.Name
	}
}

// defaultImpl applies the code generated defaults to the Database resource
func (database *Database) defaultImpl() { database.defaultAzureName() }

var _ genruntime.KubernetesResource = &Database{}

// AzureName returns the Azure name of the resource
func (database *Database) AzureName() string {
	return database.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-06-01"
func (database Database) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (database *Database) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (database *Database) GetSpec() genruntime.ConvertibleSpec {
	return &database.Spec
}

// GetStatus returns the status of this resource
func (database *Database) GetStatus() genruntime.ConvertibleStatus {
	return &database.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMariaDB/servers/databases"
func (database *Database) GetType() string {
	return "Microsoft.DBforMariaDB/servers/databases"
}

// NewEmptyStatus returns a new empty (blank) status
func (database *Database) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Database_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (database *Database) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(database.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  database.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (database *Database) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Database_STATUS); ok {
		database.Status = *st
		return nil
	}

	// Convert status to required version
	var st Database_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	database.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-dbformariadb-azure-com-v1beta20180601-database,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbformariadb.azure.com,resources=databases,verbs=create;update,versions=v1beta20180601,name=validate.v1beta20180601.databases.dbformariadb.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &Database{}

// ValidateCreate validates the creation of the resource
func (database *Database) ValidateCreate() error {
	validations := database.createValidations()
	var temp interface{} = database
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateDelete validates the deletion of the resource
func (database *Database) ValidateDelete() error {
	validations := database.deleteValidations()
	var temp interface{} = database
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateUpdate validates an update of the resource
func (database *Database) ValidateUpdate(old runtime.Object) error {
	validations := database.updateValidations()
	var temp interface{} = database
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation(old)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// createValidations validates the creation of the resource
func (database *Database) createValidations() []func() error {
	return []func() error{database.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (database *Database) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (database *Database) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return database.validateResourceReferences()
		},
		database.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (database *Database) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&database.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (database *Database) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*Database)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, database)
}

// AssignProperties_From_Database populates our Database from the provided source Database
func (database *Database) AssignProperties_From_Database(source *v20180601s.Database) error {

	// ObjectMeta
	database.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Servers_Database_Spec
	err := spec.AssignProperties_From_Servers_Database_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_Database_Spec() to populate field Spec")
	}
	database.Spec = spec

	// Status
	var status Database_STATUS
	err = status.AssignProperties_From_Database_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Database_STATUS() to populate field Status")
	}
	database.Status = status

	// No error
	return nil
}

// AssignProperties_To_Database populates the provided destination Database from our Database
func (database *Database) AssignProperties_To_Database(destination *v20180601s.Database) error {

	// ObjectMeta
	destination.ObjectMeta = *database.ObjectMeta.DeepCopy()

	// Spec
	var spec v20180601s.Servers_Database_Spec
	err := database.Spec.AssignProperties_To_Servers_Database_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_Database_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20180601s.Database_STATUS
	err = database.Status.AssignProperties_To_Database_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Database_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (database *Database) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: database.Spec.OriginalVersion(),
		Kind:    "Database",
	}
}

// +kubebuilder:object:root=true
// Generated from: https://schema.management.azure.com/schemas/2018-06-01/Microsoft.DBforMariaDB.json#/resourceDefinitions/servers_databases
type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Database `json:"items"`
}

type Database_STATUS struct {
	// Charset: The charset of the database.
	Charset *string `json:"charset,omitempty"`

	// Collation: The collation of the database.
	Collation *string `json:"collation,omitempty"`

	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Database_STATUS{}

// ConvertStatusFrom populates our Database_STATUS from the provided source
func (database *Database_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20180601s.Database_STATUS)
	if ok {
		// Populate our instance from source
		return database.AssignProperties_From_Database_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20180601s.Database_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = database.AssignProperties_From_Database_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Database_STATUS
func (database *Database_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20180601s.Database_STATUS)
	if ok {
		// Populate destination from our instance
		return database.AssignProperties_To_Database_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20180601s.Database_STATUS{}
	err := database.AssignProperties_To_Database_STATUS(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

var _ genruntime.FromARMConverter = &Database_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (database *Database_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Database_STATUSARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (database *Database_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Database_STATUSARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Database_STATUSARM, got %T", armInput)
	}

	// Set property ‘Charset’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Charset != nil {
			charset := *typedInput.Properties.Charset
			database.Charset = &charset
		}
	}

	// Set property ‘Collation’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Collation != nil {
			collation := *typedInput.Properties.Collation
			database.Collation = &collation
		}
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		database.Id = &id
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		database.Name = &name
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		database.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_Database_STATUS populates our Database_STATUS from the provided source Database_STATUS
func (database *Database_STATUS) AssignProperties_From_Database_STATUS(source *v20180601s.Database_STATUS) error {

	// Charset
	database.Charset = genruntime.ClonePointerToString(source.Charset)

	// Collation
	database.Collation = genruntime.ClonePointerToString(source.Collation)

	// Conditions
	database.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	database.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	database.Name = genruntime.ClonePointerToString(source.Name)

	// Type
	database.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_Database_STATUS populates the provided destination Database_STATUS from our Database_STATUS
func (database *Database_STATUS) AssignProperties_To_Database_STATUS(destination *v20180601s.Database_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Charset
	destination.Charset = genruntime.ClonePointerToString(database.Charset)

	// Collation
	destination.Collation = genruntime.ClonePointerToString(database.Collation)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(database.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(database.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(database.Name)

	// Type
	destination.Type = genruntime.ClonePointerToString(database.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type Servers_Database_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// Charset: The charset of the database.
	Charset *string `json:"charset,omitempty"`

	// Collation: The collation of the database.
	Collation *string `json:"collation,omitempty"`

	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a dbformariadb.azure.com/Server resource
	Owner *genruntime.KnownResourceReference `group:"dbformariadb.azure.com" json:"owner,omitempty" kind:"Server"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &Servers_Database_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (database *Servers_Database_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if database == nil {
		return nil, nil
	}
	result := &Servers_Database_SpecARM{}

	// Set property ‘Location’:
	if database.Location != nil {
		location := *database.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if database.Charset != nil || database.Collation != nil {
		result.Properties = &DatabasePropertiesARM{}
	}
	if database.Charset != nil {
		charset := *database.Charset
		result.Properties.Charset = &charset
	}
	if database.Collation != nil {
		collation := *database.Collation
		result.Properties.Collation = &collation
	}

	// Set property ‘Tags’:
	if database.Tags != nil {
		result.Tags = make(map[string]string, len(database.Tags))
		for key, value := range database.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (database *Servers_Database_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_Database_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (database *Servers_Database_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_Database_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_Database_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	database.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Charset’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Charset != nil {
			charset := *typedInput.Properties.Charset
			database.Charset = &charset
		}
	}

	// Set property ‘Collation’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Collation != nil {
			collation := *typedInput.Properties.Collation
			database.Collation = &collation
		}
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		database.Location = &location
	}

	// Set property ‘Owner’:
	database.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		database.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			database.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Servers_Database_Spec{}

// ConvertSpecFrom populates our Servers_Database_Spec from the provided source
func (database *Servers_Database_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20180601s.Servers_Database_Spec)
	if ok {
		// Populate our instance from source
		return database.AssignProperties_From_Servers_Database_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20180601s.Servers_Database_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = database.AssignProperties_From_Servers_Database_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Servers_Database_Spec
func (database *Servers_Database_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20180601s.Servers_Database_Spec)
	if ok {
		// Populate destination from our instance
		return database.AssignProperties_To_Servers_Database_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20180601s.Servers_Database_Spec{}
	err := database.AssignProperties_To_Servers_Database_Spec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_Servers_Database_Spec populates our Servers_Database_Spec from the provided source Servers_Database_Spec
func (database *Servers_Database_Spec) AssignProperties_From_Servers_Database_Spec(source *v20180601s.Servers_Database_Spec) error {

	// AzureName
	database.AzureName = source.AzureName

	// Charset
	database.Charset = genruntime.ClonePointerToString(source.Charset)

	// Collation
	database.Collation = genruntime.ClonePointerToString(source.Collation)

	// Location
	database.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		database.Owner = &owner
	} else {
		database.Owner = nil
	}

	// Tags
	database.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignProperties_To_Servers_Database_Spec populates the provided destination Servers_Database_Spec from our Servers_Database_Spec
func (database *Servers_Database_Spec) AssignProperties_To_Servers_Database_Spec(destination *v20180601s.Servers_Database_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = database.AzureName

	// Charset
	destination.Charset = genruntime.ClonePointerToString(database.Charset)

	// Collation
	destination.Collation = genruntime.ClonePointerToString(database.Collation)

	// Location
	destination.Location = genruntime.ClonePointerToString(database.Location)

	// OriginalVersion
	destination.OriginalVersion = database.OriginalVersion()

	// Owner
	if database.Owner != nil {
		owner := database.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(database.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (database *Servers_Database_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (database *Servers_Database_Spec) SetAzureName(azureName string) { database.AzureName = azureName }

func init() {
	SchemeBuilder.Register(&Database{}, &DatabaseList{})
}