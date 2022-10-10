// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401

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

func Test_StorageQueue_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageQueue_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageQueue_STATUS_ARM, StorageQueue_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageQueue_STATUS_ARM runs a test to see if a specific instance of StorageQueue_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageQueue_STATUS_ARM(subject StorageQueue_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageQueue_STATUS_ARM
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

// Generator of StorageQueue_STATUS_ARM instances for property testing - lazily instantiated by
// StorageQueue_STATUS_ARMGenerator()
var storageQueue_STATUS_ARMGenerator gopter.Gen

// StorageQueue_STATUS_ARMGenerator returns a generator of StorageQueue_STATUS_ARM instances for property testing.
// We first initialize storageQueue_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageQueue_STATUS_ARMGenerator() gopter.Gen {
	if storageQueue_STATUS_ARMGenerator != nil {
		return storageQueue_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageQueue_STATUS_ARM(generators)
	storageQueue_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(StorageQueue_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageQueue_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForStorageQueue_STATUS_ARM(generators)
	storageQueue_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(StorageQueue_STATUS_ARM{}), generators)

	return storageQueue_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForStorageQueue_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageQueue_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageQueue_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageQueue_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(QueueProperties_STATUS_ARMGenerator())
}

func Test_QueueProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of QueueProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForQueueProperties_STATUS_ARM, QueueProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForQueueProperties_STATUS_ARM runs a test to see if a specific instance of QueueProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForQueueProperties_STATUS_ARM(subject QueueProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual QueueProperties_STATUS_ARM
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

// Generator of QueueProperties_STATUS_ARM instances for property testing - lazily instantiated by
// QueueProperties_STATUS_ARMGenerator()
var queueProperties_STATUS_ARMGenerator gopter.Gen

// QueueProperties_STATUS_ARMGenerator returns a generator of QueueProperties_STATUS_ARM instances for property testing.
func QueueProperties_STATUS_ARMGenerator() gopter.Gen {
	if queueProperties_STATUS_ARMGenerator != nil {
		return queueProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForQueueProperties_STATUS_ARM(generators)
	queueProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(QueueProperties_STATUS_ARM{}), generators)

	return queueProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForQueueProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForQueueProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ApproximateMessageCount"] = gen.PtrOf(gen.Int())
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}