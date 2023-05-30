// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210301

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

func Test_RedisEnterprise_Database_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisEnterprise_Database_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisEnterprise_Database_Spec_ARM, RedisEnterprise_Database_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisEnterprise_Database_Spec_ARM runs a test to see if a specific instance of RedisEnterprise_Database_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisEnterprise_Database_Spec_ARM(subject RedisEnterprise_Database_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisEnterprise_Database_Spec_ARM
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

// Generator of RedisEnterprise_Database_Spec_ARM instances for property testing - lazily instantiated by
// RedisEnterprise_Database_Spec_ARMGenerator()
var redisEnterprise_Database_Spec_ARMGenerator gopter.Gen

// RedisEnterprise_Database_Spec_ARMGenerator returns a generator of RedisEnterprise_Database_Spec_ARM instances for property testing.
// We first initialize redisEnterprise_Database_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisEnterprise_Database_Spec_ARMGenerator() gopter.Gen {
	if redisEnterprise_Database_Spec_ARMGenerator != nil {
		return redisEnterprise_Database_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterprise_Database_Spec_ARM(generators)
	redisEnterprise_Database_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(RedisEnterprise_Database_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterprise_Database_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForRedisEnterprise_Database_Spec_ARM(generators)
	redisEnterprise_Database_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(RedisEnterprise_Database_Spec_ARM{}), generators)

	return redisEnterprise_Database_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisEnterprise_Database_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisEnterprise_Database_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForRedisEnterprise_Database_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisEnterprise_Database_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DatabaseProperties_ARMGenerator())
}

func Test_DatabaseProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseProperties_ARM, DatabaseProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseProperties_ARM runs a test to see if a specific instance of DatabaseProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseProperties_ARM(subject DatabaseProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseProperties_ARM
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

// Generator of DatabaseProperties_ARM instances for property testing - lazily instantiated by
// DatabaseProperties_ARMGenerator()
var databaseProperties_ARMGenerator gopter.Gen

// DatabaseProperties_ARMGenerator returns a generator of DatabaseProperties_ARM instances for property testing.
// We first initialize databaseProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseProperties_ARMGenerator() gopter.Gen {
	if databaseProperties_ARMGenerator != nil {
		return databaseProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseProperties_ARM(generators)
	databaseProperties_ARMGenerator = gen.Struct(reflect.TypeOf(DatabaseProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseProperties_ARM(generators)
	databaseProperties_ARMGenerator = gen.Struct(reflect.TypeOf(DatabaseProperties_ARM{}), generators)

	return databaseProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseProperties_ARM(gens map[string]gopter.Gen) {
	gens["ClientProtocol"] = gen.PtrOf(gen.OneConstOf(DatabaseProperties_ClientProtocol_Encrypted, DatabaseProperties_ClientProtocol_Plaintext))
	gens["ClusteringPolicy"] = gen.PtrOf(gen.OneConstOf(DatabaseProperties_ClusteringPolicy_EnterpriseCluster, DatabaseProperties_ClusteringPolicy_OSSCluster))
	gens["EvictionPolicy"] = gen.PtrOf(gen.OneConstOf(
		DatabaseProperties_EvictionPolicy_AllKeysLFU,
		DatabaseProperties_EvictionPolicy_AllKeysLRU,
		DatabaseProperties_EvictionPolicy_AllKeysRandom,
		DatabaseProperties_EvictionPolicy_NoEviction,
		DatabaseProperties_EvictionPolicy_VolatileLFU,
		DatabaseProperties_EvictionPolicy_VolatileLRU,
		DatabaseProperties_EvictionPolicy_VolatileRandom,
		DatabaseProperties_EvictionPolicy_VolatileTTL))
	gens["Port"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForDatabaseProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseProperties_ARM(gens map[string]gopter.Gen) {
	gens["Modules"] = gen.SliceOf(Module_ARMGenerator())
	gens["Persistence"] = gen.PtrOf(Persistence_ARMGenerator())
}

func Test_Module_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Module_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForModule_ARM, Module_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForModule_ARM runs a test to see if a specific instance of Module_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForModule_ARM(subject Module_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Module_ARM
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

// Generator of Module_ARM instances for property testing - lazily instantiated by Module_ARMGenerator()
var module_ARMGenerator gopter.Gen

// Module_ARMGenerator returns a generator of Module_ARM instances for property testing.
func Module_ARMGenerator() gopter.Gen {
	if module_ARMGenerator != nil {
		return module_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForModule_ARM(generators)
	module_ARMGenerator = gen.Struct(reflect.TypeOf(Module_ARM{}), generators)

	return module_ARMGenerator
}

// AddIndependentPropertyGeneratorsForModule_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForModule_ARM(gens map[string]gopter.Gen) {
	gens["Args"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

func Test_Persistence_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Persistence_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPersistence_ARM, Persistence_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPersistence_ARM runs a test to see if a specific instance of Persistence_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPersistence_ARM(subject Persistence_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Persistence_ARM
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

// Generator of Persistence_ARM instances for property testing - lazily instantiated by Persistence_ARMGenerator()
var persistence_ARMGenerator gopter.Gen

// Persistence_ARMGenerator returns a generator of Persistence_ARM instances for property testing.
func Persistence_ARMGenerator() gopter.Gen {
	if persistence_ARMGenerator != nil {
		return persistence_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPersistence_ARM(generators)
	persistence_ARMGenerator = gen.Struct(reflect.TypeOf(Persistence_ARM{}), generators)

	return persistence_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPersistence_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPersistence_ARM(gens map[string]gopter.Gen) {
	gens["AofEnabled"] = gen.PtrOf(gen.Bool())
	gens["AofFrequency"] = gen.PtrOf(gen.OneConstOf(Persistence_AofFrequency_1S, Persistence_AofFrequency_Always))
	gens["RdbEnabled"] = gen.PtrOf(gen.Bool())
	gens["RdbFrequency"] = gen.PtrOf(gen.OneConstOf(Persistence_RdbFrequency_12H, Persistence_RdbFrequency_1H, Persistence_RdbFrequency_6H))
}