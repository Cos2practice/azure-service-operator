// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180601

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

func Test_Server_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Server_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServer_STATUSARM, Server_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServer_STATUSARM runs a test to see if a specific instance of Server_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServer_STATUSARM(subject Server_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Server_STATUSARM
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

// Generator of Server_STATUSARM instances for property testing - lazily instantiated by Server_STATUSARMGenerator()
var server_STATUSARMGenerator gopter.Gen

// Server_STATUSARMGenerator returns a generator of Server_STATUSARM instances for property testing.
// We first initialize server_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Server_STATUSARMGenerator() gopter.Gen {
	if server_STATUSARMGenerator != nil {
		return server_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServer_STATUSARM(generators)
	server_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Server_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServer_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForServer_STATUSARM(generators)
	server_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Server_STATUSARM{}), generators)

	return server_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForServer_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServer_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServer_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServer_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ServerProperties_STATUSARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUSARMGenerator())
}

func Test_ServerProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServerProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerProperties_STATUSARM, ServerProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerProperties_STATUSARM runs a test to see if a specific instance of ServerProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServerProperties_STATUSARM(subject ServerProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServerProperties_STATUSARM
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

// Generator of ServerProperties_STATUSARM instances for property testing - lazily instantiated by
// ServerProperties_STATUSARMGenerator()
var serverProperties_STATUSARMGenerator gopter.Gen

// ServerProperties_STATUSARMGenerator returns a generator of ServerProperties_STATUSARM instances for property testing.
// We first initialize serverProperties_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServerProperties_STATUSARMGenerator() gopter.Gen {
	if serverProperties_STATUSARMGenerator != nil {
		return serverProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerProperties_STATUSARM(generators)
	serverProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ServerProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerProperties_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForServerProperties_STATUSARM(generators)
	serverProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ServerProperties_STATUSARM{}), generators)

	return serverProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForServerProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["AdministratorLogin"] = gen.PtrOf(gen.AlphaString())
	gens["EarliestRestoreDate"] = gen.PtrOf(gen.AlphaString())
	gens["FullyQualifiedDomainName"] = gen.PtrOf(gen.AlphaString())
	gens["MasterServerId"] = gen.PtrOf(gen.AlphaString())
	gens["MinimalTlsVersion"] = gen.PtrOf(gen.OneConstOf(
		MinimalTlsVersion_STATUS_TLS1_0,
		MinimalTlsVersion_STATUS_TLS1_1,
		MinimalTlsVersion_STATUS_TLS1_2,
		MinimalTlsVersion_STATUS_TLSEnforcementDisabled))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccess_STATUS_Disabled, PublicNetworkAccess_STATUS_Enabled))
	gens["ReplicaCapacity"] = gen.PtrOf(gen.Int())
	gens["ReplicationRole"] = gen.PtrOf(gen.AlphaString())
	gens["SslEnforcement"] = gen.PtrOf(gen.OneConstOf(SslEnforcement_STATUS_Disabled, SslEnforcement_STATUS_Enabled))
	gens["UserVisibleState"] = gen.PtrOf(gen.OneConstOf(ServerProperties_STATUS_UserVisibleState_Disabled, ServerProperties_STATUS_UserVisibleState_Dropping, ServerProperties_STATUS_UserVisibleState_Ready))
	gens["Version"] = gen.PtrOf(gen.OneConstOf(ServerVersion_STATUS_102, ServerVersion_STATUS_103))
}

// AddRelatedPropertyGeneratorsForServerProperties_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServerProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["PrivateEndpointConnections"] = gen.SliceOf(ServerPrivateEndpointConnection_STATUSARMGenerator())
	gens["StorageProfile"] = gen.PtrOf(StorageProfile_STATUSARMGenerator())
}

func Test_Sku_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku_STATUSARM, Sku_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_STATUSARM runs a test to see if a specific instance of Sku_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_STATUSARM(subject Sku_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_STATUSARM
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

// Generator of Sku_STATUSARM instances for property testing - lazily instantiated by Sku_STATUSARMGenerator()
var sku_STATUSARMGenerator gopter.Gen

// Sku_STATUSARMGenerator returns a generator of Sku_STATUSARM instances for property testing.
func Sku_STATUSARMGenerator() gopter.Gen {
	if sku_STATUSARMGenerator != nil {
		return sku_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_STATUSARM(generators)
	sku_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Sku_STATUSARM{}), generators)

	return sku_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSku_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_STATUSARM(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Family"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Size"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.OneConstOf(Sku_STATUS_Tier_Basic, Sku_STATUS_Tier_GeneralPurpose, Sku_STATUS_Tier_MemoryOptimized))
}

