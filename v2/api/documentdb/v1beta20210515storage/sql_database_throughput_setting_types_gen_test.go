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

func Test_SqlDatabaseThroughputSetting_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseThroughputSetting via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseThroughputSetting, SqlDatabaseThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseThroughputSetting runs a test to see if a specific instance of SqlDatabaseThroughputSetting round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseThroughputSetting(subject SqlDatabaseThroughputSetting) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseThroughputSetting
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

// Generator of SqlDatabaseThroughputSetting instances for property testing - lazily instantiated by
// SqlDatabaseThroughputSettingGenerator()
var sqlDatabaseThroughputSettingGenerator gopter.Gen

// SqlDatabaseThroughputSettingGenerator returns a generator of SqlDatabaseThroughputSetting instances for property testing.
func SqlDatabaseThroughputSettingGenerator() gopter.Gen {
	if sqlDatabaseThroughputSettingGenerator != nil {
		return sqlDatabaseThroughputSettingGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlDatabaseThroughputSetting(generators)
	sqlDatabaseThroughputSettingGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseThroughputSetting{}), generators)

	return sqlDatabaseThroughputSettingGenerator
}

// AddRelatedPropertyGeneratorsForSqlDatabaseThroughputSetting is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseThroughputSetting(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccounts_SqlDatabases_ThroughputSetting_SpecGenerator()
	gens["Status"] = ThroughputSettingsGetResults_STATUSGenerator()
}

func Test_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec, DatabaseAccounts_SqlDatabases_ThroughputSetting_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec runs a test to see if a specific instance of DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(subject DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec
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

// Generator of DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec instances for property testing - lazily
// instantiated by DatabaseAccounts_SqlDatabases_ThroughputSetting_SpecGenerator()
var databaseAccounts_SqlDatabases_ThroughputSetting_SpecGenerator gopter.Gen

// DatabaseAccounts_SqlDatabases_ThroughputSetting_SpecGenerator returns a generator of DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec instances for property testing.
// We first initialize databaseAccounts_SqlDatabases_ThroughputSetting_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccounts_SqlDatabases_ThroughputSetting_SpecGenerator() gopter.Gen {
	if databaseAccounts_SqlDatabases_ThroughputSetting_SpecGenerator != nil {
		return databaseAccounts_SqlDatabases_ThroughputSetting_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(generators)
	databaseAccounts_SqlDatabases_ThroughputSetting_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(generators)
	databaseAccounts_SqlDatabases_ThroughputSetting_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec{}), generators)

	return databaseAccounts_SqlDatabases_ThroughputSetting_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(ThroughputSettingsResourceGenerator())
}