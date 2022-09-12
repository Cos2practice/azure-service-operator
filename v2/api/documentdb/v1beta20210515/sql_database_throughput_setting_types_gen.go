// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515

import (
	"fmt"
	v20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1beta20210515storage"
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
// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_throughputSettings
type SqlDatabaseThroughputSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec `json:"spec,omitempty"`
	Status            ThroughputSettingsGetResults_STATUS                  `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseThroughputSetting{}

// GetConditions returns the conditions of the resource
func (setting *SqlDatabaseThroughputSetting) GetConditions() conditions.Conditions {
	return setting.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (setting *SqlDatabaseThroughputSetting) SetConditions(conditions conditions.Conditions) {
	setting.Status.Conditions = conditions
}

var _ conversion.Convertible = &SqlDatabaseThroughputSetting{}

// ConvertFrom populates our SqlDatabaseThroughputSetting from the provided hub SqlDatabaseThroughputSetting
func (setting *SqlDatabaseThroughputSetting) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20210515s.SqlDatabaseThroughputSetting)
	if !ok {
		return fmt.Errorf("expected documentdb/v1beta20210515storage/SqlDatabaseThroughputSetting but received %T instead", hub)
	}

	return setting.AssignProperties_From_SqlDatabaseThroughputSetting(source)
}

// ConvertTo populates the provided hub SqlDatabaseThroughputSetting from our SqlDatabaseThroughputSetting
func (setting *SqlDatabaseThroughputSetting) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210515s.SqlDatabaseThroughputSetting)
	if !ok {
		return fmt.Errorf("expected documentdb/v1beta20210515storage/SqlDatabaseThroughputSetting but received %T instead", hub)
	}

	return setting.AssignProperties_To_SqlDatabaseThroughputSetting(destination)
}

// +kubebuilder:webhook:path=/mutate-documentdb-azure-com-v1beta20210515-sqldatabasethroughputsetting,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqldatabasethroughputsettings,verbs=create;update,versions=v1beta20210515,name=default.v1beta20210515.sqldatabasethroughputsettings.documentdb.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &SqlDatabaseThroughputSetting{}

// Default applies defaults to the SqlDatabaseThroughputSetting resource
func (setting *SqlDatabaseThroughputSetting) Default() {
	setting.defaultImpl()
	var temp interface{} = setting
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the SqlDatabaseThroughputSetting resource
func (setting *SqlDatabaseThroughputSetting) defaultImpl() {}

var _ genruntime.KubernetesResource = &SqlDatabaseThroughputSetting{}

// AzureName returns the Azure name of the resource (always "default")
func (setting *SqlDatabaseThroughputSetting) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (setting SqlDatabaseThroughputSetting) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (setting *SqlDatabaseThroughputSetting) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (setting *SqlDatabaseThroughputSetting) GetSpec() genruntime.ConvertibleSpec {
	return &setting.Spec
}

// GetStatus returns the status of this resource
func (setting *SqlDatabaseThroughputSetting) GetStatus() genruntime.ConvertibleStatus {
	return &setting.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/throughputSettings"
func (setting *SqlDatabaseThroughputSetting) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/throughputSettings"
}

// NewEmptyStatus returns a new empty (blank) status
func (setting *SqlDatabaseThroughputSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ThroughputSettingsGetResults_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (setting *SqlDatabaseThroughputSetting) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(setting.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  setting.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (setting *SqlDatabaseThroughputSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ThroughputSettingsGetResults_STATUS); ok {
		setting.Status = *st
		return nil
	}

	// Convert status to required version
	var st ThroughputSettingsGetResults_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	setting.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-documentdb-azure-com-v1beta20210515-sqldatabasethroughputsetting,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqldatabasethroughputsettings,verbs=create;update,versions=v1beta20210515,name=validate.v1beta20210515.sqldatabasethroughputsettings.documentdb.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &SqlDatabaseThroughputSetting{}

// ValidateCreate validates the creation of the resource
func (setting *SqlDatabaseThroughputSetting) ValidateCreate() error {
	validations := setting.createValidations()
	var temp interface{} = setting
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
func (setting *SqlDatabaseThroughputSetting) ValidateDelete() error {
	validations := setting.deleteValidations()
	var temp interface{} = setting
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
func (setting *SqlDatabaseThroughputSetting) ValidateUpdate(old runtime.Object) error {
	validations := setting.updateValidations()
	var temp interface{} = setting
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
func (setting *SqlDatabaseThroughputSetting) createValidations() []func() error {
	return []func() error{setting.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (setting *SqlDatabaseThroughputSetting) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (setting *SqlDatabaseThroughputSetting) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return setting.validateResourceReferences()
		},
		setting.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (setting *SqlDatabaseThroughputSetting) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&setting.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (setting *SqlDatabaseThroughputSetting) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*SqlDatabaseThroughputSetting)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, setting)
}

// AssignProperties_From_SqlDatabaseThroughputSetting populates our SqlDatabaseThroughputSetting from the provided source SqlDatabaseThroughputSetting
func (setting *SqlDatabaseThroughputSetting) AssignProperties_From_SqlDatabaseThroughputSetting(source *v20210515s.SqlDatabaseThroughputSetting) error {

	// ObjectMeta
	setting.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec
	err := spec.AssignProperties_From_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec() to populate field Spec")
	}
	setting.Spec = spec

	// Status
	var status ThroughputSettingsGetResults_STATUS
	err = status.AssignProperties_From_ThroughputSettingsGetResults_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_ThroughputSettingsGetResults_STATUS() to populate field Status")
	}
	setting.Status = status

	// No error
	return nil
}

// AssignProperties_To_SqlDatabaseThroughputSetting populates the provided destination SqlDatabaseThroughputSetting from our SqlDatabaseThroughputSetting
func (setting *SqlDatabaseThroughputSetting) AssignProperties_To_SqlDatabaseThroughputSetting(destination *v20210515s.SqlDatabaseThroughputSetting) error {

	// ObjectMeta
	destination.ObjectMeta = *setting.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210515s.DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec
	err := setting.Spec.AssignProperties_To_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210515s.ThroughputSettingsGetResults_STATUS
	err = setting.Status.AssignProperties_To_ThroughputSettingsGetResults_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_ThroughputSettingsGetResults_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (setting *SqlDatabaseThroughputSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: setting.Spec.OriginalVersion(),
		Kind:    "SqlDatabaseThroughputSetting",
	}
}

// +kubebuilder:object:root=true
// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_throughputSettings
type SqlDatabaseThroughputSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseThroughputSetting `json:"items"`
}

type DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec struct {
	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/SqlDatabase resource
	Owner *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"SqlDatabase"`

	// +kubebuilder:validation:Required
	// Resource: Cosmos DB resource throughput object. Either throughput is required or autoscaleSettings is required, but not
	// both.
	Resource *ThroughputSettingsResource `json:"resource,omitempty"`

	// Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
	// resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
	// greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
	// type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph",
	// "DocumentDB", and "MongoDB".
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (setting *DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if setting == nil {
		return nil, nil
	}
	result := &DatabaseAccounts_SqlDatabases_ThroughputSetting_SpecARM{}

	// Set property ‘Location’:
	if setting.Location != nil {
		location := *setting.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if setting.Resource != nil {
		result.Properties = &ThroughputSettingsUpdatePropertiesARM{}
	}
	if setting.Resource != nil {
		resourceARM, err := (*setting.Resource).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		resource := *resourceARM.(*ThroughputSettingsResourceARM)
		result.Properties.Resource = &resource
	}

	// Set property ‘Tags’:
	if setting.Tags != nil {
		result.Tags = make(map[string]string, len(setting.Tags))
		for key, value := range setting.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (setting *DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DatabaseAccounts_SqlDatabases_ThroughputSetting_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (setting *DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DatabaseAccounts_SqlDatabases_ThroughputSetting_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DatabaseAccounts_SqlDatabases_ThroughputSetting_SpecARM, got %T", armInput)
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		setting.Location = &location
	}

	// Set property ‘Owner’:
	setting.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Resource’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Resource != nil {
			var resource1 ThroughputSettingsResource
			err := resource1.PopulateFromARM(owner, *typedInput.Properties.Resource)
			if err != nil {
				return err
			}
			resource := resource1
			setting.Resource = &resource
		}
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		setting.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			setting.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec{}

// ConvertSpecFrom populates our DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec from the provided source
func (setting *DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210515s.DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec)
	if ok {
		// Populate our instance from source
		return setting.AssignProperties_From_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = setting.AssignProperties_From_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec
func (setting *DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210515s.DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec)
	if ok {
		// Populate destination from our instance
		return setting.AssignProperties_To_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec{}
	err := setting.AssignProperties_To_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(dst)
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

// AssignProperties_From_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec populates our DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec from the provided source DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec
func (setting *DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec) AssignProperties_From_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(source *v20210515s.DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec) error {

	// Location
	setting.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		setting.Owner = &owner
	} else {
		setting.Owner = nil
	}

	// Resource
	if source.Resource != nil {
		var resource ThroughputSettingsResource
		err := resource.AssignProperties_From_ThroughputSettingsResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_ThroughputSettingsResource() to populate field Resource")
		}
		setting.Resource = &resource
	} else {
		setting.Resource = nil
	}

	// Tags
	setting.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignProperties_To_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec populates the provided destination DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec from our DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec
func (setting *DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec) AssignProperties_To_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(destination *v20210515s.DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Location
	destination.Location = genruntime.ClonePointerToString(setting.Location)

	// OriginalVersion
	destination.OriginalVersion = setting.OriginalVersion()

	// Owner
	if setting.Owner != nil {
		owner := setting.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Resource
	if setting.Resource != nil {
		var resource v20210515s.ThroughputSettingsResource
		err := setting.Resource.AssignProperties_To_ThroughputSettingsResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_ThroughputSettingsResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(setting.Tags)

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
func (setting *DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

func init() {
	SchemeBuilder.Register(&SqlDatabaseThroughputSetting{}, &SqlDatabaseThroughputSettingList{})
}