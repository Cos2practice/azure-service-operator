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

func Test_ManagementPolicy_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicy_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicyStatusARM, ManagementPolicyStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicyStatusARM runs a test to see if a specific instance of ManagementPolicy_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicyStatusARM(subject ManagementPolicy_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicy_StatusARM
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

// Generator of ManagementPolicy_StatusARM instances for property testing - lazily instantiated by
// ManagementPolicyStatusARMGenerator()
var managementPolicyStatusARMGenerator gopter.Gen

// ManagementPolicyStatusARMGenerator returns a generator of ManagementPolicy_StatusARM instances for property testing.
// We first initialize managementPolicyStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagementPolicyStatusARMGenerator() gopter.Gen {
	if managementPolicyStatusARMGenerator != nil {
		return managementPolicyStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagementPolicyStatusARM(generators)
	managementPolicyStatusARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicy_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagementPolicyStatusARM(generators)
	AddRelatedPropertyGeneratorsForManagementPolicyStatusARM(generators)
	managementPolicyStatusARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicy_StatusARM{}), generators)

	return managementPolicyStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForManagementPolicyStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagementPolicyStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForManagementPolicyStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicyStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ManagementPolicyPropertiesStatusARMGenerator())
}

func Test_ManagementPolicyProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicyProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicyPropertiesStatusARM, ManagementPolicyPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicyPropertiesStatusARM runs a test to see if a specific instance of ManagementPolicyProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicyPropertiesStatusARM(subject ManagementPolicyProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicyProperties_StatusARM
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

// Generator of ManagementPolicyProperties_StatusARM instances for property testing - lazily instantiated by
// ManagementPolicyPropertiesStatusARMGenerator()
var managementPolicyPropertiesStatusARMGenerator gopter.Gen

// ManagementPolicyPropertiesStatusARMGenerator returns a generator of ManagementPolicyProperties_StatusARM instances for property testing.
// We first initialize managementPolicyPropertiesStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagementPolicyPropertiesStatusARMGenerator() gopter.Gen {
	if managementPolicyPropertiesStatusARMGenerator != nil {
		return managementPolicyPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagementPolicyPropertiesStatusARM(generators)
	managementPolicyPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagementPolicyPropertiesStatusARM(generators)
	AddRelatedPropertyGeneratorsForManagementPolicyPropertiesStatusARM(generators)
	managementPolicyPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyProperties_StatusARM{}), generators)

	return managementPolicyPropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForManagementPolicyPropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagementPolicyPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["LastModifiedTime"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForManagementPolicyPropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicyPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["Policy"] = gen.PtrOf(ManagementPolicySchemaStatusARMGenerator())
}

func Test_ManagementPolicySchema_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicySchema_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicySchemaStatusARM, ManagementPolicySchemaStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicySchemaStatusARM runs a test to see if a specific instance of ManagementPolicySchema_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicySchemaStatusARM(subject ManagementPolicySchema_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicySchema_StatusARM
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

// Generator of ManagementPolicySchema_StatusARM instances for property testing - lazily instantiated by
// ManagementPolicySchemaStatusARMGenerator()
var managementPolicySchemaStatusARMGenerator gopter.Gen

// ManagementPolicySchemaStatusARMGenerator returns a generator of ManagementPolicySchema_StatusARM instances for property testing.
func ManagementPolicySchemaStatusARMGenerator() gopter.Gen {
	if managementPolicySchemaStatusARMGenerator != nil {
		return managementPolicySchemaStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForManagementPolicySchemaStatusARM(generators)
	managementPolicySchemaStatusARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicySchema_StatusARM{}), generators)

	return managementPolicySchemaStatusARMGenerator
}

// AddRelatedPropertyGeneratorsForManagementPolicySchemaStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicySchemaStatusARM(gens map[string]gopter.Gen) {
	gens["Rules"] = gen.SliceOf(ManagementPolicyRuleStatusARMGenerator())
}

func Test_ManagementPolicyRule_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicyRule_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicyRuleStatusARM, ManagementPolicyRuleStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicyRuleStatusARM runs a test to see if a specific instance of ManagementPolicyRule_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicyRuleStatusARM(subject ManagementPolicyRule_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicyRule_StatusARM
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

// Generator of ManagementPolicyRule_StatusARM instances for property testing - lazily instantiated by
// ManagementPolicyRuleStatusARMGenerator()
var managementPolicyRuleStatusARMGenerator gopter.Gen

// ManagementPolicyRuleStatusARMGenerator returns a generator of ManagementPolicyRule_StatusARM instances for property testing.
// We first initialize managementPolicyRuleStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagementPolicyRuleStatusARMGenerator() gopter.Gen {
	if managementPolicyRuleStatusARMGenerator != nil {
		return managementPolicyRuleStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagementPolicyRuleStatusARM(generators)
	managementPolicyRuleStatusARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyRule_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagementPolicyRuleStatusARM(generators)
	AddRelatedPropertyGeneratorsForManagementPolicyRuleStatusARM(generators)
	managementPolicyRuleStatusARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyRule_StatusARM{}), generators)

	return managementPolicyRuleStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForManagementPolicyRuleStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagementPolicyRuleStatusARM(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(ManagementPolicyRuleStatusTypeLifecycle))
}

