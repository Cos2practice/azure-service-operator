// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501preview

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

func Test_PolicyFragmentContractProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PolicyFragmentContractProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPolicyFragmentContractProperties_STATUS_ARM, PolicyFragmentContractProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPolicyFragmentContractProperties_STATUS_ARM runs a test to see if a specific instance of PolicyFragmentContractProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPolicyFragmentContractProperties_STATUS_ARM(subject PolicyFragmentContractProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PolicyFragmentContractProperties_STATUS_ARM
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

// Generator of PolicyFragmentContractProperties_STATUS_ARM instances for property testing - lazily instantiated by
// PolicyFragmentContractProperties_STATUS_ARMGenerator()
var policyFragmentContractProperties_STATUS_ARMGenerator gopter.Gen

// PolicyFragmentContractProperties_STATUS_ARMGenerator returns a generator of PolicyFragmentContractProperties_STATUS_ARM instances for property testing.
func PolicyFragmentContractProperties_STATUS_ARMGenerator() gopter.Gen {
	if policyFragmentContractProperties_STATUS_ARMGenerator != nil {
		return policyFragmentContractProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPolicyFragmentContractProperties_STATUS_ARM(generators)
	policyFragmentContractProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PolicyFragmentContractProperties_STATUS_ARM{}), generators)

	return policyFragmentContractProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPolicyFragmentContractProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPolicyFragmentContractProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Format"] = gen.PtrOf(gen.OneConstOf(PolicyFragmentContractProperties_Format_STATUS_Rawxml, PolicyFragmentContractProperties_Format_STATUS_Xml))
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_Service_PolicyFragment_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_PolicyFragment_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_PolicyFragment_STATUS_ARM, Service_PolicyFragment_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_PolicyFragment_STATUS_ARM runs a test to see if a specific instance of Service_PolicyFragment_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForService_PolicyFragment_STATUS_ARM(subject Service_PolicyFragment_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_PolicyFragment_STATUS_ARM
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

// Generator of Service_PolicyFragment_STATUS_ARM instances for property testing - lazily instantiated by
// Service_PolicyFragment_STATUS_ARMGenerator()
var service_PolicyFragment_STATUS_ARMGenerator gopter.Gen

// Service_PolicyFragment_STATUS_ARMGenerator returns a generator of Service_PolicyFragment_STATUS_ARM instances for property testing.
// We first initialize service_PolicyFragment_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Service_PolicyFragment_STATUS_ARMGenerator() gopter.Gen {
	if service_PolicyFragment_STATUS_ARMGenerator != nil {
		return service_PolicyFragment_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_PolicyFragment_STATUS_ARM(generators)
	service_PolicyFragment_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Service_PolicyFragment_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_PolicyFragment_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForService_PolicyFragment_STATUS_ARM(generators)
	service_PolicyFragment_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Service_PolicyFragment_STATUS_ARM{}), generators)

	return service_PolicyFragment_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForService_PolicyFragment_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_PolicyFragment_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForService_PolicyFragment_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForService_PolicyFragment_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PolicyFragmentContractProperties_STATUS_ARMGenerator())
}