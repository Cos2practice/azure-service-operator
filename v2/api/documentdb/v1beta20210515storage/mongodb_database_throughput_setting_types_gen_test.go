// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515storage

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

func Test_MongodbDatabaseThroughputSetting_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongodbDatabaseThroughputSetting via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongodbDatabaseThroughputSetting, MongodbDatabaseThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongodbDatabaseThroughputSetting runs a test to see if a specific instance of MongodbDatabaseThroughputSetting round trips to JSON and back losslessly
func RunJSONSerializationTestForMongodbDatabaseThroughputSetting(subject MongodbDatabaseThroughputSetting) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongodbDatabaseThroughputSetting
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

// Generator of MongodbDatabaseThroughputSetting instances for property testing - lazily instantiated by
// MongodbDatabaseThroughputSettingGenerator()
var mongodbDatabaseThroughputSettingGenerator gopter.Gen

// MongodbDatabaseThroughputSettingGenerator returns a generator of MongodbDatabaseThroughputSetting instances for property testing.
func MongodbDatabaseThroughputSettingGenerator() gopter.Gen {
	if mongodbDatabaseThroughputSettingGenerator != nil {
		return mongodbDatabaseThroughputSettingGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongodbDatabaseThroughputSetting(generators)
	mongodbDatabaseThroughputSettingGenerator = gen.Struct(reflect.TypeOf(MongodbDatabaseThroughputSetting{}), generators)

	return mongodbDatabaseThroughputSettingGenerator
}

// AddRelatedPropertyGeneratorsForMongodbDatabaseThroughputSetting is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongodbDatabaseThroughputSetting(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccounts_MongodbDatabases_ThroughputSettings_SpecGenerator()
	gens["Status"] = ThroughputSettingsGetResults_STATUSGenerator()
}

func Test_DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec, DatabaseAccounts_MongodbDatabases_ThroughputSettings_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec runs a test to see if a specific instance of DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec(subject DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec
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

// Generator of DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec instances for property testing - lazily
// instantiated by DatabaseAccounts_MongodbDatabases_ThroughputSettings_SpecGenerator()
var databaseAccounts_MongodbDatabases_ThroughputSettings_SpecGenerator gopter.Gen

// DatabaseAccounts_MongodbDatabases_ThroughputSettings_SpecGenerator returns a generator of DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec instances for property testing.
// We first initialize databaseAccounts_MongodbDatabases_ThroughputSettings_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccounts_MongodbDatabases_ThroughputSettings_SpecGenerator() gopter.Gen {
	if databaseAccounts_MongodbDatabases_ThroughputSettings_SpecGenerator != nil {
		return databaseAccounts_MongodbDatabases_ThroughputSettings_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec(generators)
	databaseAccounts_MongodbDatabases_ThroughputSettings_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec(generators)
	databaseAccounts_MongodbDatabases_ThroughputSettings_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec{}), generators)

	return databaseAccounts_MongodbDatabases_ThroughputSettings_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(ThroughputSettingsResourceGenerator())
}