// AddRelatedPropertyGeneratorsForManagementPolicyRuleStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicyRuleStatusARM(gens map[string]gopter.Gen) {
	gens["Definition"] = gen.PtrOf(ManagementPolicyDefinitionStatusARMGenerator())
}

func Test_ManagementPolicyDefinition_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicyDefinition_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicyDefinitionStatusARM, ManagementPolicyDefinitionStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicyDefinitionStatusARM runs a test to see if a specific instance of ManagementPolicyDefinition_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicyDefinitionStatusARM(subject ManagementPolicyDefinition_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicyDefinition_StatusARM
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

// Generator of ManagementPolicyDefinition_StatusARM instances for property testing - lazily instantiated by
// ManagementPolicyDefinitionStatusARMGenerator()
var managementPolicyDefinitionStatusARMGenerator gopter.Gen

// ManagementPolicyDefinitionStatusARMGenerator returns a generator of ManagementPolicyDefinition_StatusARM instances for property testing.
func ManagementPolicyDefinitionStatusARMGenerator() gopter.Gen {
	if managementPolicyDefinitionStatusARMGenerator != nil {
		return managementPolicyDefinitionStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForManagementPolicyDefinitionStatusARM(generators)
	managementPolicyDefinitionStatusARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyDefinition_StatusARM{}), generators)

	return managementPolicyDefinitionStatusARMGenerator
}

// AddRelatedPropertyGeneratorsForManagementPolicyDefinitionStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicyDefinitionStatusARM(gens map[string]gopter.Gen) {
	gens["Actions"] = gen.PtrOf(ManagementPolicyActionStatusARMGenerator())
	gens["Filters"] = gen.PtrOf(ManagementPolicyFilterStatusARMGenerator())
}

func Test_ManagementPolicyAction_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicyAction_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicyActionStatusARM, ManagementPolicyActionStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicyActionStatusARM runs a test to see if a specific instance of ManagementPolicyAction_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicyActionStatusARM(subject ManagementPolicyAction_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicyAction_StatusARM
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

// Generator of ManagementPolicyAction_StatusARM instances for property testing - lazily instantiated by
// ManagementPolicyActionStatusARMGenerator()
var managementPolicyActionStatusARMGenerator gopter.Gen

// ManagementPolicyActionStatusARMGenerator returns a generator of ManagementPolicyAction_StatusARM instances for property testing.
func ManagementPolicyActionStatusARMGenerator() gopter.Gen {
	if managementPolicyActionStatusARMGenerator != nil {
		return managementPolicyActionStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForManagementPolicyActionStatusARM(generators)
	managementPolicyActionStatusARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyAction_StatusARM{}), generators)

	return managementPolicyActionStatusARMGenerator
}

// AddRelatedPropertyGeneratorsForManagementPolicyActionStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicyActionStatusARM(gens map[string]gopter.Gen) {
	gens["BaseBlob"] = gen.PtrOf(ManagementPolicyBaseBlobStatusARMGenerator())
	gens["Snapshot"] = gen.PtrOf(ManagementPolicySnapShotStatusARMGenerator())
	gens["Version"] = gen.PtrOf(ManagementPolicyVersionStatusARMGenerator())
}

