// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20200601

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

func Test_PrivateDnsZones_A_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZones_A_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZones_A_STATUS_ARM, PrivateDnsZones_A_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZones_A_STATUS_ARM runs a test to see if a specific instance of PrivateDnsZones_A_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZones_A_STATUS_ARM(subject PrivateDnsZones_A_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZones_A_STATUS_ARM
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

// Generator of PrivateDnsZones_A_STATUS_ARM instances for property testing - lazily instantiated by
// PrivateDnsZones_A_STATUS_ARMGenerator()
var privateDnsZones_A_STATUS_ARMGenerator gopter.Gen

// PrivateDnsZones_A_STATUS_ARMGenerator returns a generator of PrivateDnsZones_A_STATUS_ARM instances for property testing.
// We first initialize privateDnsZones_A_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZones_A_STATUS_ARMGenerator() gopter.Gen {
	if privateDnsZones_A_STATUS_ARMGenerator != nil {
		return privateDnsZones_A_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_A_STATUS_ARM(generators)
	privateDnsZones_A_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_A_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_A_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZones_A_STATUS_ARM(generators)
	privateDnsZones_A_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_A_STATUS_ARM{}), generators)

	return privateDnsZones_A_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZones_A_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZones_A_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZones_A_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZones_A_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RecordSetProperties_STATUS_ARMGenerator())
}