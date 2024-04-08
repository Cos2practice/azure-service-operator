// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20201201

import (
	"encoding/json"
	v20201201s "github.com/Azure/azure-service-operator/v2/api/compute/v1api20201201/storage"
	v20220301s "github.com/Azure/azure-service-operator/v2/api/compute/v1api20220301/storage"
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

func Test_VirtualMachinesExtension_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualMachinesExtension to hub returns original",
		prop.ForAll(RunResourceConversionTestForVirtualMachinesExtension, VirtualMachinesExtensionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForVirtualMachinesExtension tests if a specific instance of VirtualMachinesExtension round trips to the hub storage version and back losslessly
func RunResourceConversionTestForVirtualMachinesExtension(subject VirtualMachinesExtension) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20220301s.VirtualMachinesExtension
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual VirtualMachinesExtension
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

func Test_VirtualMachinesExtension_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualMachinesExtension to VirtualMachinesExtension via AssignProperties_To_VirtualMachinesExtension & AssignProperties_From_VirtualMachinesExtension returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualMachinesExtension, VirtualMachinesExtensionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualMachinesExtension tests if a specific instance of VirtualMachinesExtension can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForVirtualMachinesExtension(subject VirtualMachinesExtension) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201201s.VirtualMachinesExtension
	err := copied.AssignProperties_To_VirtualMachinesExtension(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualMachinesExtension
	err = actual.AssignProperties_From_VirtualMachinesExtension(&other)
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

func Test_VirtualMachinesExtension_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachinesExtension via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachinesExtension, VirtualMachinesExtensionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachinesExtension runs a test to see if a specific instance of VirtualMachinesExtension round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachinesExtension(subject VirtualMachinesExtension) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachinesExtension
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

// Generator of VirtualMachinesExtension instances for property testing - lazily instantiated by
// VirtualMachinesExtensionGenerator()
var virtualMachinesExtensionGenerator gopter.Gen

// VirtualMachinesExtensionGenerator returns a generator of VirtualMachinesExtension instances for property testing.
func VirtualMachinesExtensionGenerator() gopter.Gen {
	if virtualMachinesExtensionGenerator != nil {
		return virtualMachinesExtensionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForVirtualMachinesExtension(generators)
	virtualMachinesExtensionGenerator = gen.Struct(reflect.TypeOf(VirtualMachinesExtension{}), generators)

	return virtualMachinesExtensionGenerator
}

// AddRelatedPropertyGeneratorsForVirtualMachinesExtension is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualMachinesExtension(gens map[string]gopter.Gen) {
	gens["Spec"] = VirtualMachines_Extension_SpecGenerator()
	gens["Status"] = VirtualMachines_Extension_STATUSGenerator()
}

func Test_VirtualMachines_Extension_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualMachines_Extension_Spec to VirtualMachines_Extension_Spec via AssignProperties_To_VirtualMachines_Extension_Spec & AssignProperties_From_VirtualMachines_Extension_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualMachines_Extension_Spec, VirtualMachines_Extension_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualMachines_Extension_Spec tests if a specific instance of VirtualMachines_Extension_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForVirtualMachines_Extension_Spec(subject VirtualMachines_Extension_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201201s.VirtualMachines_Extension_Spec
	err := copied.AssignProperties_To_VirtualMachines_Extension_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualMachines_Extension_Spec
	err = actual.AssignProperties_From_VirtualMachines_Extension_Spec(&other)
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

func Test_VirtualMachines_Extension_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachines_Extension_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachines_Extension_Spec, VirtualMachines_Extension_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachines_Extension_Spec runs a test to see if a specific instance of VirtualMachines_Extension_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachines_Extension_Spec(subject VirtualMachines_Extension_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachines_Extension_Spec
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

// Generator of VirtualMachines_Extension_Spec instances for property testing - lazily instantiated by
// VirtualMachines_Extension_SpecGenerator()
var virtualMachines_Extension_SpecGenerator gopter.Gen

// VirtualMachines_Extension_SpecGenerator returns a generator of VirtualMachines_Extension_Spec instances for property testing.
// We first initialize virtualMachines_Extension_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualMachines_Extension_SpecGenerator() gopter.Gen {
	if virtualMachines_Extension_SpecGenerator != nil {
		return virtualMachines_Extension_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachines_Extension_Spec(generators)
	virtualMachines_Extension_SpecGenerator = gen.Struct(reflect.TypeOf(VirtualMachines_Extension_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachines_Extension_Spec(generators)
	AddRelatedPropertyGeneratorsForVirtualMachines_Extension_Spec(generators)
	virtualMachines_Extension_SpecGenerator = gen.Struct(reflect.TypeOf(VirtualMachines_Extension_Spec{}), generators)

	return virtualMachines_Extension_SpecGenerator
}

// AddIndependentPropertyGeneratorsForVirtualMachines_Extension_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualMachines_Extension_Spec(gens map[string]gopter.Gen) {
	gens["AutoUpgradeMinorVersion"] = gen.PtrOf(gen.Bool())
	gens["AzureName"] = gen.AlphaString()
	gens["EnableAutomaticUpgrade"] = gen.PtrOf(gen.Bool())
	gens["ForceUpdateTag"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Publisher"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["TypeHandlerVersion"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualMachines_Extension_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualMachines_Extension_Spec(gens map[string]gopter.Gen) {
	gens["InstanceView"] = gen.PtrOf(VirtualMachineExtensionInstanceViewGenerator())
}

func Test_VirtualMachines_Extension_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualMachines_Extension_STATUS to VirtualMachines_Extension_STATUS via AssignProperties_To_VirtualMachines_Extension_STATUS & AssignProperties_From_VirtualMachines_Extension_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualMachines_Extension_STATUS, VirtualMachines_Extension_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualMachines_Extension_STATUS tests if a specific instance of VirtualMachines_Extension_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForVirtualMachines_Extension_STATUS(subject VirtualMachines_Extension_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201201s.VirtualMachines_Extension_STATUS
	err := copied.AssignProperties_To_VirtualMachines_Extension_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualMachines_Extension_STATUS
	err = actual.AssignProperties_From_VirtualMachines_Extension_STATUS(&other)
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

func Test_VirtualMachines_Extension_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachines_Extension_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachines_Extension_STATUS, VirtualMachines_Extension_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachines_Extension_STATUS runs a test to see if a specific instance of VirtualMachines_Extension_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachines_Extension_STATUS(subject VirtualMachines_Extension_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachines_Extension_STATUS
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

// Generator of VirtualMachines_Extension_STATUS instances for property testing - lazily instantiated by
// VirtualMachines_Extension_STATUSGenerator()
var virtualMachines_Extension_STATUSGenerator gopter.Gen

// VirtualMachines_Extension_STATUSGenerator returns a generator of VirtualMachines_Extension_STATUS instances for property testing.
// We first initialize virtualMachines_Extension_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualMachines_Extension_STATUSGenerator() gopter.Gen {
	if virtualMachines_Extension_STATUSGenerator != nil {
		return virtualMachines_Extension_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachines_Extension_STATUS(generators)
	virtualMachines_Extension_STATUSGenerator = gen.Struct(reflect.TypeOf(VirtualMachines_Extension_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachines_Extension_STATUS(generators)
	AddRelatedPropertyGeneratorsForVirtualMachines_Extension_STATUS(generators)
	virtualMachines_Extension_STATUSGenerator = gen.Struct(reflect.TypeOf(VirtualMachines_Extension_STATUS{}), generators)

	return virtualMachines_Extension_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForVirtualMachines_Extension_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualMachines_Extension_STATUS(gens map[string]gopter.Gen) {
	gens["AutoUpgradeMinorVersion"] = gen.PtrOf(gen.Bool())
	gens["EnableAutomaticUpgrade"] = gen.PtrOf(gen.Bool())
	gens["ForceUpdateTag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PropertiesType"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["Publisher"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["TypeHandlerVersion"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualMachines_Extension_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualMachines_Extension_STATUS(gens map[string]gopter.Gen) {
	gens["InstanceView"] = gen.PtrOf(VirtualMachineExtensionInstanceView_STATUSGenerator())
}

func Test_VirtualMachineExtensionInstanceView_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualMachineExtensionInstanceView to VirtualMachineExtensionInstanceView via AssignProperties_To_VirtualMachineExtensionInstanceView & AssignProperties_From_VirtualMachineExtensionInstanceView returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualMachineExtensionInstanceView, VirtualMachineExtensionInstanceViewGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualMachineExtensionInstanceView tests if a specific instance of VirtualMachineExtensionInstanceView can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForVirtualMachineExtensionInstanceView(subject VirtualMachineExtensionInstanceView) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201201s.VirtualMachineExtensionInstanceView
	err := copied.AssignProperties_To_VirtualMachineExtensionInstanceView(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualMachineExtensionInstanceView
	err = actual.AssignProperties_From_VirtualMachineExtensionInstanceView(&other)
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

func Test_VirtualMachineExtensionInstanceView_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachineExtensionInstanceView via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachineExtensionInstanceView, VirtualMachineExtensionInstanceViewGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachineExtensionInstanceView runs a test to see if a specific instance of VirtualMachineExtensionInstanceView round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachineExtensionInstanceView(subject VirtualMachineExtensionInstanceView) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachineExtensionInstanceView
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

// Generator of VirtualMachineExtensionInstanceView instances for property testing - lazily instantiated by
// VirtualMachineExtensionInstanceViewGenerator()
var virtualMachineExtensionInstanceViewGenerator gopter.Gen

// VirtualMachineExtensionInstanceViewGenerator returns a generator of VirtualMachineExtensionInstanceView instances for property testing.
// We first initialize virtualMachineExtensionInstanceViewGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualMachineExtensionInstanceViewGenerator() gopter.Gen {
	if virtualMachineExtensionInstanceViewGenerator != nil {
		return virtualMachineExtensionInstanceViewGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineExtensionInstanceView(generators)
	virtualMachineExtensionInstanceViewGenerator = gen.Struct(reflect.TypeOf(VirtualMachineExtensionInstanceView{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineExtensionInstanceView(generators)
	AddRelatedPropertyGeneratorsForVirtualMachineExtensionInstanceView(generators)
	virtualMachineExtensionInstanceViewGenerator = gen.Struct(reflect.TypeOf(VirtualMachineExtensionInstanceView{}), generators)

	return virtualMachineExtensionInstanceViewGenerator
}

// AddIndependentPropertyGeneratorsForVirtualMachineExtensionInstanceView is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualMachineExtensionInstanceView(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["TypeHandlerVersion"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualMachineExtensionInstanceView is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualMachineExtensionInstanceView(gens map[string]gopter.Gen) {
	gens["Statuses"] = gen.SliceOf(InstanceViewStatusGenerator())
	gens["Substatuses"] = gen.SliceOf(InstanceViewStatusGenerator())
}

func Test_VirtualMachineExtensionInstanceView_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualMachineExtensionInstanceView_STATUS to VirtualMachineExtensionInstanceView_STATUS via AssignProperties_To_VirtualMachineExtensionInstanceView_STATUS & AssignProperties_From_VirtualMachineExtensionInstanceView_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualMachineExtensionInstanceView_STATUS, VirtualMachineExtensionInstanceView_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualMachineExtensionInstanceView_STATUS tests if a specific instance of VirtualMachineExtensionInstanceView_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForVirtualMachineExtensionInstanceView_STATUS(subject VirtualMachineExtensionInstanceView_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201201s.VirtualMachineExtensionInstanceView_STATUS
	err := copied.AssignProperties_To_VirtualMachineExtensionInstanceView_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualMachineExtensionInstanceView_STATUS
	err = actual.AssignProperties_From_VirtualMachineExtensionInstanceView_STATUS(&other)
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

func Test_VirtualMachineExtensionInstanceView_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachineExtensionInstanceView_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachineExtensionInstanceView_STATUS, VirtualMachineExtensionInstanceView_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachineExtensionInstanceView_STATUS runs a test to see if a specific instance of VirtualMachineExtensionInstanceView_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachineExtensionInstanceView_STATUS(subject VirtualMachineExtensionInstanceView_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachineExtensionInstanceView_STATUS
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

// Generator of VirtualMachineExtensionInstanceView_STATUS instances for property testing - lazily instantiated by
// VirtualMachineExtensionInstanceView_STATUSGenerator()
var virtualMachineExtensionInstanceView_STATUSGenerator gopter.Gen

// VirtualMachineExtensionInstanceView_STATUSGenerator returns a generator of VirtualMachineExtensionInstanceView_STATUS instances for property testing.
// We first initialize virtualMachineExtensionInstanceView_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualMachineExtensionInstanceView_STATUSGenerator() gopter.Gen {
	if virtualMachineExtensionInstanceView_STATUSGenerator != nil {
		return virtualMachineExtensionInstanceView_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineExtensionInstanceView_STATUS(generators)
	virtualMachineExtensionInstanceView_STATUSGenerator = gen.Struct(reflect.TypeOf(VirtualMachineExtensionInstanceView_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineExtensionInstanceView_STATUS(generators)
	AddRelatedPropertyGeneratorsForVirtualMachineExtensionInstanceView_STATUS(generators)
	virtualMachineExtensionInstanceView_STATUSGenerator = gen.Struct(reflect.TypeOf(VirtualMachineExtensionInstanceView_STATUS{}), generators)

	return virtualMachineExtensionInstanceView_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForVirtualMachineExtensionInstanceView_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualMachineExtensionInstanceView_STATUS(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["TypeHandlerVersion"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualMachineExtensionInstanceView_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualMachineExtensionInstanceView_STATUS(gens map[string]gopter.Gen) {
	gens["Statuses"] = gen.SliceOf(InstanceViewStatus_STATUSGenerator())
	gens["Substatuses"] = gen.SliceOf(InstanceViewStatus_STATUSGenerator())
}

func Test_InstanceViewStatus_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from InstanceViewStatus to InstanceViewStatus via AssignProperties_To_InstanceViewStatus & AssignProperties_From_InstanceViewStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForInstanceViewStatus, InstanceViewStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForInstanceViewStatus tests if a specific instance of InstanceViewStatus can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForInstanceViewStatus(subject InstanceViewStatus) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201201s.InstanceViewStatus
	err := copied.AssignProperties_To_InstanceViewStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual InstanceViewStatus
	err = actual.AssignProperties_From_InstanceViewStatus(&other)
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

func Test_InstanceViewStatus_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of InstanceViewStatus via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForInstanceViewStatus, InstanceViewStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInstanceViewStatus runs a test to see if a specific instance of InstanceViewStatus round trips to JSON and back losslessly
func RunJSONSerializationTestForInstanceViewStatus(subject InstanceViewStatus) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual InstanceViewStatus
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

// Generator of InstanceViewStatus instances for property testing - lazily instantiated by InstanceViewStatusGenerator()
var instanceViewStatusGenerator gopter.Gen

// InstanceViewStatusGenerator returns a generator of InstanceViewStatus instances for property testing.
func InstanceViewStatusGenerator() gopter.Gen {
	if instanceViewStatusGenerator != nil {
		return instanceViewStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInstanceViewStatus(generators)
	instanceViewStatusGenerator = gen.Struct(reflect.TypeOf(InstanceViewStatus{}), generators)

	return instanceViewStatusGenerator
}

// AddIndependentPropertyGeneratorsForInstanceViewStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForInstanceViewStatus(gens map[string]gopter.Gen) {
	gens["Code"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayStatus"] = gen.PtrOf(gen.AlphaString())
	gens["Level"] = gen.PtrOf(gen.OneConstOf(InstanceViewStatus_Level_Error, InstanceViewStatus_Level_Info, InstanceViewStatus_Level_Warning))
	gens["Message"] = gen.PtrOf(gen.AlphaString())
	gens["Time"] = gen.PtrOf(gen.AlphaString())
}