// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_AzureFirstPartyManagedCertificateParameters_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AzureFirstPartyManagedCertificateParameters_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAzureFirstPartyManagedCertificateParameters_STATUS_ARM, AzureFirstPartyManagedCertificateParameters_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAzureFirstPartyManagedCertificateParameters_STATUS_ARM runs a test to see if a specific instance of AzureFirstPartyManagedCertificateParameters_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAzureFirstPartyManagedCertificateParameters_STATUS_ARM(subject AzureFirstPartyManagedCertificateParameters_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AzureFirstPartyManagedCertificateParameters_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of AzureFirstPartyManagedCertificateParameters_STATUS_ARM instances for property testing - lazily
// instantiated by AzureFirstPartyManagedCertificateParameters_STATUS_ARMGenerator()
var azureFirstPartyManagedCertificateParameters_STATUS_ARMGenerator gopter.Gen

// AzureFirstPartyManagedCertificateParameters_STATUS_ARMGenerator returns a generator of AzureFirstPartyManagedCertificateParameters_STATUS_ARM instances for property testing.
// We first initialize azureFirstPartyManagedCertificateParameters_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AzureFirstPartyManagedCertificateParameters_STATUS_ARMGenerator() gopter.Gen {
	if azureFirstPartyManagedCertificateParameters_STATUS_ARMGenerator != nil {
		return azureFirstPartyManagedCertificateParameters_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAzureFirstPartyManagedCertificateParameters_STATUS_ARM(generators)
	azureFirstPartyManagedCertificateParameters_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AzureFirstPartyManagedCertificateParameters_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAzureFirstPartyManagedCertificateParameters_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForAzureFirstPartyManagedCertificateParameters_STATUS_ARM(generators)
	azureFirstPartyManagedCertificateParameters_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AzureFirstPartyManagedCertificateParameters_STATUS_ARM{}), generators)

	return azureFirstPartyManagedCertificateParameters_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAzureFirstPartyManagedCertificateParameters_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAzureFirstPartyManagedCertificateParameters_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CertificateAuthority"] = gen.PtrOf(gen.AlphaString())
	gens["ExpirationDate"] = gen.PtrOf(gen.AlphaString())
	gens["Subject"] = gen.PtrOf(gen.AlphaString())
	gens["SubjectAlternativeNames"] = gen.SliceOf(gen.AlphaString())
	gens["Thumbprint"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.OneConstOf(AzureFirstPartyManagedCertificateParameters_Type_STATUS_AzureFirstPartyManagedCertificate)
}

// AddRelatedPropertyGeneratorsForAzureFirstPartyManagedCertificateParameters_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAzureFirstPartyManagedCertificateParameters_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["SecretSource"] = gen.PtrOf(ResourceReference_STATUS_ARMGenerator())
}