func Test_ManagementPolicyFilter_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicyFilter_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicyFilterStatusARM, ManagementPolicyFilterStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicyFilterStatusARM runs a test to see if a specific instance of ManagementPolicyFilter_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicyFilterStatusARM(subject ManagementPolicyFilter_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicyFilter_StatusARM
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

// Generator of ManagementPolicyFilter_StatusARM instances for property testing - lazily instantiated by
// ManagementPolicyFilterStatusARMGenerator()
var managementPolicyFilterStatusARMGenerator gopter.Gen

// ManagementPolicyFilterStatusARMGenerator returns a generator of ManagementPolicyFilter_StatusARM instances for property testing.
// We first initialize managementPolicyFilterStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagementPolicyFilterStatusARMGenerator() gopter.Gen {
	if managementPolicyFilterStatusARMGenerator != nil {
		return managementPolicyFilterStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagementPolicyFilterStatusARM(generators)
	managementPolicyFilterStatusARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyFilter_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagementPolicyFilterStatusARM(generators)
	AddRelatedPropertyGeneratorsForManagementPolicyFilterStatusARM(generators)
	managementPolicyFilterStatusARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyFilter_StatusARM{}), generators)

	return managementPolicyFilterStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForManagementPolicyFilterStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagementPolicyFilterStatusARM(gens map[string]gopter.Gen) {
	gens["BlobTypes"] = gen.SliceOf(gen.AlphaString())
	gens["PrefixMatch"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForManagementPolicyFilterStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicyFilterStatusARM(gens map[string]gopter.Gen) {
	gens["BlobIndexMatch"] = gen.SliceOf(TagFilterStatusARMGenerator())
}

func Test_ManagementPolicyBaseBlob_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicyBaseBlob_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicyBaseBlobStatusARM, ManagementPolicyBaseBlobStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicyBaseBlobStatusARM runs a test to see if a specific instance of ManagementPolicyBaseBlob_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicyBaseBlobStatusARM(subject ManagementPolicyBaseBlob_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicyBaseBlob_StatusARM
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

// Generator of ManagementPolicyBaseBlob_StatusARM instances for property testing - lazily instantiated by
// ManagementPolicyBaseBlobStatusARMGenerator()
var managementPolicyBaseBlobStatusARMGenerator gopter.Gen

// ManagementPolicyBaseBlobStatusARMGenerator returns a generator of ManagementPolicyBaseBlob_StatusARM instances for property testing.
// We first initialize managementPolicyBaseBlobStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagementPolicyBaseBlobStatusARMGenerator() gopter.Gen {
	if managementPolicyBaseBlobStatusARMGenerator != nil {
		return managementPolicyBaseBlobStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagementPolicyBaseBlobStatusARM(generators)
	managementPolicyBaseBlobStatusARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyBaseBlob_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagementPolicyBaseBlobStatusARM(generators)
	AddRelatedPropertyGeneratorsForManagementPolicyBaseBlobStatusARM(generators)
	managementPolicyBaseBlobStatusARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyBaseBlob_StatusARM{}), generators)

	return managementPolicyBaseBlobStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForManagementPolicyBaseBlobStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagementPolicyBaseBlobStatusARM(gens map[string]gopter.Gen) {
	gens["EnableAutoTierToHotFromCool"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForManagementPolicyBaseBlobStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicyBaseBlobStatusARM(gens map[string]gopter.Gen) {
	gens["Delete"] = gen.PtrOf(DateAfterModificationStatusARMGenerator())
	gens["TierToArchive"] = gen.PtrOf(DateAfterModificationStatusARMGenerator())
	gens["TierToCool"] = gen.PtrOf(DateAfterModificationStatusARMGenerator())
}

func Test_ManagementPolicySnapShot_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicySnapShot_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicySnapShotStatusARM, ManagementPolicySnapShotStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicySnapShotStatusARM runs a test to see if a specific instance of ManagementPolicySnapShot_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicySnapShotStatusARM(subject ManagementPolicySnapShot_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicySnapShot_StatusARM
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

// Generator of ManagementPolicySnapShot_StatusARM instances for property testing - lazily instantiated by
// ManagementPolicySnapShotStatusARMGenerator()
var managementPolicySnapShotStatusARMGenerator gopter.Gen

// ManagementPolicySnapShotStatusARMGenerator returns a generator of ManagementPolicySnapShot_StatusARM instances for property testing.
func ManagementPolicySnapShotStatusARMGenerator() gopter.Gen {
	if managementPolicySnapShotStatusARMGenerator != nil {
		return managementPolicySnapShotStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForManagementPolicySnapShotStatusARM(generators)
	managementPolicySnapShotStatusARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicySnapShot_StatusARM{}), generators)

	return managementPolicySnapShotStatusARMGenerator
}

// AddRelatedPropertyGeneratorsForManagementPolicySnapShotStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicySnapShotStatusARM(gens map[string]gopter.Gen) {
	gens["Delete"] = gen.PtrOf(DateAfterCreationStatusARMGenerator())
	gens["TierToArchive"] = gen.PtrOf(DateAfterCreationStatusARMGenerator())
	gens["TierToCool"] = gen.PtrOf(DateAfterCreationStatusARMGenerator())
}

func Test_ManagementPolicyVersion_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicyVersion_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicyVersionStatusARM, ManagementPolicyVersionStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicyVersionStatusARM runs a test to see if a specific instance of ManagementPolicyVersion_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicyVersionStatusARM(subject ManagementPolicyVersion_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicyVersion_StatusARM
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

// Generator of ManagementPolicyVersion_StatusARM instances for property testing - lazily instantiated by
// ManagementPolicyVersionStatusARMGenerator()
var managementPolicyVersionStatusARMGenerator gopter.Gen

// ManagementPolicyVersionStatusARMGenerator returns a generator of ManagementPolicyVersion_StatusARM instances for property testing.
func ManagementPolicyVersionStatusARMGenerator() gopter.Gen {
	if managementPolicyVersionStatusARMGenerator != nil {
		return managementPolicyVersionStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForManagementPolicyVersionStatusARM(generators)
	managementPolicyVersionStatusARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyVersion_StatusARM{}), generators)

	return managementPolicyVersionStatusARMGenerator
}

// AddRelatedPropertyGeneratorsForManagementPolicyVersionStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicyVersionStatusARM(gens map[string]gopter.Gen) {
	gens["Delete"] = gen.PtrOf(DateAfterCreationStatusARMGenerator())
	gens["TierToArchive"] = gen.PtrOf(DateAfterCreationStatusARMGenerator())
	gens["TierToCool"] = gen.PtrOf(DateAfterCreationStatusARMGenerator())
}

func Test_TagFilter_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TagFilter_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTagFilterStatusARM, TagFilterStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTagFilterStatusARM runs a test to see if a specific instance of TagFilter_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForTagFilterStatusARM(subject TagFilter_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TagFilter_StatusARM
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

// Generator of TagFilter_StatusARM instances for property testing - lazily instantiated by TagFilterStatusARMGenerator()
var tagFilterStatusARMGenerator gopter.Gen

// TagFilterStatusARMGenerator returns a generator of TagFilter_StatusARM instances for property testing.
func TagFilterStatusARMGenerator() gopter.Gen {
	if tagFilterStatusARMGenerator != nil {
		return tagFilterStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTagFilterStatusARM(generators)
	tagFilterStatusARMGenerator = gen.Struct(reflect.TypeOf(TagFilter_StatusARM{}), generators)

	return tagFilterStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForTagFilterStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTagFilterStatusARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Op"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_DateAfterCreation_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DateAfterCreation_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDateAfterCreationStatusARM, DateAfterCreationStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDateAfterCreationStatusARM runs a test to see if a specific instance of DateAfterCreation_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDateAfterCreationStatusARM(subject DateAfterCreation_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DateAfterCreation_StatusARM
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

// Generator of DateAfterCreation_StatusARM instances for property testing - lazily instantiated by
// DateAfterCreationStatusARMGenerator()
var dateAfterCreationStatusARMGenerator gopter.Gen

// DateAfterCreationStatusARMGenerator returns a generator of DateAfterCreation_StatusARM instances for property testing.
func DateAfterCreationStatusARMGenerator() gopter.Gen {
	if dateAfterCreationStatusARMGenerator != nil {
		return dateAfterCreationStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDateAfterCreationStatusARM(generators)
	dateAfterCreationStatusARMGenerator = gen.Struct(reflect.TypeOf(DateAfterCreation_StatusARM{}), generators)

	return dateAfterCreationStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForDateAfterCreationStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDateAfterCreationStatusARM(gens map[string]gopter.Gen) {
	gens["DaysAfterCreationGreaterThan"] = gen.PtrOf(gen.Float64())
}

func Test_DateAfterModification_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DateAfterModification_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDateAfterModificationStatusARM, DateAfterModificationStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDateAfterModificationStatusARM runs a test to see if a specific instance of DateAfterModification_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDateAfterModificationStatusARM(subject DateAfterModification_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DateAfterModification_StatusARM
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

// Generator of DateAfterModification_StatusARM instances for property testing - lazily instantiated by
// DateAfterModificationStatusARMGenerator()
var dateAfterModificationStatusARMGenerator gopter.Gen

// DateAfterModificationStatusARMGenerator returns a generator of DateAfterModification_StatusARM instances for property testing.
func DateAfterModificationStatusARMGenerator() gopter.Gen {
	if dateAfterModificationStatusARMGenerator != nil {
		return dateAfterModificationStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDateAfterModificationStatusARM(generators)
	dateAfterModificationStatusARMGenerator = gen.Struct(reflect.TypeOf(DateAfterModification_StatusARM{}), generators)

	return dateAfterModificationStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForDateAfterModificationStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDateAfterModificationStatusARM(gens map[string]gopter.Gen) {
	gens["DaysAfterLastAccessTimeGreaterThan"] = gen.PtrOf(gen.Float64())
	gens["DaysAfterModificationGreaterThan"] = gen.PtrOf(gen.Float64())
}