func Test_ServerPrivateEndpointConnection_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServerPrivateEndpointConnection_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerPrivateEndpointConnection_STATUSARM, ServerPrivateEndpointConnection_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerPrivateEndpointConnection_STATUSARM runs a test to see if a specific instance of ServerPrivateEndpointConnection_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServerPrivateEndpointConnection_STATUSARM(subject ServerPrivateEndpointConnection_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServerPrivateEndpointConnection_STATUSARM
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

// Generator of ServerPrivateEndpointConnection_STATUSARM instances for property testing - lazily instantiated by
// ServerPrivateEndpointConnection_STATUSARMGenerator()
var serverPrivateEndpointConnection_STATUSARMGenerator gopter.Gen

// ServerPrivateEndpointConnection_STATUSARMGenerator returns a generator of ServerPrivateEndpointConnection_STATUSARM instances for property testing.
// We first initialize serverPrivateEndpointConnection_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServerPrivateEndpointConnection_STATUSARMGenerator() gopter.Gen {
	if serverPrivateEndpointConnection_STATUSARMGenerator != nil {
		return serverPrivateEndpointConnection_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerPrivateEndpointConnection_STATUSARM(generators)
	serverPrivateEndpointConnection_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ServerPrivateEndpointConnection_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerPrivateEndpointConnection_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForServerPrivateEndpointConnection_STATUSARM(generators)
	serverPrivateEndpointConnection_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ServerPrivateEndpointConnection_STATUSARM{}), generators)

	return serverPrivateEndpointConnection_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForServerPrivateEndpointConnection_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerPrivateEndpointConnection_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServerPrivateEndpointConnection_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServerPrivateEndpointConnection_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ServerPrivateEndpointConnectionProperties_STATUSARMGenerator())
}

func Test_StorageProfile_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageProfile_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageProfile_STATUSARM, StorageProfile_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageProfile_STATUSARM runs a test to see if a specific instance of StorageProfile_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageProfile_STATUSARM(subject StorageProfile_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageProfile_STATUSARM
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

// Generator of StorageProfile_STATUSARM instances for property testing - lazily instantiated by
// StorageProfile_STATUSARMGenerator()
var storageProfile_STATUSARMGenerator gopter.Gen

// StorageProfile_STATUSARMGenerator returns a generator of StorageProfile_STATUSARM instances for property testing.
func StorageProfile_STATUSARMGenerator() gopter.Gen {
	if storageProfile_STATUSARMGenerator != nil {
		return storageProfile_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageProfile_STATUSARM(generators)
	storageProfile_STATUSARMGenerator = gen.Struct(reflect.TypeOf(StorageProfile_STATUSARM{}), generators)

	return storageProfile_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForStorageProfile_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageProfile_STATUSARM(gens map[string]gopter.Gen) {
	gens["BackupRetentionDays"] = gen.PtrOf(gen.Int())
	gens["GeoRedundantBackup"] = gen.PtrOf(gen.OneConstOf(StorageProfile_STATUS_GeoRedundantBackup_Disabled, StorageProfile_STATUS_GeoRedundantBackup_Enabled))
	gens["StorageAutogrow"] = gen.PtrOf(gen.OneConstOf(StorageProfile_STATUS_StorageAutogrow_Disabled, StorageProfile_STATUS_StorageAutogrow_Enabled))
	gens["StorageMB"] = gen.PtrOf(gen.Int())
}

func Test_ServerPrivateEndpointConnectionProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServerPrivateEndpointConnectionProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerPrivateEndpointConnectionProperties_STATUSARM, ServerPrivateEndpointConnectionProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerPrivateEndpointConnectionProperties_STATUSARM runs a test to see if a specific instance of ServerPrivateEndpointConnectionProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServerPrivateEndpointConnectionProperties_STATUSARM(subject ServerPrivateEndpointConnectionProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServerPrivateEndpointConnectionProperties_STATUSARM
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

// Generator of ServerPrivateEndpointConnectionProperties_STATUSARM instances for property testing - lazily instantiated
// by ServerPrivateEndpointConnectionProperties_STATUSARMGenerator()
var serverPrivateEndpointConnectionProperties_STATUSARMGenerator gopter.Gen

// ServerPrivateEndpointConnectionProperties_STATUSARMGenerator returns a generator of ServerPrivateEndpointConnectionProperties_STATUSARM instances for property testing.
// We first initialize serverPrivateEndpointConnectionProperties_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServerPrivateEndpointConnectionProperties_STATUSARMGenerator() gopter.Gen {
	if serverPrivateEndpointConnectionProperties_STATUSARMGenerator != nil {
		return serverPrivateEndpointConnectionProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerPrivateEndpointConnectionProperties_STATUSARM(generators)
	serverPrivateEndpointConnectionProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ServerPrivateEndpointConnectionProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerPrivateEndpointConnectionProperties_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForServerPrivateEndpointConnectionProperties_STATUSARM(generators)
	serverPrivateEndpointConnectionProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ServerPrivateEndpointConnectionProperties_STATUSARM{}), generators)

	return serverPrivateEndpointConnectionProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForServerPrivateEndpointConnectionProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerPrivateEndpointConnectionProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ServerPrivateEndpointConnectionProperties_STATUS_ProvisioningState_Approving,
		ServerPrivateEndpointConnectionProperties_STATUS_ProvisioningState_Dropping,
		ServerPrivateEndpointConnectionProperties_STATUS_ProvisioningState_Failed,
		ServerPrivateEndpointConnectionProperties_STATUS_ProvisioningState_Ready,
		ServerPrivateEndpointConnectionProperties_STATUS_ProvisioningState_Rejecting))
}

