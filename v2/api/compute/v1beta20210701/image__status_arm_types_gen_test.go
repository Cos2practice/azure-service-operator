// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210701

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

func Test_Image_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Image_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImageStatusARM, ImageStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImageStatusARM runs a test to see if a specific instance of Image_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImageStatusARM(subject Image_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Image_StatusARM
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

// Generator of Image_StatusARM instances for property testing - lazily instantiated by ImageStatusARMGenerator()
var imageStatusARMGenerator gopter.Gen

// ImageStatusARMGenerator returns a generator of Image_StatusARM instances for property testing.
// We first initialize imageStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImageStatusARMGenerator() gopter.Gen {
	if imageStatusARMGenerator != nil {
		return imageStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageStatusARM(generators)
	imageStatusARMGenerator = gen.Struct(reflect.TypeOf(Image_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageStatusARM(generators)
	AddRelatedPropertyGeneratorsForImageStatusARM(generators)
	imageStatusARMGenerator = gen.Struct(reflect.TypeOf(Image_StatusARM{}), generators)

	return imageStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForImageStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImageStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForImageStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImageStatusARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationStatusARMGenerator())
	gens["Properties"] = gen.PtrOf(ImagePropertiesStatusARMGenerator())
}

func Test_ExtendedLocation_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ExtendedLocation_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForExtendedLocationStatusARM, ExtendedLocationStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForExtendedLocationStatusARM runs a test to see if a specific instance of ExtendedLocation_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForExtendedLocationStatusARM(subject ExtendedLocation_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ExtendedLocation_StatusARM
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

// Generator of ExtendedLocation_StatusARM instances for property testing - lazily instantiated by
// ExtendedLocationStatusARMGenerator()
var extendedLocationStatusARMGenerator gopter.Gen

// ExtendedLocationStatusARMGenerator returns a generator of ExtendedLocation_StatusARM instances for property testing.
func ExtendedLocationStatusARMGenerator() gopter.Gen {
	if extendedLocationStatusARMGenerator != nil {
		return extendedLocationStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForExtendedLocationStatusARM(generators)
	extendedLocationStatusARMGenerator = gen.Struct(reflect.TypeOf(ExtendedLocation_StatusARM{}), generators)

	return extendedLocationStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForExtendedLocationStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForExtendedLocationStatusARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(ExtendedLocationType_StatusEdgeZone))
}

func Test_ImageProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImageProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImagePropertiesStatusARM, ImagePropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImagePropertiesStatusARM runs a test to see if a specific instance of ImageProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImagePropertiesStatusARM(subject ImageProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImageProperties_StatusARM
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

// Generator of ImageProperties_StatusARM instances for property testing - lazily instantiated by
// ImagePropertiesStatusARMGenerator()
var imagePropertiesStatusARMGenerator gopter.Gen

// ImagePropertiesStatusARMGenerator returns a generator of ImageProperties_StatusARM instances for property testing.
// We first initialize imagePropertiesStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImagePropertiesStatusARMGenerator() gopter.Gen {
	if imagePropertiesStatusARMGenerator != nil {
		return imagePropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImagePropertiesStatusARM(generators)
	imagePropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(ImageProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImagePropertiesStatusARM(generators)
	AddRelatedPropertyGeneratorsForImagePropertiesStatusARM(generators)
	imagePropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(ImageProperties_StatusARM{}), generators)

	return imagePropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForImagePropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImagePropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["HyperVGeneration"] = gen.PtrOf(gen.OneConstOf(HyperVGenerationType_StatusV1, HyperVGenerationType_StatusV2))
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForImagePropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImagePropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["SourceVirtualMachine"] = gen.PtrOf(SubResourceStatusARMGenerator())
	gens["StorageProfile"] = gen.PtrOf(ImageStorageProfileStatusARMGenerator())
}

func Test_ImageStorageProfile_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImageStorageProfile_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImageStorageProfileStatusARM, ImageStorageProfileStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImageStorageProfileStatusARM runs a test to see if a specific instance of ImageStorageProfile_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImageStorageProfileStatusARM(subject ImageStorageProfile_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImageStorageProfile_StatusARM
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

// Generator of ImageStorageProfile_StatusARM instances for property testing - lazily instantiated by
// ImageStorageProfileStatusARMGenerator()
var imageStorageProfileStatusARMGenerator gopter.Gen

// ImageStorageProfileStatusARMGenerator returns a generator of ImageStorageProfile_StatusARM instances for property testing.
// We first initialize imageStorageProfileStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImageStorageProfileStatusARMGenerator() gopter.Gen {
	if imageStorageProfileStatusARMGenerator != nil {
		return imageStorageProfileStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageStorageProfileStatusARM(generators)
	imageStorageProfileStatusARMGenerator = gen.Struct(reflect.TypeOf(ImageStorageProfile_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageStorageProfileStatusARM(generators)
	AddRelatedPropertyGeneratorsForImageStorageProfileStatusARM(generators)
	imageStorageProfileStatusARMGenerator = gen.Struct(reflect.TypeOf(ImageStorageProfile_StatusARM{}), generators)

	return imageStorageProfileStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForImageStorageProfileStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImageStorageProfileStatusARM(gens map[string]gopter.Gen) {
	gens["ZoneResilient"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForImageStorageProfileStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImageStorageProfileStatusARM(gens map[string]gopter.Gen) {
	gens["DataDisks"] = gen.SliceOf(ImageDataDiskStatusARMGenerator())
	gens["OsDisk"] = gen.PtrOf(ImageOSDiskStatusARMGenerator())
}

func Test_SubResource_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubResource_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubResourceStatusARM, SubResourceStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubResourceStatusARM runs a test to see if a specific instance of SubResource_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubResourceStatusARM(subject SubResource_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubResource_StatusARM
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

// Generator of SubResource_StatusARM instances for property testing - lazily instantiated by
// SubResourceStatusARMGenerator()
var subResourceStatusARMGenerator gopter.Gen

// SubResourceStatusARMGenerator returns a generator of SubResource_StatusARM instances for property testing.
func SubResourceStatusARMGenerator() gopter.Gen {
	if subResourceStatusARMGenerator != nil {
		return subResourceStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubResourceStatusARM(generators)
	subResourceStatusARMGenerator = gen.Struct(reflect.TypeOf(SubResource_StatusARM{}), generators)

	return subResourceStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSubResourceStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubResourceStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_ImageDataDisk_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImageDataDisk_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImageDataDiskStatusARM, ImageDataDiskStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImageDataDiskStatusARM runs a test to see if a specific instance of ImageDataDisk_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImageDataDiskStatusARM(subject ImageDataDisk_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImageDataDisk_StatusARM
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

// Generator of ImageDataDisk_StatusARM instances for property testing - lazily instantiated by
// ImageDataDiskStatusARMGenerator()
var imageDataDiskStatusARMGenerator gopter.Gen

// ImageDataDiskStatusARMGenerator returns a generator of ImageDataDisk_StatusARM instances for property testing.
// We first initialize imageDataDiskStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImageDataDiskStatusARMGenerator() gopter.Gen {
	if imageDataDiskStatusARMGenerator != nil {
		return imageDataDiskStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageDataDiskStatusARM(generators)
	imageDataDiskStatusARMGenerator = gen.Struct(reflect.TypeOf(ImageDataDisk_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageDataDiskStatusARM(generators)
	AddRelatedPropertyGeneratorsForImageDataDiskStatusARM(generators)
	imageDataDiskStatusARMGenerator = gen.Struct(reflect.TypeOf(ImageDataDisk_StatusARM{}), generators)

	return imageDataDiskStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForImageDataDiskStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImageDataDiskStatusARM(gens map[string]gopter.Gen) {
	gens["BlobUri"] = gen.PtrOf(gen.AlphaString())
	gens["Caching"] = gen.PtrOf(gen.OneConstOf(ImageDataDiskStatusCachingNone, ImageDataDiskStatusCachingReadOnly, ImageDataDiskStatusCachingReadWrite))
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["Lun"] = gen.PtrOf(gen.Int())
	gens["StorageAccountType"] = gen.PtrOf(gen.OneConstOf(
		StorageAccountType_StatusPremiumLRS,
		StorageAccountType_StatusPremiumZRS,
		StorageAccountType_StatusStandardLRS,
		StorageAccountType_StatusStandardSSDLRS,
		StorageAccountType_StatusStandardSSDZRS,
		StorageAccountType_StatusUltraSSDLRS))
}

// AddRelatedPropertyGeneratorsForImageDataDiskStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImageDataDiskStatusARM(gens map[string]gopter.Gen) {
	gens["DiskEncryptionSet"] = gen.PtrOf(SubResourceStatusARMGenerator())
	gens["ManagedDisk"] = gen.PtrOf(SubResourceStatusARMGenerator())
	gens["Snapshot"] = gen.PtrOf(SubResourceStatusARMGenerator())
}

func Test_ImageOSDisk_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImageOSDisk_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImageOSDiskStatusARM, ImageOSDiskStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImageOSDiskStatusARM runs a test to see if a specific instance of ImageOSDisk_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImageOSDiskStatusARM(subject ImageOSDisk_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImageOSDisk_StatusARM
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

// Generator of ImageOSDisk_StatusARM instances for property testing - lazily instantiated by
// ImageOSDiskStatusARMGenerator()
var imageOSDiskStatusARMGenerator gopter.Gen

// ImageOSDiskStatusARMGenerator returns a generator of ImageOSDisk_StatusARM instances for property testing.
// We first initialize imageOSDiskStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImageOSDiskStatusARMGenerator() gopter.Gen {
	if imageOSDiskStatusARMGenerator != nil {
		return imageOSDiskStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageOSDiskStatusARM(generators)
	imageOSDiskStatusARMGenerator = gen.Struct(reflect.TypeOf(ImageOSDisk_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageOSDiskStatusARM(generators)
	AddRelatedPropertyGeneratorsForImageOSDiskStatusARM(generators)
	imageOSDiskStatusARMGenerator = gen.Struct(reflect.TypeOf(ImageOSDisk_StatusARM{}), generators)

	return imageOSDiskStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForImageOSDiskStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImageOSDiskStatusARM(gens map[string]gopter.Gen) {
	gens["BlobUri"] = gen.PtrOf(gen.AlphaString())
	gens["Caching"] = gen.PtrOf(gen.OneConstOf(ImageOSDiskStatusCachingNone, ImageOSDiskStatusCachingReadOnly, ImageOSDiskStatusCachingReadWrite))
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["OsState"] = gen.PtrOf(gen.OneConstOf(ImageOSDiskStatusOsStateGeneralized, ImageOSDiskStatusOsStateSpecialized))
	gens["OsType"] = gen.PtrOf(gen.OneConstOf(ImageOSDiskStatusOsTypeLinux, ImageOSDiskStatusOsTypeWindows))
	gens["StorageAccountType"] = gen.PtrOf(gen.OneConstOf(
		StorageAccountType_StatusPremiumLRS,
		StorageAccountType_StatusPremiumZRS,
		StorageAccountType_StatusStandardLRS,
		StorageAccountType_StatusStandardSSDLRS,
		StorageAccountType_StatusStandardSSDZRS,
		StorageAccountType_StatusUltraSSDLRS))
}

// AddRelatedPropertyGeneratorsForImageOSDiskStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImageOSDiskStatusARM(gens map[string]gopter.Gen) {
	gens["DiskEncryptionSet"] = gen.PtrOf(SubResourceStatusARMGenerator())
	gens["ManagedDisk"] = gen.PtrOf(SubResourceStatusARMGenerator())
	gens["Snapshot"] = gen.PtrOf(SubResourceStatusARMGenerator())
}
