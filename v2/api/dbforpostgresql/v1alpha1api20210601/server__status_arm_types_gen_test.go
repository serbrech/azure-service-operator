// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

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

func Test_Server_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Server_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerStatusARM, ServerStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerStatusARM runs a test to see if a specific instance of Server_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServerStatusARM(subject Server_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Server_StatusARM
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

// Generator of Server_StatusARM instances for property testing - lazily instantiated by ServerStatusARMGenerator()
var serverStatusARMGenerator gopter.Gen

// ServerStatusARMGenerator returns a generator of Server_StatusARM instances for property testing.
// We first initialize serverStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServerStatusARMGenerator() gopter.Gen {
	if serverStatusARMGenerator != nil {
		return serverStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerStatusARM(generators)
	serverStatusARMGenerator = gen.Struct(reflect.TypeOf(Server_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerStatusARM(generators)
	AddRelatedPropertyGeneratorsForServerStatusARM(generators)
	serverStatusARMGenerator = gen.Struct(reflect.TypeOf(Server_StatusARM{}), generators)

	return serverStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForServerStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServerStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServerStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ServerPropertiesStatusARMGenerator())
	gens["Sku"] = gen.PtrOf(SkuStatusARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataStatusARMGenerator())
}

func Test_ServerProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServerProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerPropertiesStatusARM, ServerPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerPropertiesStatusARM runs a test to see if a specific instance of ServerProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServerPropertiesStatusARM(subject ServerProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServerProperties_StatusARM
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

// Generator of ServerProperties_StatusARM instances for property testing - lazily instantiated by
//ServerPropertiesStatusARMGenerator()
var serverPropertiesStatusARMGenerator gopter.Gen

// ServerPropertiesStatusARMGenerator returns a generator of ServerProperties_StatusARM instances for property testing.
// We first initialize serverPropertiesStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServerPropertiesStatusARMGenerator() gopter.Gen {
	if serverPropertiesStatusARMGenerator != nil {
		return serverPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerPropertiesStatusARM(generators)
	serverPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(ServerProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerPropertiesStatusARM(generators)
	AddRelatedPropertyGeneratorsForServerPropertiesStatusARM(generators)
	serverPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(ServerProperties_StatusARM{}), generators)

	return serverPropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForServerPropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["AdministratorLogin"] = gen.PtrOf(gen.AlphaString())
	gens["AvailabilityZone"] = gen.PtrOf(gen.AlphaString())
	gens["CreateMode"] = gen.PtrOf(gen.OneConstOf(
		ServerPropertiesStatusCreateModeCreate,
		ServerPropertiesStatusCreateModeDefault,
		ServerPropertiesStatusCreateModePointInTimeRestore,
		ServerPropertiesStatusCreateModeUpdate))
	gens["FullyQualifiedDomainName"] = gen.PtrOf(gen.AlphaString())
	gens["MinorVersion"] = gen.PtrOf(gen.AlphaString())
	gens["PointInTimeUTC"] = gen.PtrOf(gen.AlphaString())
	gens["SourceServerResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.OneConstOf(
		ServerPropertiesStatusStateDisabled,
		ServerPropertiesStatusStateDropping,
		ServerPropertiesStatusStateReady,
		ServerPropertiesStatusStateStarting,
		ServerPropertiesStatusStateStopped,
		ServerPropertiesStatusStateStopping,
		ServerPropertiesStatusStateUpdating))
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Version"] = gen.PtrOf(gen.OneConstOf(ServerVersion_Status11, ServerVersion_Status12, ServerVersion_Status13))
}

// AddRelatedPropertyGeneratorsForServerPropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServerPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["Backup"] = gen.PtrOf(BackupStatusARMGenerator())
	gens["HighAvailability"] = gen.PtrOf(HighAvailabilityStatusARMGenerator())
	gens["MaintenanceWindow"] = gen.PtrOf(MaintenanceWindowStatusARMGenerator())
	gens["Network"] = gen.PtrOf(NetworkStatusARMGenerator())
	gens["Storage"] = gen.PtrOf(StorageStatusARMGenerator())
}

func Test_Sku_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSkuStatusARM, SkuStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSkuStatusARM runs a test to see if a specific instance of Sku_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSkuStatusARM(subject Sku_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_StatusARM
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

// Generator of Sku_StatusARM instances for property testing - lazily instantiated by SkuStatusARMGenerator()
var skuStatusARMGenerator gopter.Gen

// SkuStatusARMGenerator returns a generator of Sku_StatusARM instances for property testing.
func SkuStatusARMGenerator() gopter.Gen {
	if skuStatusARMGenerator != nil {
		return skuStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSkuStatusARM(generators)
	skuStatusARMGenerator = gen.Struct(reflect.TypeOf(Sku_StatusARM{}), generators)

	return skuStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSkuStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSkuStatusARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
	gens["Tier"] = gen.OneConstOf(SkuStatusTierBurstable, SkuStatusTierGeneralPurpose, SkuStatusTierMemoryOptimized)
}

func Test_Backup_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Backup_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackupStatusARM, BackupStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackupStatusARM runs a test to see if a specific instance of Backup_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBackupStatusARM(subject Backup_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Backup_StatusARM
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

// Generator of Backup_StatusARM instances for property testing - lazily instantiated by BackupStatusARMGenerator()
var backupStatusARMGenerator gopter.Gen

// BackupStatusARMGenerator returns a generator of Backup_StatusARM instances for property testing.
func BackupStatusARMGenerator() gopter.Gen {
	if backupStatusARMGenerator != nil {
		return backupStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackupStatusARM(generators)
	backupStatusARMGenerator = gen.Struct(reflect.TypeOf(Backup_StatusARM{}), generators)

	return backupStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForBackupStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackupStatusARM(gens map[string]gopter.Gen) {
	gens["BackupRetentionDays"] = gen.PtrOf(gen.Int())
	gens["EarliestRestoreDate"] = gen.PtrOf(gen.AlphaString())
	gens["GeoRedundantBackup"] = gen.PtrOf(gen.OneConstOf(BackupStatusGeoRedundantBackupDisabled, BackupStatusGeoRedundantBackupEnabled))
}

func Test_HighAvailability_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of HighAvailability_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForHighAvailabilityStatusARM, HighAvailabilityStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForHighAvailabilityStatusARM runs a test to see if a specific instance of HighAvailability_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForHighAvailabilityStatusARM(subject HighAvailability_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual HighAvailability_StatusARM
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

// Generator of HighAvailability_StatusARM instances for property testing - lazily instantiated by
//HighAvailabilityStatusARMGenerator()
var highAvailabilityStatusARMGenerator gopter.Gen

// HighAvailabilityStatusARMGenerator returns a generator of HighAvailability_StatusARM instances for property testing.
func HighAvailabilityStatusARMGenerator() gopter.Gen {
	if highAvailabilityStatusARMGenerator != nil {
		return highAvailabilityStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForHighAvailabilityStatusARM(generators)
	highAvailabilityStatusARMGenerator = gen.Struct(reflect.TypeOf(HighAvailability_StatusARM{}), generators)

	return highAvailabilityStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForHighAvailabilityStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForHighAvailabilityStatusARM(gens map[string]gopter.Gen) {
	gens["Mode"] = gen.PtrOf(gen.OneConstOf(HighAvailabilityStatusModeDisabled, HighAvailabilityStatusModeZoneRedundant))
	gens["StandbyAvailabilityZone"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.OneConstOf(
		HighAvailabilityStatusStateCreatingStandby,
		HighAvailabilityStatusStateFailingOver,
		HighAvailabilityStatusStateHealthy,
		HighAvailabilityStatusStateNotEnabled,
		HighAvailabilityStatusStateRemovingStandby,
		HighAvailabilityStatusStateReplicatingData))
}

func Test_MaintenanceWindow_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MaintenanceWindow_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMaintenanceWindowStatusARM, MaintenanceWindowStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMaintenanceWindowStatusARM runs a test to see if a specific instance of MaintenanceWindow_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMaintenanceWindowStatusARM(subject MaintenanceWindow_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MaintenanceWindow_StatusARM
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

// Generator of MaintenanceWindow_StatusARM instances for property testing - lazily instantiated by
//MaintenanceWindowStatusARMGenerator()
var maintenanceWindowStatusARMGenerator gopter.Gen

// MaintenanceWindowStatusARMGenerator returns a generator of MaintenanceWindow_StatusARM instances for property testing.
func MaintenanceWindowStatusARMGenerator() gopter.Gen {
	if maintenanceWindowStatusARMGenerator != nil {
		return maintenanceWindowStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMaintenanceWindowStatusARM(generators)
	maintenanceWindowStatusARMGenerator = gen.Struct(reflect.TypeOf(MaintenanceWindow_StatusARM{}), generators)

	return maintenanceWindowStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForMaintenanceWindowStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMaintenanceWindowStatusARM(gens map[string]gopter.Gen) {
	gens["CustomWindow"] = gen.PtrOf(gen.AlphaString())
	gens["DayOfWeek"] = gen.PtrOf(gen.Int())
	gens["StartHour"] = gen.PtrOf(gen.Int())
	gens["StartMinute"] = gen.PtrOf(gen.Int())
}

func Test_Network_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Network_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkStatusARM, NetworkStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkStatusARM runs a test to see if a specific instance of Network_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkStatusARM(subject Network_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Network_StatusARM
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

// Generator of Network_StatusARM instances for property testing - lazily instantiated by NetworkStatusARMGenerator()
var networkStatusARMGenerator gopter.Gen

// NetworkStatusARMGenerator returns a generator of Network_StatusARM instances for property testing.
func NetworkStatusARMGenerator() gopter.Gen {
	if networkStatusARMGenerator != nil {
		return networkStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkStatusARM(generators)
	networkStatusARMGenerator = gen.Struct(reflect.TypeOf(Network_StatusARM{}), generators)

	return networkStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkStatusARM(gens map[string]gopter.Gen) {
	gens["DelegatedSubnetResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateDnsZoneArmResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(NetworkStatusPublicNetworkAccessDisabled, NetworkStatusPublicNetworkAccessEnabled))
}

func Test_Storage_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Storage_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageStatusARM, StorageStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageStatusARM runs a test to see if a specific instance of Storage_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageStatusARM(subject Storage_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Storage_StatusARM
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

// Generator of Storage_StatusARM instances for property testing - lazily instantiated by StorageStatusARMGenerator()
var storageStatusARMGenerator gopter.Gen

// StorageStatusARMGenerator returns a generator of Storage_StatusARM instances for property testing.
func StorageStatusARMGenerator() gopter.Gen {
	if storageStatusARMGenerator != nil {
		return storageStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageStatusARM(generators)
	storageStatusARMGenerator = gen.Struct(reflect.TypeOf(Storage_StatusARM{}), generators)

	return storageStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForStorageStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageStatusARM(gens map[string]gopter.Gen) {
	gens["StorageSizeGB"] = gen.PtrOf(gen.Int())
}
