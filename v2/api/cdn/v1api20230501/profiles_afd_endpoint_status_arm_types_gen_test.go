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

func Test_AFDEndpointProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AFDEndpointProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAFDEndpointProperties_STATUS_ARM, AFDEndpointProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAFDEndpointProperties_STATUS_ARM runs a test to see if a specific instance of AFDEndpointProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAFDEndpointProperties_STATUS_ARM(subject AFDEndpointProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AFDEndpointProperties_STATUS_ARM
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

// Generator of AFDEndpointProperties_STATUS_ARM instances for property testing - lazily instantiated by
// AFDEndpointProperties_STATUS_ARMGenerator()
var afdEndpointProperties_STATUS_ARMGenerator gopter.Gen

// AFDEndpointProperties_STATUS_ARMGenerator returns a generator of AFDEndpointProperties_STATUS_ARM instances for property testing.
func AFDEndpointProperties_STATUS_ARMGenerator() gopter.Gen {
	if afdEndpointProperties_STATUS_ARMGenerator != nil {
		return afdEndpointProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAFDEndpointProperties_STATUS_ARM(generators)
	afdEndpointProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AFDEndpointProperties_STATUS_ARM{}), generators)

	return afdEndpointProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAFDEndpointProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAFDEndpointProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AutoGeneratedDomainNameLabelScope"] = gen.PtrOf(gen.OneConstOf(
		AutoGeneratedDomainNameLabelScope_STATUS_NoReuse,
		AutoGeneratedDomainNameLabelScope_STATUS_ResourceGroupReuse,
		AutoGeneratedDomainNameLabelScope_STATUS_SubscriptionReuse,
		AutoGeneratedDomainNameLabelScope_STATUS_TenantReuse))
	gens["DeploymentStatus"] = gen.PtrOf(gen.OneConstOf(
		AFDEndpointProperties_DeploymentStatus_STATUS_Failed,
		AFDEndpointProperties_DeploymentStatus_STATUS_InProgress,
		AFDEndpointProperties_DeploymentStatus_STATUS_NotStarted,
		AFDEndpointProperties_DeploymentStatus_STATUS_Succeeded))
	gens["EnabledState"] = gen.PtrOf(gen.OneConstOf(AFDEndpointProperties_EnabledState_STATUS_Disabled, AFDEndpointProperties_EnabledState_STATUS_Enabled))
	gens["HostName"] = gen.PtrOf(gen.AlphaString())
	gens["ProfileName"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		AFDEndpointProperties_ProvisioningState_STATUS_Creating,
		AFDEndpointProperties_ProvisioningState_STATUS_Deleting,
		AFDEndpointProperties_ProvisioningState_STATUS_Failed,
		AFDEndpointProperties_ProvisioningState_STATUS_Succeeded,
		AFDEndpointProperties_ProvisioningState_STATUS_Updating))
}

func Test_Profiles_AfdEndpoint_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Profiles_AfdEndpoint_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfiles_AfdEndpoint_STATUS_ARM, Profiles_AfdEndpoint_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfiles_AfdEndpoint_STATUS_ARM runs a test to see if a specific instance of Profiles_AfdEndpoint_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForProfiles_AfdEndpoint_STATUS_ARM(subject Profiles_AfdEndpoint_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Profiles_AfdEndpoint_STATUS_ARM
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

// Generator of Profiles_AfdEndpoint_STATUS_ARM instances for property testing - lazily instantiated by
// Profiles_AfdEndpoint_STATUS_ARMGenerator()
var profiles_AfdEndpoint_STATUS_ARMGenerator gopter.Gen

// Profiles_AfdEndpoint_STATUS_ARMGenerator returns a generator of Profiles_AfdEndpoint_STATUS_ARM instances for property testing.
// We first initialize profiles_AfdEndpoint_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Profiles_AfdEndpoint_STATUS_ARMGenerator() gopter.Gen {
	if profiles_AfdEndpoint_STATUS_ARMGenerator != nil {
		return profiles_AfdEndpoint_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfiles_AfdEndpoint_STATUS_ARM(generators)
	profiles_AfdEndpoint_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Profiles_AfdEndpoint_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfiles_AfdEndpoint_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForProfiles_AfdEndpoint_STATUS_ARM(generators)
	profiles_AfdEndpoint_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Profiles_AfdEndpoint_STATUS_ARM{}), generators)

	return profiles_AfdEndpoint_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForProfiles_AfdEndpoint_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProfiles_AfdEndpoint_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForProfiles_AfdEndpoint_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProfiles_AfdEndpoint_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(AFDEndpointProperties_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}