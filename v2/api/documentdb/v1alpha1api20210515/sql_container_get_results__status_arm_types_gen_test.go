// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

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

func Test_SqlContainerGetResults_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlContainerGetResults_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlContainerGetResultsStatusARM, SqlContainerGetResultsStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlContainerGetResultsStatusARM runs a test to see if a specific instance of SqlContainerGetResults_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlContainerGetResultsStatusARM(subject SqlContainerGetResults_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlContainerGetResults_StatusARM
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

// Generator of SqlContainerGetResults_StatusARM instances for property testing - lazily instantiated by
// SqlContainerGetResultsStatusARMGenerator()
var sqlContainerGetResultsStatusARMGenerator gopter.Gen

// SqlContainerGetResultsStatusARMGenerator returns a generator of SqlContainerGetResults_StatusARM instances for property testing.
// We first initialize sqlContainerGetResultsStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlContainerGetResultsStatusARMGenerator() gopter.Gen {
	if sqlContainerGetResultsStatusARMGenerator != nil {
		return sqlContainerGetResultsStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlContainerGetResultsStatusARM(generators)
	sqlContainerGetResultsStatusARMGenerator = gen.Struct(reflect.TypeOf(SqlContainerGetResults_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlContainerGetResultsStatusARM(generators)
	AddRelatedPropertyGeneratorsForSqlContainerGetResultsStatusARM(generators)
	sqlContainerGetResultsStatusARMGenerator = gen.Struct(reflect.TypeOf(SqlContainerGetResults_StatusARM{}), generators)

	return sqlContainerGetResultsStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlContainerGetResultsStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlContainerGetResultsStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlContainerGetResultsStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlContainerGetResultsStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SqlContainerGetPropertiesStatusARMGenerator())
}

func Test_SqlContainerGetProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlContainerGetProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlContainerGetPropertiesStatusARM, SqlContainerGetPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlContainerGetPropertiesStatusARM runs a test to see if a specific instance of SqlContainerGetProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlContainerGetPropertiesStatusARM(subject SqlContainerGetProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlContainerGetProperties_StatusARM
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

// Generator of SqlContainerGetProperties_StatusARM instances for property testing - lazily instantiated by
// SqlContainerGetPropertiesStatusARMGenerator()
var sqlContainerGetPropertiesStatusARMGenerator gopter.Gen

// SqlContainerGetPropertiesStatusARMGenerator returns a generator of SqlContainerGetProperties_StatusARM instances for property testing.
func SqlContainerGetPropertiesStatusARMGenerator() gopter.Gen {
	if sqlContainerGetPropertiesStatusARMGenerator != nil {
		return sqlContainerGetPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlContainerGetPropertiesStatusARM(generators)
	sqlContainerGetPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(SqlContainerGetProperties_StatusARM{}), generators)

	return sqlContainerGetPropertiesStatusARMGenerator
}

// AddRelatedPropertyGeneratorsForSqlContainerGetPropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlContainerGetPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(OptionsResourceStatusARMGenerator())
	gens["Resource"] = gen.PtrOf(SqlContainerGetPropertiesStatusResourceARMGenerator())
}

func Test_SqlContainerGetProperties_Status_ResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlContainerGetProperties_Status_ResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlContainerGetPropertiesStatusResourceARM, SqlContainerGetPropertiesStatusResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlContainerGetPropertiesStatusResourceARM runs a test to see if a specific instance of SqlContainerGetProperties_Status_ResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlContainerGetPropertiesStatusResourceARM(subject SqlContainerGetProperties_Status_ResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlContainerGetProperties_Status_ResourceARM
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

// Generator of SqlContainerGetProperties_Status_ResourceARM instances for property testing - lazily instantiated by
// SqlContainerGetPropertiesStatusResourceARMGenerator()
var sqlContainerGetPropertiesStatusResourceARMGenerator gopter.Gen

// SqlContainerGetPropertiesStatusResourceARMGenerator returns a generator of SqlContainerGetProperties_Status_ResourceARM instances for property testing.
// We first initialize sqlContainerGetPropertiesStatusResourceARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlContainerGetPropertiesStatusResourceARMGenerator() gopter.Gen {
	if sqlContainerGetPropertiesStatusResourceARMGenerator != nil {
		return sqlContainerGetPropertiesStatusResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlContainerGetPropertiesStatusResourceARM(generators)
	sqlContainerGetPropertiesStatusResourceARMGenerator = gen.Struct(reflect.TypeOf(SqlContainerGetProperties_Status_ResourceARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlContainerGetPropertiesStatusResourceARM(generators)
	AddRelatedPropertyGeneratorsForSqlContainerGetPropertiesStatusResourceARM(generators)
	sqlContainerGetPropertiesStatusResourceARMGenerator = gen.Struct(reflect.TypeOf(SqlContainerGetProperties_Status_ResourceARM{}), generators)

	return sqlContainerGetPropertiesStatusResourceARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlContainerGetPropertiesStatusResourceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlContainerGetPropertiesStatusResourceARM(gens map[string]gopter.Gen) {
	gens["AnalyticalStorageTtl"] = gen.PtrOf(gen.Int())
	gens["DefaultTtl"] = gen.PtrOf(gen.Int())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["Ts"] = gen.PtrOf(gen.Float64())
}

// AddRelatedPropertyGeneratorsForSqlContainerGetPropertiesStatusResourceARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlContainerGetPropertiesStatusResourceARM(gens map[string]gopter.Gen) {
	gens["ConflictResolutionPolicy"] = gen.PtrOf(ConflictResolutionPolicyStatusARMGenerator())
	gens["IndexingPolicy"] = gen.PtrOf(IndexingPolicyStatusARMGenerator())
	gens["PartitionKey"] = gen.PtrOf(ContainerPartitionKeyStatusARMGenerator())
	gens["UniqueKeyPolicy"] = gen.PtrOf(UniqueKeyPolicyStatusARMGenerator())
}

func Test_ConflictResolutionPolicy_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ConflictResolutionPolicy_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConflictResolutionPolicyStatusARM, ConflictResolutionPolicyStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConflictResolutionPolicyStatusARM runs a test to see if a specific instance of ConflictResolutionPolicy_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForConflictResolutionPolicyStatusARM(subject ConflictResolutionPolicy_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ConflictResolutionPolicy_StatusARM
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

// Generator of ConflictResolutionPolicy_StatusARM instances for property testing - lazily instantiated by
// ConflictResolutionPolicyStatusARMGenerator()
var conflictResolutionPolicyStatusARMGenerator gopter.Gen

// ConflictResolutionPolicyStatusARMGenerator returns a generator of ConflictResolutionPolicy_StatusARM instances for property testing.
func ConflictResolutionPolicyStatusARMGenerator() gopter.Gen {
	if conflictResolutionPolicyStatusARMGenerator != nil {
		return conflictResolutionPolicyStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConflictResolutionPolicyStatusARM(generators)
	conflictResolutionPolicyStatusARMGenerator = gen.Struct(reflect.TypeOf(ConflictResolutionPolicy_StatusARM{}), generators)

	return conflictResolutionPolicyStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForConflictResolutionPolicyStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConflictResolutionPolicyStatusARM(gens map[string]gopter.Gen) {
	gens["ConflictResolutionPath"] = gen.PtrOf(gen.AlphaString())
	gens["ConflictResolutionProcedure"] = gen.PtrOf(gen.AlphaString())
	gens["Mode"] = gen.PtrOf(gen.OneConstOf(ConflictResolutionPolicyStatusModeCustom, ConflictResolutionPolicyStatusModeLastWriterWins))
}

func Test_ContainerPartitionKey_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ContainerPartitionKey_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForContainerPartitionKeyStatusARM, ContainerPartitionKeyStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForContainerPartitionKeyStatusARM runs a test to see if a specific instance of ContainerPartitionKey_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForContainerPartitionKeyStatusARM(subject ContainerPartitionKey_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ContainerPartitionKey_StatusARM
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

// Generator of ContainerPartitionKey_StatusARM instances for property testing - lazily instantiated by
// ContainerPartitionKeyStatusARMGenerator()
var containerPartitionKeyStatusARMGenerator gopter.Gen

// ContainerPartitionKeyStatusARMGenerator returns a generator of ContainerPartitionKey_StatusARM instances for property testing.
func ContainerPartitionKeyStatusARMGenerator() gopter.Gen {
	if containerPartitionKeyStatusARMGenerator != nil {
		return containerPartitionKeyStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForContainerPartitionKeyStatusARM(generators)
	containerPartitionKeyStatusARMGenerator = gen.Struct(reflect.TypeOf(ContainerPartitionKey_StatusARM{}), generators)

	return containerPartitionKeyStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForContainerPartitionKeyStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForContainerPartitionKeyStatusARM(gens map[string]gopter.Gen) {
	gens["Kind"] = gen.PtrOf(gen.OneConstOf(ContainerPartitionKeyStatusKindHash, ContainerPartitionKeyStatusKindMultiHash, ContainerPartitionKeyStatusKindRange))
	gens["Paths"] = gen.SliceOf(gen.AlphaString())
	gens["SystemKey"] = gen.PtrOf(gen.Bool())
	gens["Version"] = gen.PtrOf(gen.Int())
}

func Test_IndexingPolicy_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IndexingPolicy_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIndexingPolicyStatusARM, IndexingPolicyStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIndexingPolicyStatusARM runs a test to see if a specific instance of IndexingPolicy_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIndexingPolicyStatusARM(subject IndexingPolicy_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IndexingPolicy_StatusARM
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

// Generator of IndexingPolicy_StatusARM instances for property testing - lazily instantiated by
// IndexingPolicyStatusARMGenerator()
var indexingPolicyStatusARMGenerator gopter.Gen

// IndexingPolicyStatusARMGenerator returns a generator of IndexingPolicy_StatusARM instances for property testing.
// We first initialize indexingPolicyStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func IndexingPolicyStatusARMGenerator() gopter.Gen {
	if indexingPolicyStatusARMGenerator != nil {
		return indexingPolicyStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIndexingPolicyStatusARM(generators)
	indexingPolicyStatusARMGenerator = gen.Struct(reflect.TypeOf(IndexingPolicy_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIndexingPolicyStatusARM(generators)
	AddRelatedPropertyGeneratorsForIndexingPolicyStatusARM(generators)
	indexingPolicyStatusARMGenerator = gen.Struct(reflect.TypeOf(IndexingPolicy_StatusARM{}), generators)

	return indexingPolicyStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForIndexingPolicyStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIndexingPolicyStatusARM(gens map[string]gopter.Gen) {
	gens["Automatic"] = gen.PtrOf(gen.Bool())
	gens["IndexingMode"] = gen.PtrOf(gen.OneConstOf(IndexingPolicyStatusIndexingModeConsistent, IndexingPolicyStatusIndexingModeLazy, IndexingPolicyStatusIndexingModeNone))
}

// AddRelatedPropertyGeneratorsForIndexingPolicyStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIndexingPolicyStatusARM(gens map[string]gopter.Gen) {
	gens["CompositeIndexes"] = gen.SliceOf(gen.SliceOf(CompositePathStatusARMGenerator()))
	gens["ExcludedPaths"] = gen.SliceOf(ExcludedPathStatusARMGenerator())
	gens["IncludedPaths"] = gen.SliceOf(IncludedPathStatusARMGenerator())
	gens["SpatialIndexes"] = gen.SliceOf(SpatialSpecStatusARMGenerator())
}

func Test_UniqueKeyPolicy_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UniqueKeyPolicy_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUniqueKeyPolicyStatusARM, UniqueKeyPolicyStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUniqueKeyPolicyStatusARM runs a test to see if a specific instance of UniqueKeyPolicy_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUniqueKeyPolicyStatusARM(subject UniqueKeyPolicy_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UniqueKeyPolicy_StatusARM
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

// Generator of UniqueKeyPolicy_StatusARM instances for property testing - lazily instantiated by
// UniqueKeyPolicyStatusARMGenerator()
var uniqueKeyPolicyStatusARMGenerator gopter.Gen

// UniqueKeyPolicyStatusARMGenerator returns a generator of UniqueKeyPolicy_StatusARM instances for property testing.
func UniqueKeyPolicyStatusARMGenerator() gopter.Gen {
	if uniqueKeyPolicyStatusARMGenerator != nil {
		return uniqueKeyPolicyStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForUniqueKeyPolicyStatusARM(generators)
	uniqueKeyPolicyStatusARMGenerator = gen.Struct(reflect.TypeOf(UniqueKeyPolicy_StatusARM{}), generators)

	return uniqueKeyPolicyStatusARMGenerator
}

// AddRelatedPropertyGeneratorsForUniqueKeyPolicyStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForUniqueKeyPolicyStatusARM(gens map[string]gopter.Gen) {
	gens["UniqueKeys"] = gen.SliceOf(UniqueKeyStatusARMGenerator())
}

func Test_CompositePath_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CompositePath_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCompositePathStatusARM, CompositePathStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCompositePathStatusARM runs a test to see if a specific instance of CompositePath_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCompositePathStatusARM(subject CompositePath_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CompositePath_StatusARM
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

// Generator of CompositePath_StatusARM instances for property testing - lazily instantiated by
// CompositePathStatusARMGenerator()
var compositePathStatusARMGenerator gopter.Gen

// CompositePathStatusARMGenerator returns a generator of CompositePath_StatusARM instances for property testing.
func CompositePathStatusARMGenerator() gopter.Gen {
	if compositePathStatusARMGenerator != nil {
		return compositePathStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCompositePathStatusARM(generators)
	compositePathStatusARMGenerator = gen.Struct(reflect.TypeOf(CompositePath_StatusARM{}), generators)

	return compositePathStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForCompositePathStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCompositePathStatusARM(gens map[string]gopter.Gen) {
	gens["Order"] = gen.PtrOf(gen.OneConstOf(CompositePathStatusOrderAscending, CompositePathStatusOrderDescending))
	gens["Path"] = gen.PtrOf(gen.AlphaString())
}

func Test_ExcludedPath_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ExcludedPath_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForExcludedPathStatusARM, ExcludedPathStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForExcludedPathStatusARM runs a test to see if a specific instance of ExcludedPath_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForExcludedPathStatusARM(subject ExcludedPath_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ExcludedPath_StatusARM
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

// Generator of ExcludedPath_StatusARM instances for property testing - lazily instantiated by
// ExcludedPathStatusARMGenerator()
var excludedPathStatusARMGenerator gopter.Gen

// ExcludedPathStatusARMGenerator returns a generator of ExcludedPath_StatusARM instances for property testing.
func ExcludedPathStatusARMGenerator() gopter.Gen {
	if excludedPathStatusARMGenerator != nil {
		return excludedPathStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForExcludedPathStatusARM(generators)
	excludedPathStatusARMGenerator = gen.Struct(reflect.TypeOf(ExcludedPath_StatusARM{}), generators)

	return excludedPathStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForExcludedPathStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForExcludedPathStatusARM(gens map[string]gopter.Gen) {
	gens["Path"] = gen.PtrOf(gen.AlphaString())
}

func Test_IncludedPath_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IncludedPath_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIncludedPathStatusARM, IncludedPathStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIncludedPathStatusARM runs a test to see if a specific instance of IncludedPath_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIncludedPathStatusARM(subject IncludedPath_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IncludedPath_StatusARM
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

// Generator of IncludedPath_StatusARM instances for property testing - lazily instantiated by
// IncludedPathStatusARMGenerator()
var includedPathStatusARMGenerator gopter.Gen

// IncludedPathStatusARMGenerator returns a generator of IncludedPath_StatusARM instances for property testing.
// We first initialize includedPathStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func IncludedPathStatusARMGenerator() gopter.Gen {
	if includedPathStatusARMGenerator != nil {
		return includedPathStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIncludedPathStatusARM(generators)
	includedPathStatusARMGenerator = gen.Struct(reflect.TypeOf(IncludedPath_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIncludedPathStatusARM(generators)
	AddRelatedPropertyGeneratorsForIncludedPathStatusARM(generators)
	includedPathStatusARMGenerator = gen.Struct(reflect.TypeOf(IncludedPath_StatusARM{}), generators)

	return includedPathStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForIncludedPathStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIncludedPathStatusARM(gens map[string]gopter.Gen) {
	gens["Path"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForIncludedPathStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIncludedPathStatusARM(gens map[string]gopter.Gen) {
	gens["Indexes"] = gen.SliceOf(IndexesStatusARMGenerator())
}

func Test_SpatialSpec_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SpatialSpec_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSpatialSpecStatusARM, SpatialSpecStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSpatialSpecStatusARM runs a test to see if a specific instance of SpatialSpec_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSpatialSpecStatusARM(subject SpatialSpec_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SpatialSpec_StatusARM
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

// Generator of SpatialSpec_StatusARM instances for property testing - lazily instantiated by
// SpatialSpecStatusARMGenerator()
var spatialSpecStatusARMGenerator gopter.Gen

// SpatialSpecStatusARMGenerator returns a generator of SpatialSpec_StatusARM instances for property testing.
func SpatialSpecStatusARMGenerator() gopter.Gen {
	if spatialSpecStatusARMGenerator != nil {
		return spatialSpecStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSpatialSpecStatusARM(generators)
	spatialSpecStatusARMGenerator = gen.Struct(reflect.TypeOf(SpatialSpec_StatusARM{}), generators)

	return spatialSpecStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSpatialSpecStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSpatialSpecStatusARM(gens map[string]gopter.Gen) {
	gens["Path"] = gen.PtrOf(gen.AlphaString())
	gens["Types"] = gen.SliceOf(gen.OneConstOf(
		SpatialType_StatusLineString,
		SpatialType_StatusMultiPolygon,
		SpatialType_StatusPoint,
		SpatialType_StatusPolygon))
}

func Test_UniqueKey_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UniqueKey_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUniqueKeyStatusARM, UniqueKeyStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUniqueKeyStatusARM runs a test to see if a specific instance of UniqueKey_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUniqueKeyStatusARM(subject UniqueKey_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UniqueKey_StatusARM
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

// Generator of UniqueKey_StatusARM instances for property testing - lazily instantiated by UniqueKeyStatusARMGenerator()
var uniqueKeyStatusARMGenerator gopter.Gen

// UniqueKeyStatusARMGenerator returns a generator of UniqueKey_StatusARM instances for property testing.
func UniqueKeyStatusARMGenerator() gopter.Gen {
	if uniqueKeyStatusARMGenerator != nil {
		return uniqueKeyStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUniqueKeyStatusARM(generators)
	uniqueKeyStatusARMGenerator = gen.Struct(reflect.TypeOf(UniqueKey_StatusARM{}), generators)

	return uniqueKeyStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForUniqueKeyStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUniqueKeyStatusARM(gens map[string]gopter.Gen) {
	gens["Paths"] = gen.SliceOf(gen.AlphaString())
}

func Test_Indexes_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Indexes_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIndexesStatusARM, IndexesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIndexesStatusARM runs a test to see if a specific instance of Indexes_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIndexesStatusARM(subject Indexes_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Indexes_StatusARM
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

// Generator of Indexes_StatusARM instances for property testing - lazily instantiated by IndexesStatusARMGenerator()
var indexesStatusARMGenerator gopter.Gen

// IndexesStatusARMGenerator returns a generator of Indexes_StatusARM instances for property testing.
func IndexesStatusARMGenerator() gopter.Gen {
	if indexesStatusARMGenerator != nil {
		return indexesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIndexesStatusARM(generators)
	indexesStatusARMGenerator = gen.Struct(reflect.TypeOf(Indexes_StatusARM{}), generators)

	return indexesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForIndexesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIndexesStatusARM(gens map[string]gopter.Gen) {
	gens["DataType"] = gen.PtrOf(gen.OneConstOf(
		IndexesStatusDataTypeLineString,
		IndexesStatusDataTypeMultiPolygon,
		IndexesStatusDataTypeNumber,
		IndexesStatusDataTypePoint,
		IndexesStatusDataTypePolygon,
		IndexesStatusDataTypeString))
	gens["Kind"] = gen.PtrOf(gen.OneConstOf(IndexesStatusKindHash, IndexesStatusKindRange, IndexesStatusKindSpatial))
	gens["Precision"] = gen.PtrOf(gen.Int())
}