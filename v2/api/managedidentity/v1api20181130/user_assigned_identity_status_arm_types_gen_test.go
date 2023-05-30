// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20181130

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

func Test_UserAssignedIdentity_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentity_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentity_STATUS_ARM, UserAssignedIdentity_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentity_STATUS_ARM runs a test to see if a specific instance of UserAssignedIdentity_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentity_STATUS_ARM(subject UserAssignedIdentity_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentity_STATUS_ARM
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

// Generator of UserAssignedIdentity_STATUS_ARM instances for property testing - lazily instantiated by
// UserAssignedIdentity_STATUS_ARMGenerator()
var userAssignedIdentity_STATUS_ARMGenerator gopter.Gen

// UserAssignedIdentity_STATUS_ARMGenerator returns a generator of UserAssignedIdentity_STATUS_ARM instances for property testing.
// We first initialize userAssignedIdentity_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func UserAssignedIdentity_STATUS_ARMGenerator() gopter.Gen {
	if userAssignedIdentity_STATUS_ARMGenerator != nil {
		return userAssignedIdentity_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUS_ARM(generators)
	userAssignedIdentity_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentity_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForUserAssignedIdentity_STATUS_ARM(generators)
	userAssignedIdentity_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentity_STATUS_ARM{}), generators)

	return userAssignedIdentity_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForUserAssignedIdentity_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForUserAssignedIdentity_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(UserAssignedIdentityProperties_STATUS_ARMGenerator())
}

func Test_UserAssignedIdentityProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentityProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentityProperties_STATUS_ARM, UserAssignedIdentityProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentityProperties_STATUS_ARM runs a test to see if a specific instance of UserAssignedIdentityProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentityProperties_STATUS_ARM(subject UserAssignedIdentityProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentityProperties_STATUS_ARM
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

// Generator of UserAssignedIdentityProperties_STATUS_ARM instances for property testing - lazily instantiated by
// UserAssignedIdentityProperties_STATUS_ARMGenerator()
var userAssignedIdentityProperties_STATUS_ARMGenerator gopter.Gen

// UserAssignedIdentityProperties_STATUS_ARMGenerator returns a generator of UserAssignedIdentityProperties_STATUS_ARM instances for property testing.
func UserAssignedIdentityProperties_STATUS_ARMGenerator() gopter.Gen {
	if userAssignedIdentityProperties_STATUS_ARMGenerator != nil {
		return userAssignedIdentityProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentityProperties_STATUS_ARM(generators)
	userAssignedIdentityProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentityProperties_STATUS_ARM{}), generators)

	return userAssignedIdentityProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentityProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentityProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
}