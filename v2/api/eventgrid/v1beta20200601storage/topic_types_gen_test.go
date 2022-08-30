// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200601storage

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

func Test_Topic_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Topic via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTopic, TopicGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTopic runs a test to see if a specific instance of Topic round trips to JSON and back losslessly
func RunJSONSerializationTestForTopic(subject Topic) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Topic
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

// Generator of Topic instances for property testing - lazily instantiated by TopicGenerator()
var topicGenerator gopter.Gen

// TopicGenerator returns a generator of Topic instances for property testing.
func TopicGenerator() gopter.Gen {
	if topicGenerator != nil {
		return topicGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForTopic(generators)
	topicGenerator = gen.Struct(reflect.TypeOf(Topic{}), generators)

	return topicGenerator
}

// AddRelatedPropertyGeneratorsForTopic is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTopic(gens map[string]gopter.Gen) {
	gens["Spec"] = Topics_SpecGenerator()
	gens["Status"] = Topic_STATUSGenerator()
}

func Test_Topic_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Topic_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTopic_STATUS, Topic_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTopic_STATUS runs a test to see if a specific instance of Topic_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForTopic_STATUS(subject Topic_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Topic_STATUS
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

// Generator of Topic_STATUS instances for property testing - lazily instantiated by Topic_STATUSGenerator()
var topic_STATUSGenerator gopter.Gen

// Topic_STATUSGenerator returns a generator of Topic_STATUS instances for property testing.
// We first initialize topic_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Topic_STATUSGenerator() gopter.Gen {
	if topic_STATUSGenerator != nil {
		return topic_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTopic_STATUS(generators)
	topic_STATUSGenerator = gen.Struct(reflect.TypeOf(Topic_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTopic_STATUS(generators)
	AddRelatedPropertyGeneratorsForTopic_STATUS(generators)
	topic_STATUSGenerator = gen.Struct(reflect.TypeOf(Topic_STATUS{}), generators)

	return topic_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForTopic_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTopic_STATUS(gens map[string]gopter.Gen) {
	gens["Endpoint"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["InputSchema"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MetricResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForTopic_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTopic_STATUS(gens map[string]gopter.Gen) {
	gens["InboundIpRules"] = gen.SliceOf(InboundIpRule_STATUSGenerator())
	gens["InputSchemaMapping"] = gen.PtrOf(InputSchemaMapping_STATUSGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnection_STATUS_Topic_SubResourceEmbeddedGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_Topics_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Topics_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTopics_Spec, Topics_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTopics_Spec runs a test to see if a specific instance of Topics_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForTopics_Spec(subject Topics_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Topics_Spec
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

// Generator of Topics_Spec instances for property testing - lazily instantiated by Topics_SpecGenerator()
var topics_SpecGenerator gopter.Gen

// Topics_SpecGenerator returns a generator of Topics_Spec instances for property testing.
func Topics_SpecGenerator() gopter.Gen {
	if topics_SpecGenerator != nil {
		return topics_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTopics_Spec(generators)
	topics_SpecGenerator = gen.Struct(reflect.TypeOf(Topics_Spec{}), generators)

	return topics_SpecGenerator
}

// AddIndependentPropertyGeneratorsForTopics_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTopics_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

func Test_PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded, PrivateEndpointConnection_STATUS_Topic_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded runs a test to see if a specific instance of PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded(subject PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded
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

// Generator of PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded instances for property testing - lazily
// instantiated by PrivateEndpointConnection_STATUS_Topic_SubResourceEmbeddedGenerator()
var privateEndpointConnection_STATUS_Topic_SubResourceEmbeddedGenerator gopter.Gen

// PrivateEndpointConnection_STATUS_Topic_SubResourceEmbeddedGenerator returns a generator of PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded instances for property testing.
func PrivateEndpointConnection_STATUS_Topic_SubResourceEmbeddedGenerator() gopter.Gen {
	if privateEndpointConnection_STATUS_Topic_SubResourceEmbeddedGenerator != nil {
		return privateEndpointConnection_STATUS_Topic_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded(generators)
	privateEndpointConnection_STATUS_Topic_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded{}), generators)

	return privateEndpointConnection_STATUS_Topic_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}