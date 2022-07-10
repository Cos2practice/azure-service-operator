// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

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

func Test_StorageAccountsQueueServicesQueues_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsQueueServicesQueues_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsQueueServicesQueuesSpecARM, StorageAccountsQueueServicesQueuesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsQueueServicesQueuesSpecARM runs a test to see if a specific instance of StorageAccountsQueueServicesQueues_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsQueueServicesQueuesSpecARM(subject StorageAccountsQueueServicesQueues_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsQueueServicesQueues_SpecARM
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

// Generator of StorageAccountsQueueServicesQueues_SpecARM instances for property testing - lazily instantiated by
// StorageAccountsQueueServicesQueuesSpecARMGenerator()
var storageAccountsQueueServicesQueuesSpecARMGenerator gopter.Gen

// StorageAccountsQueueServicesQueuesSpecARMGenerator returns a generator of StorageAccountsQueueServicesQueues_SpecARM instances for property testing.
// We first initialize storageAccountsQueueServicesQueuesSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsQueueServicesQueuesSpecARMGenerator() gopter.Gen {
	if storageAccountsQueueServicesQueuesSpecARMGenerator != nil {
		return storageAccountsQueueServicesQueuesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesQueuesSpecARM(generators)
	storageAccountsQueueServicesQueuesSpecARMGenerator = gen.Struct(reflect.TypeOf(StorageAccountsQueueServicesQueues_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesQueuesSpecARM(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsQueueServicesQueuesSpecARM(generators)
	storageAccountsQueueServicesQueuesSpecARMGenerator = gen.Struct(reflect.TypeOf(StorageAccountsQueueServicesQueues_SpecARM{}), generators)

	return storageAccountsQueueServicesQueuesSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesQueuesSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesQueuesSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageAccountsQueueServicesQueuesSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsQueueServicesQueuesSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(QueuePropertiesARMGenerator())
}

func Test_QueuePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of QueuePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForQueuePropertiesARM, QueuePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForQueuePropertiesARM runs a test to see if a specific instance of QueuePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForQueuePropertiesARM(subject QueuePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual QueuePropertiesARM
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

// Generator of QueuePropertiesARM instances for property testing - lazily instantiated by QueuePropertiesARMGenerator()
var queuePropertiesARMGenerator gopter.Gen

// QueuePropertiesARMGenerator returns a generator of QueuePropertiesARM instances for property testing.
func QueuePropertiesARMGenerator() gopter.Gen {
	if queuePropertiesARMGenerator != nil {
		return queuePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForQueuePropertiesARM(generators)
	queuePropertiesARMGenerator = gen.Struct(reflect.TypeOf(QueuePropertiesARM{}), generators)

	return queuePropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForQueuePropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForQueuePropertiesARM(gens map[string]gopter.Gen) {
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}