// AddRelatedPropertyGeneratorsForServerPrivateEndpointConnectionProperties_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServerPrivateEndpointConnectionProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["PrivateEndpoint"] = gen.PtrOf(PrivateEndpointProperty_STATUSARMGenerator())
	gens["PrivateLinkServiceConnectionState"] = gen.PtrOf(ServerPrivateLinkServiceConnectionStateProperty_STATUSARMGenerator())
}

func Test_PrivateEndpointProperty_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointProperty_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointProperty_STATUSARM, PrivateEndpointProperty_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointProperty_STATUSARM runs a test to see if a specific instance of PrivateEndpointProperty_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointProperty_STATUSARM(subject PrivateEndpointProperty_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointProperty_STATUSARM
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

// Generator of PrivateEndpointProperty_STATUSARM instances for property testing - lazily instantiated by
// PrivateEndpointProperty_STATUSARMGenerator()
var privateEndpointProperty_STATUSARMGenerator gopter.Gen

// PrivateEndpointProperty_STATUSARMGenerator returns a generator of PrivateEndpointProperty_STATUSARM instances for property testing.
func PrivateEndpointProperty_STATUSARMGenerator() gopter.Gen {
	if privateEndpointProperty_STATUSARMGenerator != nil {
		return privateEndpointProperty_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointProperty_STATUSARM(generators)
	privateEndpointProperty_STATUSARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointProperty_STATUSARM{}), generators)

	return privateEndpointProperty_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointProperty_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointProperty_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServerPrivateLinkServiceConnectionStateProperty_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServerPrivateLinkServiceConnectionStateProperty_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerPrivateLinkServiceConnectionStateProperty_STATUSARM, ServerPrivateLinkServiceConnectionStateProperty_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerPrivateLinkServiceConnectionStateProperty_STATUSARM runs a test to see if a specific instance of ServerPrivateLinkServiceConnectionStateProperty_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServerPrivateLinkServiceConnectionStateProperty_STATUSARM(subject ServerPrivateLinkServiceConnectionStateProperty_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServerPrivateLinkServiceConnectionStateProperty_STATUSARM
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

// Generator of ServerPrivateLinkServiceConnectionStateProperty_STATUSARM instances for property testing - lazily
// instantiated by ServerPrivateLinkServiceConnectionStateProperty_STATUSARMGenerator()
var serverPrivateLinkServiceConnectionStateProperty_STATUSARMGenerator gopter.Gen

// ServerPrivateLinkServiceConnectionStateProperty_STATUSARMGenerator returns a generator of ServerPrivateLinkServiceConnectionStateProperty_STATUSARM instances for property testing.
func ServerPrivateLinkServiceConnectionStateProperty_STATUSARMGenerator() gopter.Gen {
	if serverPrivateLinkServiceConnectionStateProperty_STATUSARMGenerator != nil {
		return serverPrivateLinkServiceConnectionStateProperty_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerPrivateLinkServiceConnectionStateProperty_STATUSARM(generators)
	serverPrivateLinkServiceConnectionStateProperty_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ServerPrivateLinkServiceConnectionStateProperty_STATUSARM{}), generators)

	return serverPrivateLinkServiceConnectionStateProperty_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForServerPrivateLinkServiceConnectionStateProperty_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerPrivateLinkServiceConnectionStateProperty_STATUSARM(gens map[string]gopter.Gen) {
	gens["ActionsRequired"] = gen.PtrOf(gen.OneConstOf(ServerPrivateLinkServiceConnectionStateProperty_STATUS_ActionsRequired_None))
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(
		ServerPrivateLinkServiceConnectionStateProperty_STATUS_Status_Approved,
		ServerPrivateLinkServiceConnectionStateProperty_STATUS_Status_Disconnected,
		ServerPrivateLinkServiceConnectionStateProperty_STATUS_Status_Pending,
		ServerPrivateLinkServiceConnectionStateProperty_STATUS_Status_Rejected))
}