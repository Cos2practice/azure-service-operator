// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20201201

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

func Test_VirtualMachineExtensionInstanceView_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachineExtensionInstanceView_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachineExtensionInstanceView_STATUS_ARM, VirtualMachineExtensionInstanceView_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachineExtensionInstanceView_STATUS_ARM runs a test to see if a specific instance of VirtualMachineExtensionInstanceView_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachineExtensionInstanceView_STATUS_ARM(subject VirtualMachineExtensionInstanceView_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachineExtensionInstanceView_STATUS_ARM
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

// Generator of VirtualMachineExtensionInstanceView_STATUS_ARM instances for property testing - lazily instantiated by
// VirtualMachineExtensionInstanceView_STATUS_ARMGenerator()
var virtualMachineExtensionInstanceView_STATUS_ARMGenerator gopter.Gen

// VirtualMachineExtensionInstanceView_STATUS_ARMGenerator returns a generator of VirtualMachineExtensionInstanceView_STATUS_ARM instances for property testing.
// We first initialize virtualMachineExtensionInstanceView_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualMachineExtensionInstanceView_STATUS_ARMGenerator() gopter.Gen {
	if virtualMachineExtensionInstanceView_STATUS_ARMGenerator != nil {
		return virtualMachineExtensionInstanceView_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineExtensionInstanceView_STATUS_ARM(generators)
	virtualMachineExtensionInstanceView_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualMachineExtensionInstanceView_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineExtensionInstanceView_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForVirtualMachineExtensionInstanceView_STATUS_ARM(generators)
	virtualMachineExtensionInstanceView_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualMachineExtensionInstanceView_STATUS_ARM{}), generators)

	return virtualMachineExtensionInstanceView_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualMachineExtensionInstanceView_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualMachineExtensionInstanceView_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["TypeHandlerVersion"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualMachineExtensionInstanceView_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualMachineExtensionInstanceView_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Statuses"] = gen.SliceOf(InstanceViewStatus_STATUS_ARMGenerator())
	gens["Substatuses"] = gen.SliceOf(InstanceViewStatus_STATUS_ARMGenerator())
}

func Test_VirtualMachineExtensionProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachineExtensionProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachineExtensionProperties_STATUS_ARM, VirtualMachineExtensionProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachineExtensionProperties_STATUS_ARM runs a test to see if a specific instance of VirtualMachineExtensionProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachineExtensionProperties_STATUS_ARM(subject VirtualMachineExtensionProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachineExtensionProperties_STATUS_ARM
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

// Generator of VirtualMachineExtensionProperties_STATUS_ARM instances for property testing - lazily instantiated by
// VirtualMachineExtensionProperties_STATUS_ARMGenerator()
var virtualMachineExtensionProperties_STATUS_ARMGenerator gopter.Gen

// VirtualMachineExtensionProperties_STATUS_ARMGenerator returns a generator of VirtualMachineExtensionProperties_STATUS_ARM instances for property testing.
// We first initialize virtualMachineExtensionProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualMachineExtensionProperties_STATUS_ARMGenerator() gopter.Gen {
	if virtualMachineExtensionProperties_STATUS_ARMGenerator != nil {
		return virtualMachineExtensionProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineExtensionProperties_STATUS_ARM(generators)
	virtualMachineExtensionProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualMachineExtensionProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineExtensionProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForVirtualMachineExtensionProperties_STATUS_ARM(generators)
	virtualMachineExtensionProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualMachineExtensionProperties_STATUS_ARM{}), generators)

	return virtualMachineExtensionProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualMachineExtensionProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualMachineExtensionProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AutoUpgradeMinorVersion"] = gen.PtrOf(gen.Bool())
	gens["EnableAutomaticUpgrade"] = gen.PtrOf(gen.Bool())
	gens["ForceUpdateTag"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["Publisher"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["TypeHandlerVersion"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualMachineExtensionProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualMachineExtensionProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["InstanceView"] = gen.PtrOf(VirtualMachineExtensionInstanceView_STATUS_ARMGenerator())
}

func Test_VirtualMachines_Extension_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachines_Extension_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachines_Extension_STATUS_ARM, VirtualMachines_Extension_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachines_Extension_STATUS_ARM runs a test to see if a specific instance of VirtualMachines_Extension_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachines_Extension_STATUS_ARM(subject VirtualMachines_Extension_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachines_Extension_STATUS_ARM
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

// Generator of VirtualMachines_Extension_STATUS_ARM instances for property testing - lazily instantiated by
// VirtualMachines_Extension_STATUS_ARMGenerator()
var virtualMachines_Extension_STATUS_ARMGenerator gopter.Gen

// VirtualMachines_Extension_STATUS_ARMGenerator returns a generator of VirtualMachines_Extension_STATUS_ARM instances for property testing.
// We first initialize virtualMachines_Extension_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualMachines_Extension_STATUS_ARMGenerator() gopter.Gen {
	if virtualMachines_Extension_STATUS_ARMGenerator != nil {
		return virtualMachines_Extension_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachines_Extension_STATUS_ARM(generators)
	virtualMachines_Extension_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualMachines_Extension_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachines_Extension_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForVirtualMachines_Extension_STATUS_ARM(generators)
	virtualMachines_Extension_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualMachines_Extension_STATUS_ARM{}), generators)

	return virtualMachines_Extension_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualMachines_Extension_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualMachines_Extension_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualMachines_Extension_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualMachines_Extension_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(VirtualMachineExtensionProperties_STATUS_ARMGenerator())
}