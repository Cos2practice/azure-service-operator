// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101preview

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

func Test_SBTopic_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SBTopic_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBTopicStatusARM, SBTopicStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBTopicStatusARM runs a test to see if a specific instance of SBTopic_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSBTopicStatusARM(subject SBTopic_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SBTopic_StatusARM
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

// Generator of SBTopic_StatusARM instances for property testing - lazily instantiated by SBTopicStatusARMGenerator()
var sbTopicStatusARMGenerator gopter.Gen

// SBTopicStatusARMGenerator returns a generator of SBTopic_StatusARM instances for property testing.
// We first initialize sbTopicStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SBTopicStatusARMGenerator() gopter.Gen {
	if sbTopicStatusARMGenerator != nil {
		return sbTopicStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBTopicStatusARM(generators)
	sbTopicStatusARMGenerator = gen.Struct(reflect.TypeOf(SBTopic_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBTopicStatusARM(generators)
	AddRelatedPropertyGeneratorsForSBTopicStatusARM(generators)
	sbTopicStatusARMGenerator = gen.Struct(reflect.TypeOf(SBTopic_StatusARM{}), generators)

	return sbTopicStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSBTopicStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBTopicStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSBTopicStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSBTopicStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SBTopicPropertiesStatusARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataStatusARMGenerator())
}

func Test_SBTopicProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SBTopicProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBTopicPropertiesStatusARM, SBTopicPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBTopicPropertiesStatusARM runs a test to see if a specific instance of SBTopicProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSBTopicPropertiesStatusARM(subject SBTopicProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SBTopicProperties_StatusARM
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

// Generator of SBTopicProperties_StatusARM instances for property testing - lazily instantiated by
// SBTopicPropertiesStatusARMGenerator()
var sbTopicPropertiesStatusARMGenerator gopter.Gen

// SBTopicPropertiesStatusARMGenerator returns a generator of SBTopicProperties_StatusARM instances for property testing.
// We first initialize sbTopicPropertiesStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SBTopicPropertiesStatusARMGenerator() gopter.Gen {
	if sbTopicPropertiesStatusARMGenerator != nil {
		return sbTopicPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBTopicPropertiesStatusARM(generators)
	sbTopicPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(SBTopicProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBTopicPropertiesStatusARM(generators)
	AddRelatedPropertyGeneratorsForSBTopicPropertiesStatusARM(generators)
	sbTopicPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(SBTopicProperties_StatusARM{}), generators)

	return sbTopicPropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSBTopicPropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBTopicPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["AccessedAt"] = gen.PtrOf(gen.AlphaString())
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["DefaultMessageTimeToLive"] = gen.PtrOf(gen.AlphaString())
	gens["DuplicateDetectionHistoryTimeWindow"] = gen.PtrOf(gen.AlphaString())
	gens["EnableBatchedOperations"] = gen.PtrOf(gen.Bool())
	gens["EnableExpress"] = gen.PtrOf(gen.Bool())
	gens["EnablePartitioning"] = gen.PtrOf(gen.Bool())
	gens["MaxSizeInMegabytes"] = gen.PtrOf(gen.Int())
	gens["RequiresDuplicateDetection"] = gen.PtrOf(gen.Bool())
	gens["SizeInBytes"] = gen.PtrOf(gen.Int())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(
		EntityStatus_StatusActive,
		EntityStatus_StatusCreating,
		EntityStatus_StatusDeleting,
		EntityStatus_StatusDisabled,
		EntityStatus_StatusReceiveDisabled,
		EntityStatus_StatusRenaming,
		EntityStatus_StatusRestoring,
		EntityStatus_StatusSendDisabled,
		EntityStatus_StatusUnknown))
	gens["SubscriptionCount"] = gen.PtrOf(gen.Int())
	gens["SupportOrdering"] = gen.PtrOf(gen.Bool())
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSBTopicPropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSBTopicPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["CountDetails"] = gen.PtrOf(MessageCountDetailsStatusARMGenerator())
}