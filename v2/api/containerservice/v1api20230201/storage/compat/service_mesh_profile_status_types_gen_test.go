// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package compat

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231001/storage"
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

func Test_IstioComponents_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from IstioComponents_STATUS to IstioComponents_STATUS via AssignProperties_To_IstioComponents_STATUS & AssignProperties_From_IstioComponents_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForIstioComponents_STATUS, IstioComponents_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForIstioComponents_STATUS tests if a specific instance of IstioComponents_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForIstioComponents_STATUS(subject IstioComponents_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.IstioComponents_STATUS
	err := copied.AssignProperties_To_IstioComponents_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual IstioComponents_STATUS
	err = actual.AssignProperties_From_IstioComponents_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_IstioComponents_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IstioComponents_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIstioComponents_STATUS, IstioComponents_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIstioComponents_STATUS runs a test to see if a specific instance of IstioComponents_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForIstioComponents_STATUS(subject IstioComponents_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IstioComponents_STATUS
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

// Generator of IstioComponents_STATUS instances for property testing - lazily instantiated by
// IstioComponents_STATUSGenerator()
var istioComponents_STATUSGenerator gopter.Gen

// IstioComponents_STATUSGenerator returns a generator of IstioComponents_STATUS instances for property testing.
func IstioComponents_STATUSGenerator() gopter.Gen {
	if istioComponents_STATUSGenerator != nil {
		return istioComponents_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForIstioComponents_STATUS(generators)
	istioComponents_STATUSGenerator = gen.Struct(reflect.TypeOf(IstioComponents_STATUS{}), generators)

	return istioComponents_STATUSGenerator
}

// AddRelatedPropertyGeneratorsForIstioComponents_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIstioComponents_STATUS(gens map[string]gopter.Gen) {
	gens["IngressGateways"] = gen.SliceOf(IstioIngressGateway_STATUSGenerator())
}

func Test_IstioIngressGateway_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from IstioIngressGateway_STATUS to IstioIngressGateway_STATUS via AssignProperties_To_IstioIngressGateway_STATUS & AssignProperties_From_IstioIngressGateway_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForIstioIngressGateway_STATUS, IstioIngressGateway_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForIstioIngressGateway_STATUS tests if a specific instance of IstioIngressGateway_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForIstioIngressGateway_STATUS(subject IstioIngressGateway_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.IstioIngressGateway_STATUS
	err := copied.AssignProperties_To_IstioIngressGateway_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual IstioIngressGateway_STATUS
	err = actual.AssignProperties_From_IstioIngressGateway_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_IstioIngressGateway_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IstioIngressGateway_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIstioIngressGateway_STATUS, IstioIngressGateway_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIstioIngressGateway_STATUS runs a test to see if a specific instance of IstioIngressGateway_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForIstioIngressGateway_STATUS(subject IstioIngressGateway_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IstioIngressGateway_STATUS
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

// Generator of IstioIngressGateway_STATUS instances for property testing - lazily instantiated by
// IstioIngressGateway_STATUSGenerator()
var istioIngressGateway_STATUSGenerator gopter.Gen

// IstioIngressGateway_STATUSGenerator returns a generator of IstioIngressGateway_STATUS instances for property testing.
func IstioIngressGateway_STATUSGenerator() gopter.Gen {
	if istioIngressGateway_STATUSGenerator != nil {
		return istioIngressGateway_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIstioIngressGateway_STATUS(generators)
	istioIngressGateway_STATUSGenerator = gen.Struct(reflect.TypeOf(IstioIngressGateway_STATUS{}), generators)

	return istioIngressGateway_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForIstioIngressGateway_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIstioIngressGateway_STATUS(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Mode"] = gen.PtrOf(gen.AlphaString())
}

func Test_IstioServiceMesh_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from IstioServiceMesh_STATUS to IstioServiceMesh_STATUS via AssignProperties_To_IstioServiceMesh_STATUS & AssignProperties_From_IstioServiceMesh_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForIstioServiceMesh_STATUS, IstioServiceMesh_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForIstioServiceMesh_STATUS tests if a specific instance of IstioServiceMesh_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForIstioServiceMesh_STATUS(subject IstioServiceMesh_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.IstioServiceMesh_STATUS
	err := copied.AssignProperties_To_IstioServiceMesh_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual IstioServiceMesh_STATUS
	err = actual.AssignProperties_From_IstioServiceMesh_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_IstioServiceMesh_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IstioServiceMesh_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIstioServiceMesh_STATUS, IstioServiceMesh_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIstioServiceMesh_STATUS runs a test to see if a specific instance of IstioServiceMesh_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForIstioServiceMesh_STATUS(subject IstioServiceMesh_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IstioServiceMesh_STATUS
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

// Generator of IstioServiceMesh_STATUS instances for property testing - lazily instantiated by
// IstioServiceMesh_STATUSGenerator()
var istioServiceMesh_STATUSGenerator gopter.Gen

// IstioServiceMesh_STATUSGenerator returns a generator of IstioServiceMesh_STATUS instances for property testing.
func IstioServiceMesh_STATUSGenerator() gopter.Gen {
	if istioServiceMesh_STATUSGenerator != nil {
		return istioServiceMesh_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForIstioServiceMesh_STATUS(generators)
	istioServiceMesh_STATUSGenerator = gen.Struct(reflect.TypeOf(IstioServiceMesh_STATUS{}), generators)

	return istioServiceMesh_STATUSGenerator
}

// AddRelatedPropertyGeneratorsForIstioServiceMesh_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIstioServiceMesh_STATUS(gens map[string]gopter.Gen) {
	gens["Components"] = gen.PtrOf(IstioComponents_STATUSGenerator())
}

func Test_ServiceMeshProfile_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServiceMeshProfile_STATUS to ServiceMeshProfile_STATUS via AssignProperties_To_ServiceMeshProfile_STATUS & AssignProperties_From_ServiceMeshProfile_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForServiceMeshProfile_STATUS, ServiceMeshProfile_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServiceMeshProfile_STATUS tests if a specific instance of ServiceMeshProfile_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForServiceMeshProfile_STATUS(subject ServiceMeshProfile_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.ServiceMeshProfile_STATUS
	err := copied.AssignProperties_To_ServiceMeshProfile_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ServiceMeshProfile_STATUS
	err = actual.AssignProperties_From_ServiceMeshProfile_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ServiceMeshProfile_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServiceMeshProfile_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServiceMeshProfile_STATUS, ServiceMeshProfile_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServiceMeshProfile_STATUS runs a test to see if a specific instance of ServiceMeshProfile_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServiceMeshProfile_STATUS(subject ServiceMeshProfile_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServiceMeshProfile_STATUS
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

// Generator of ServiceMeshProfile_STATUS instances for property testing - lazily instantiated by
// ServiceMeshProfile_STATUSGenerator()
var serviceMeshProfile_STATUSGenerator gopter.Gen

// ServiceMeshProfile_STATUSGenerator returns a generator of ServiceMeshProfile_STATUS instances for property testing.
// We first initialize serviceMeshProfile_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServiceMeshProfile_STATUSGenerator() gopter.Gen {
	if serviceMeshProfile_STATUSGenerator != nil {
		return serviceMeshProfile_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServiceMeshProfile_STATUS(generators)
	serviceMeshProfile_STATUSGenerator = gen.Struct(reflect.TypeOf(ServiceMeshProfile_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServiceMeshProfile_STATUS(generators)
	AddRelatedPropertyGeneratorsForServiceMeshProfile_STATUS(generators)
	serviceMeshProfile_STATUSGenerator = gen.Struct(reflect.TypeOf(ServiceMeshProfile_STATUS{}), generators)

	return serviceMeshProfile_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServiceMeshProfile_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServiceMeshProfile_STATUS(gens map[string]gopter.Gen) {
	gens["Mode"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServiceMeshProfile_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServiceMeshProfile_STATUS(gens map[string]gopter.Gen) {
	gens["Istio"] = gen.PtrOf(IstioServiceMesh_STATUSGenerator())
}