func Test_CustomerCertificateParameters_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CustomerCertificateParameters_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCustomerCertificateParameters_STATUS_ARM, CustomerCertificateParameters_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCustomerCertificateParameters_STATUS_ARM runs a test to see if a specific instance of CustomerCertificateParameters_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCustomerCertificateParameters_STATUS_ARM(subject CustomerCertificateParameters_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CustomerCertificateParameters_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of CustomerCertificateParameters_STATUS_ARM instances for property testing - lazily instantiated by
// CustomerCertificateParameters_STATUS_ARMGenerator()
var customerCertificateParameters_STATUS_ARMGenerator gopter.Gen

// CustomerCertificateParameters_STATUS_ARMGenerator returns a generator of CustomerCertificateParameters_STATUS_ARM instances for property testing.
// We first initialize customerCertificateParameters_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func CustomerCertificateParameters_STATUS_ARMGenerator() gopter.Gen {
	if customerCertificateParameters_STATUS_ARMGenerator != nil {
		return customerCertificateParameters_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCustomerCertificateParameters_STATUS_ARM(generators)
	customerCertificateParameters_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(CustomerCertificateParameters_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCustomerCertificateParameters_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForCustomerCertificateParameters_STATUS_ARM(generators)
	customerCertificateParameters_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(CustomerCertificateParameters_STATUS_ARM{}), generators)

	return customerCertificateParameters_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForCustomerCertificateParameters_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCustomerCertificateParameters_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CertificateAuthority"] = gen.PtrOf(gen.AlphaString())
	gens["ExpirationDate"] = gen.PtrOf(gen.AlphaString())
	gens["SecretVersion"] = gen.PtrOf(gen.AlphaString())
	gens["Subject"] = gen.PtrOf(gen.AlphaString())
	gens["SubjectAlternativeNames"] = gen.SliceOf(gen.AlphaString())
	gens["Thumbprint"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.OneConstOf(CustomerCertificateParameters_Type_STATUS_CustomerCertificate)
	gens["UseLatestVersion"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForCustomerCertificateParameters_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCustomerCertificateParameters_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["SecretSource"] = gen.PtrOf(ResourceReference_STATUS_ARMGenerator())
}

func Test_ManagedCertificateParameters_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedCertificateParameters_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedCertificateParameters_STATUS_ARM, ManagedCertificateParameters_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedCertificateParameters_STATUS_ARM runs a test to see if a specific instance of ManagedCertificateParameters_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedCertificateParameters_STATUS_ARM(subject ManagedCertificateParameters_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedCertificateParameters_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ManagedCertificateParameters_STATUS_ARM instances for property testing - lazily instantiated by
// ManagedCertificateParameters_STATUS_ARMGenerator()
var managedCertificateParameters_STATUS_ARMGenerator gopter.Gen

// ManagedCertificateParameters_STATUS_ARMGenerator returns a generator of ManagedCertificateParameters_STATUS_ARM instances for property testing.
func ManagedCertificateParameters_STATUS_ARMGenerator() gopter.Gen {
	if managedCertificateParameters_STATUS_ARMGenerator != nil {
		return managedCertificateParameters_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedCertificateParameters_STATUS_ARM(generators)
	managedCertificateParameters_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ManagedCertificateParameters_STATUS_ARM{}), generators)

	return managedCertificateParameters_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForManagedCertificateParameters_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedCertificateParameters_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ExpirationDate"] = gen.PtrOf(gen.AlphaString())
	gens["Subject"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.OneConstOf(ManagedCertificateParameters_Type_STATUS_ManagedCertificate)
}

func Test_Profiles_Secret_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Profiles_Secret_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfiles_Secret_STATUS_ARM, Profiles_Secret_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfiles_Secret_STATUS_ARM runs a test to see if a specific instance of Profiles_Secret_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForProfiles_Secret_STATUS_ARM(subject Profiles_Secret_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Profiles_Secret_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Profiles_Secret_STATUS_ARM instances for property testing - lazily instantiated by
// Profiles_Secret_STATUS_ARMGenerator()
var profiles_Secret_STATUS_ARMGenerator gopter.Gen

// Profiles_Secret_STATUS_ARMGenerator returns a generator of Profiles_Secret_STATUS_ARM instances for property testing.
// We first initialize profiles_Secret_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Profiles_Secret_STATUS_ARMGenerator() gopter.Gen {
	if profiles_Secret_STATUS_ARMGenerator != nil {
		return profiles_Secret_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfiles_Secret_STATUS_ARM(generators)
	profiles_Secret_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Profiles_Secret_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfiles_Secret_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForProfiles_Secret_STATUS_ARM(generators)
	profiles_Secret_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Profiles_Secret_STATUS_ARM{}), generators)

	return profiles_Secret_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForProfiles_Secret_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProfiles_Secret_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForProfiles_Secret_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProfiles_Secret_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SecretProperties_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}

func Test_SecretParameters_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecretParameters_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecretParameters_STATUS_ARM, SecretParameters_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecretParameters_STATUS_ARM runs a test to see if a specific instance of SecretParameters_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSecretParameters_STATUS_ARM(subject SecretParameters_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecretParameters_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SecretParameters_STATUS_ARM instances for property testing - lazily instantiated by
// SecretParameters_STATUS_ARMGenerator()
var secretParameters_STATUS_ARMGenerator gopter.Gen

// SecretParameters_STATUS_ARMGenerator returns a generator of SecretParameters_STATUS_ARM instances for property testing.
func SecretParameters_STATUS_ARMGenerator() gopter.Gen {
	if secretParameters_STATUS_ARMGenerator != nil {
		return secretParameters_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSecretParameters_STATUS_ARM(generators)

	// handle OneOf by choosing only one field to instantiate
	var gens []gopter.Gen
	for propName, propGen := range generators {
		gens = append(gens, gen.Struct(reflect.TypeOf(SecretParameters_STATUS_ARM{}), map[string]gopter.Gen{propName: propGen}))
	}
	secretParameters_STATUS_ARMGenerator = gen.OneGenOf(gens...)

	return secretParameters_STATUS_ARMGenerator
}

// AddRelatedPropertyGeneratorsForSecretParameters_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecretParameters_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AzureFirstPartyManagedCertificate"] = AzureFirstPartyManagedCertificateParameters_STATUS_ARMGenerator().Map(func(it AzureFirstPartyManagedCertificateParameters_STATUS_ARM) *AzureFirstPartyManagedCertificateParameters_STATUS_ARM {
		return &it
	}) // generate one case for OneOf type
	gens["CustomerCertificate"] = CustomerCertificateParameters_STATUS_ARMGenerator().Map(func(it CustomerCertificateParameters_STATUS_ARM) *CustomerCertificateParameters_STATUS_ARM {
		return &it
	}) // generate one case for OneOf type
	gens["ManagedCertificate"] = ManagedCertificateParameters_STATUS_ARMGenerator().Map(func(it ManagedCertificateParameters_STATUS_ARM) *ManagedCertificateParameters_STATUS_ARM {
		return &it
	}) // generate one case for OneOf type
	gens["UrlSigningKey"] = UrlSigningKeyParameters_STATUS_ARMGenerator().Map(func(it UrlSigningKeyParameters_STATUS_ARM) *UrlSigningKeyParameters_STATUS_ARM {
		return &it
	}) // generate one case for OneOf type
}

func Test_SecretProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecretProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecretProperties_STATUS_ARM, SecretProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecretProperties_STATUS_ARM runs a test to see if a specific instance of SecretProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSecretProperties_STATUS_ARM(subject SecretProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecretProperties_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SecretProperties_STATUS_ARM instances for property testing - lazily instantiated by
// SecretProperties_STATUS_ARMGenerator()
var secretProperties_STATUS_ARMGenerator gopter.Gen

// SecretProperties_STATUS_ARMGenerator returns a generator of SecretProperties_STATUS_ARM instances for property testing.
// We first initialize secretProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SecretProperties_STATUS_ARMGenerator() gopter.Gen {
	if secretProperties_STATUS_ARMGenerator != nil {
		return secretProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecretProperties_STATUS_ARM(generators)
	secretProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SecretProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecretProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForSecretProperties_STATUS_ARM(generators)
	secretProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SecretProperties_STATUS_ARM{}), generators)

	return secretProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSecretProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecretProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["DeploymentStatus"] = gen.PtrOf(gen.OneConstOf(
		SecretProperties_DeploymentStatus_STATUS_Failed,
		SecretProperties_DeploymentStatus_STATUS_InProgress,
		SecretProperties_DeploymentStatus_STATUS_NotStarted,
		SecretProperties_DeploymentStatus_STATUS_Succeeded))
	gens["ProfileName"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		SecretProperties_ProvisioningState_STATUS_Creating,
		SecretProperties_ProvisioningState_STATUS_Deleting,
		SecretProperties_ProvisioningState_STATUS_Failed,
		SecretProperties_ProvisioningState_STATUS_Succeeded,
		SecretProperties_ProvisioningState_STATUS_Updating))
}

