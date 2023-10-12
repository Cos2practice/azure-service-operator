// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20200601

import (
	"encoding/json"
	v20200601s "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1api20200601storage"
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

func Test_DomainsTopic_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DomainsTopic to hub returns original",
		prop.ForAll(RunResourceConversionTestForDomainsTopic, DomainsTopicGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForDomainsTopic tests if a specific instance of DomainsTopic round trips to the hub storage version and back losslessly
func RunResourceConversionTestForDomainsTopic(subject DomainsTopic) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20200601s.DomainsTopic
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual DomainsTopic
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_DomainsTopic_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DomainsTopic to DomainsTopic via AssignProperties_To_DomainsTopic & AssignProperties_From_DomainsTopic returns original",
		prop.ForAll(RunPropertyAssignmentTestForDomainsTopic, DomainsTopicGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDomainsTopic tests if a specific instance of DomainsTopic can be assigned to v1api20200601storage and back losslessly
func RunPropertyAssignmentTestForDomainsTopic(subject DomainsTopic) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20200601s.DomainsTopic
	err := copied.AssignProperties_To_DomainsTopic(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DomainsTopic
	err = actual.AssignProperties_From_DomainsTopic(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_DomainsTopic_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DomainsTopic via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomainsTopic, DomainsTopicGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomainsTopic runs a test to see if a specific instance of DomainsTopic round trips to JSON and back losslessly
func RunJSONSerializationTestForDomainsTopic(subject DomainsTopic) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DomainsTopic
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

// Generator of DomainsTopic instances for property testing - lazily instantiated by DomainsTopicGenerator()
var domainsTopicGenerator gopter.Gen

// DomainsTopicGenerator returns a generator of DomainsTopic instances for property testing.
func DomainsTopicGenerator() gopter.Gen {
	if domainsTopicGenerator != nil {
		return domainsTopicGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDomainsTopic(generators)
	domainsTopicGenerator = gen.Struct(reflect.TypeOf(DomainsTopic{}), generators)

	return domainsTopicGenerator
}

// AddRelatedPropertyGeneratorsForDomainsTopic is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomainsTopic(gens map[string]gopter.Gen) {
	gens["Spec"] = Domains_Topic_SpecGenerator()
	gens["Status"] = Domains_Topic_STATUSGenerator()
}

func Test_Domains_Topic_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Domains_Topic_Spec to Domains_Topic_Spec via AssignProperties_To_Domains_Topic_Spec & AssignProperties_From_Domains_Topic_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDomains_Topic_Spec, Domains_Topic_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDomains_Topic_Spec tests if a specific instance of Domains_Topic_Spec can be assigned to v1api20200601storage and back losslessly
func RunPropertyAssignmentTestForDomains_Topic_Spec(subject Domains_Topic_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20200601s.Domains_Topic_Spec
	err := copied.AssignProperties_To_Domains_Topic_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Domains_Topic_Spec
	err = actual.AssignProperties_From_Domains_Topic_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Domains_Topic_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Domains_Topic_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomains_Topic_Spec, Domains_Topic_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomains_Topic_Spec runs a test to see if a specific instance of Domains_Topic_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDomains_Topic_Spec(subject Domains_Topic_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Domains_Topic_Spec
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

// Generator of Domains_Topic_Spec instances for property testing - lazily instantiated by Domains_Topic_SpecGenerator()
var domains_Topic_SpecGenerator gopter.Gen

// Domains_Topic_SpecGenerator returns a generator of Domains_Topic_Spec instances for property testing.
func Domains_Topic_SpecGenerator() gopter.Gen {
	if domains_Topic_SpecGenerator != nil {
		return domains_Topic_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomains_Topic_Spec(generators)
	domains_Topic_SpecGenerator = gen.Struct(reflect.TypeOf(Domains_Topic_Spec{}), generators)

	return domains_Topic_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDomains_Topic_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomains_Topic_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
}

func Test_Domains_Topic_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Domains_Topic_STATUS to Domains_Topic_STATUS via AssignProperties_To_Domains_Topic_STATUS & AssignProperties_From_Domains_Topic_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForDomains_Topic_STATUS, Domains_Topic_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDomains_Topic_STATUS tests if a specific instance of Domains_Topic_STATUS can be assigned to v1api20200601storage and back losslessly
func RunPropertyAssignmentTestForDomains_Topic_STATUS(subject Domains_Topic_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20200601s.Domains_Topic_STATUS
	err := copied.AssignProperties_To_Domains_Topic_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Domains_Topic_STATUS
	err = actual.AssignProperties_From_Domains_Topic_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Domains_Topic_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Domains_Topic_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomains_Topic_STATUS, Domains_Topic_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomains_Topic_STATUS runs a test to see if a specific instance of Domains_Topic_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDomains_Topic_STATUS(subject Domains_Topic_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Domains_Topic_STATUS
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

// Generator of Domains_Topic_STATUS instances for property testing - lazily instantiated by
// Domains_Topic_STATUSGenerator()
var domains_Topic_STATUSGenerator gopter.Gen

// Domains_Topic_STATUSGenerator returns a generator of Domains_Topic_STATUS instances for property testing.
// We first initialize domains_Topic_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Domains_Topic_STATUSGenerator() gopter.Gen {
	if domains_Topic_STATUSGenerator != nil {
		return domains_Topic_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomains_Topic_STATUS(generators)
	domains_Topic_STATUSGenerator = gen.Struct(reflect.TypeOf(Domains_Topic_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomains_Topic_STATUS(generators)
	AddRelatedPropertyGeneratorsForDomains_Topic_STATUS(generators)
	domains_Topic_STATUSGenerator = gen.Struct(reflect.TypeOf(Domains_Topic_STATUS{}), generators)

	return domains_Topic_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDomains_Topic_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomains_Topic_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		DomainTopicProperties_ProvisioningState_STATUS_Canceled,
		DomainTopicProperties_ProvisioningState_STATUS_Creating,
		DomainTopicProperties_ProvisioningState_STATUS_Deleting,
		DomainTopicProperties_ProvisioningState_STATUS_Failed,
		DomainTopicProperties_ProvisioningState_STATUS_Succeeded,
		DomainTopicProperties_ProvisioningState_STATUS_Updating))
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDomains_Topic_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomains_Topic_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}