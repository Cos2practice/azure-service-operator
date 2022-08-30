// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210601storage

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

func Test_FlexibleServersConfiguration_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersConfiguration via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersConfiguration, FlexibleServersConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersConfiguration runs a test to see if a specific instance of FlexibleServersConfiguration round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersConfiguration(subject FlexibleServersConfiguration) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersConfiguration
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

// Generator of FlexibleServersConfiguration instances for property testing - lazily instantiated by
// FlexibleServersConfigurationGenerator()
var flexibleServersConfigurationGenerator gopter.Gen

// FlexibleServersConfigurationGenerator returns a generator of FlexibleServersConfiguration instances for property testing.
func FlexibleServersConfigurationGenerator() gopter.Gen {
	if flexibleServersConfigurationGenerator != nil {
		return flexibleServersConfigurationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForFlexibleServersConfiguration(generators)
	flexibleServersConfigurationGenerator = gen.Struct(reflect.TypeOf(FlexibleServersConfiguration{}), generators)

	return flexibleServersConfigurationGenerator
}

// AddRelatedPropertyGeneratorsForFlexibleServersConfiguration is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersConfiguration(gens map[string]gopter.Gen) {
	gens["Spec"] = FlexibleServers_Configurations_SpecGenerator()
	gens["Status"] = Configuration_STATUSGenerator()
}

func Test_Configuration_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Configuration_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConfiguration_STATUS, Configuration_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConfiguration_STATUS runs a test to see if a specific instance of Configuration_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForConfiguration_STATUS(subject Configuration_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Configuration_STATUS
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

// Generator of Configuration_STATUS instances for property testing - lazily instantiated by
// Configuration_STATUSGenerator()
var configuration_STATUSGenerator gopter.Gen

// Configuration_STATUSGenerator returns a generator of Configuration_STATUS instances for property testing.
// We first initialize configuration_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Configuration_STATUSGenerator() gopter.Gen {
	if configuration_STATUSGenerator != nil {
		return configuration_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConfiguration_STATUS(generators)
	configuration_STATUSGenerator = gen.Struct(reflect.TypeOf(Configuration_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConfiguration_STATUS(generators)
	AddRelatedPropertyGeneratorsForConfiguration_STATUS(generators)
	configuration_STATUSGenerator = gen.Struct(reflect.TypeOf(Configuration_STATUS{}), generators)

	return configuration_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForConfiguration_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConfiguration_STATUS(gens map[string]gopter.Gen) {
	gens["AllowedValues"] = gen.PtrOf(gen.AlphaString())
	gens["DataType"] = gen.PtrOf(gen.AlphaString())
	gens["DefaultValue"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DocumentationLink"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IsConfigPendingRestart"] = gen.PtrOf(gen.Bool())
	gens["IsDynamicConfig"] = gen.PtrOf(gen.Bool())
	gens["IsReadOnly"] = gen.PtrOf(gen.Bool())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Source"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Unit"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForConfiguration_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForConfiguration_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_FlexibleServers_Configurations_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServers_Configurations_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServers_Configurations_Spec, FlexibleServers_Configurations_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServers_Configurations_Spec runs a test to see if a specific instance of FlexibleServers_Configurations_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServers_Configurations_Spec(subject FlexibleServers_Configurations_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServers_Configurations_Spec
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

// Generator of FlexibleServers_Configurations_Spec instances for property testing - lazily instantiated by
// FlexibleServers_Configurations_SpecGenerator()
var flexibleServers_Configurations_SpecGenerator gopter.Gen

// FlexibleServers_Configurations_SpecGenerator returns a generator of FlexibleServers_Configurations_Spec instances for property testing.
func FlexibleServers_Configurations_SpecGenerator() gopter.Gen {
	if flexibleServers_Configurations_SpecGenerator != nil {
		return flexibleServers_Configurations_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServers_Configurations_Spec(generators)
	flexibleServers_Configurations_SpecGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_Configurations_Spec{}), generators)

	return flexibleServers_Configurations_SpecGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServers_Configurations_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServers_Configurations_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Source"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}