// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101storage

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

func Test_NetworkSecurityGroupsSecurityRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroupsSecurityRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule, NetworkSecurityGroupsSecurityRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule runs a test to see if a specific instance of NetworkSecurityGroupsSecurityRule round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule(subject NetworkSecurityGroupsSecurityRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroupsSecurityRule
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

// Generator of NetworkSecurityGroupsSecurityRule instances for property testing - lazily instantiated by
// NetworkSecurityGroupsSecurityRuleGenerator()
var networkSecurityGroupsSecurityRuleGenerator gopter.Gen

// NetworkSecurityGroupsSecurityRuleGenerator returns a generator of NetworkSecurityGroupsSecurityRule instances for property testing.
func NetworkSecurityGroupsSecurityRuleGenerator() gopter.Gen {
	if networkSecurityGroupsSecurityRuleGenerator != nil {
		return networkSecurityGroupsSecurityRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule(generators)
	networkSecurityGroupsSecurityRuleGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupsSecurityRule{}), generators)

	return networkSecurityGroupsSecurityRuleGenerator
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule(gens map[string]gopter.Gen) {
	gens["Spec"] = NetworkSecurityGroupsSecurityRulesSpecGenerator()
	gens["Status"] = SecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator()
}

func Test_NetworkSecurityGroupsSecurityRules_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroupsSecurityRules_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupsSecurityRulesSpec, NetworkSecurityGroupsSecurityRulesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupsSecurityRulesSpec runs a test to see if a specific instance of NetworkSecurityGroupsSecurityRules_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupsSecurityRulesSpec(subject NetworkSecurityGroupsSecurityRules_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroupsSecurityRules_Spec
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

// Generator of NetworkSecurityGroupsSecurityRules_Spec instances for property testing - lazily instantiated by
// NetworkSecurityGroupsSecurityRulesSpecGenerator()
var networkSecurityGroupsSecurityRulesSpecGenerator gopter.Gen

// NetworkSecurityGroupsSecurityRulesSpecGenerator returns a generator of NetworkSecurityGroupsSecurityRules_Spec instances for property testing.
// We first initialize networkSecurityGroupsSecurityRulesSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroupsSecurityRulesSpecGenerator() gopter.Gen {
	if networkSecurityGroupsSecurityRulesSpecGenerator != nil {
		return networkSecurityGroupsSecurityRulesSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRulesSpec(generators)
	networkSecurityGroupsSecurityRulesSpecGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupsSecurityRules_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRulesSpec(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRulesSpec(generators)
	networkSecurityGroupsSecurityRulesSpecGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupsSecurityRules_Spec{}), generators)

	return networkSecurityGroupsSecurityRulesSpecGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRulesSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRulesSpec(gens map[string]gopter.Gen) {
	gens["Access"] = gen.PtrOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["DestinationPortRange"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationPortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Direction"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Protocol"] = gen.PtrOf(gen.AlphaString())
	gens["SourceAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["SourceAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["SourcePortRange"] = gen.PtrOf(gen.AlphaString())
	gens["SourcePortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRulesSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRulesSpec(gens map[string]gopter.Gen) {
	gens["DestinationApplicationSecurityGroups"] = gen.SliceOf(SubResourceGenerator())
	gens["SourceApplicationSecurityGroups"] = gen.SliceOf(SubResourceGenerator())
}

func Test_SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded, SecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded runs a test to see if a specific instance of SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded(subject SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded
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

// Generator of SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded instances for property testing
// - lazily instantiated by SecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator()
var securityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator gopter.Gen

// SecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator returns a generator of SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded instances for property testing.
// We first initialize securityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator() gopter.Gen {
	if securityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator != nil {
		return securityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded(generators)
	securityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded(generators)
	AddRelatedPropertyGeneratorsForSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded(generators)
	securityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded{}), generators)

	return securityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Access"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["DestinationPortRange"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationPortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Direction"] = gen.PtrOf(gen.AlphaString())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Protocol"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["SourceAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["SourceAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["SourcePortRange"] = gen.PtrOf(gen.AlphaString())
	gens["SourcePortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["DestinationApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator())
	gens["SourceApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator())
}

func Test_ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded, ApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded runs a test to see if a specific instance of ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded(subject ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded
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

// Generator of ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded instances for
// property testing - lazily instantiated by
// ApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator()
var applicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator gopter.Gen

// ApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator returns a generator of ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded instances for property testing.
func ApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator() gopter.Gen {
	if applicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator != nil {
		return applicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded(generators)
	applicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded{}), generators)

	return applicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