// AddRelatedPropertyGeneratorsForSecretProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecretProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Parameters"] = gen.PtrOf(SecretParameters_STATUS_ARMGenerator())
}

func Test_UrlSigningKeyParameters_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UrlSigningKeyParameters_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUrlSigningKeyParameters_STATUS_ARM, UrlSigningKeyParameters_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUrlSigningKeyParameters_STATUS_ARM runs a test to see if a specific instance of UrlSigningKeyParameters_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUrlSigningKeyParameters_STATUS_ARM(subject UrlSigningKeyParameters_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UrlSigningKeyParameters_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of UrlSigningKeyParameters_STATUS_ARM instances for property testing - lazily instantiated by
// UrlSigningKeyParameters_STATUS_ARMGenerator()
var urlSigningKeyParameters_STATUS_ARMGenerator gopter.Gen

// UrlSigningKeyParameters_STATUS_ARMGenerator returns a generator of UrlSigningKeyParameters_STATUS_ARM instances for property testing.
// We first initialize urlSigningKeyParameters_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func UrlSigningKeyParameters_STATUS_ARMGenerator() gopter.Gen {
	if urlSigningKeyParameters_STATUS_ARMGenerator != nil {
		return urlSigningKeyParameters_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUrlSigningKeyParameters_STATUS_ARM(generators)
	urlSigningKeyParameters_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(UrlSigningKeyParameters_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUrlSigningKeyParameters_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForUrlSigningKeyParameters_STATUS_ARM(generators)
	urlSigningKeyParameters_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(UrlSigningKeyParameters_STATUS_ARM{}), generators)

	return urlSigningKeyParameters_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForUrlSigningKeyParameters_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUrlSigningKeyParameters_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["KeyId"] = gen.PtrOf(gen.AlphaString())
	gens["SecretVersion"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.OneConstOf(UrlSigningKeyParameters_Type_STATUS_UrlSigningKey)
}

// AddRelatedPropertyGeneratorsForUrlSigningKeyParameters_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForUrlSigningKeyParameters_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["SecretSource"] = gen.PtrOf(ResourceReference_STATUS_ARMGenerator())
}