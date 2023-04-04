// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220120preview

import (
	"fmt"
	v1api20220120ps "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20220120previewstorage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Generator information:
// - Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2022-01-20-preview/Databases.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/databases/{databaseName}
type FlexibleServersDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServers_Database_Spec   `json:"spec,omitempty"`
	Status            FlexibleServers_Database_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &FlexibleServersDatabase{}

// GetConditions returns the conditions of the resource
func (database *FlexibleServersDatabase) GetConditions() conditions.Conditions {
	return database.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (database *FlexibleServersDatabase) SetConditions(conditions conditions.Conditions) {
	database.Status.Conditions = conditions
}

var _ conversion.Convertible = &FlexibleServersDatabase{}

// ConvertFrom populates our FlexibleServersDatabase from the provided hub FlexibleServersDatabase
func (database *FlexibleServersDatabase) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source v1api20220120ps.FlexibleServersDatabase

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = database.AssignProperties_From_FlexibleServersDatabase(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to database")
	}

	return nil
}

// ConvertTo populates the provided hub FlexibleServersDatabase from our FlexibleServersDatabase
func (database *FlexibleServersDatabase) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination v1api20220120ps.FlexibleServersDatabase
	err := database.AssignProperties_To_FlexibleServersDatabase(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from database")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-dbforpostgresql-azure-com-v1api20220120preview-flexibleserversdatabase,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbforpostgresql.azure.com,resources=flexibleserversdatabases,verbs=create;update,versions=v1api20220120preview,name=default.v1api20220120preview.flexibleserversdatabases.dbforpostgresql.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &FlexibleServersDatabase{}

// Default applies defaults to the FlexibleServersDatabase resource
func (database *FlexibleServersDatabase) Default() {
	database.defaultImpl()
	var temp any = database
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (database *FlexibleServersDatabase) defaultAzureName() {
	if database.Spec.AzureName == "" {
		database.Spec.AzureName = database.Name
	}
}

// defaultImpl applies the code generated defaults to the FlexibleServersDatabase resource
func (database *FlexibleServersDatabase) defaultImpl() { database.defaultAzureName() }

var _ genruntime.KubernetesResource = &FlexibleServersDatabase{}

// AzureName returns the Azure name of the resource
func (database *FlexibleServersDatabase) AzureName() string {
	return database.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-01-20-preview"
func (database FlexibleServersDatabase) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (database *FlexibleServersDatabase) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (database *FlexibleServersDatabase) GetSpec() genruntime.ConvertibleSpec {
	return &database.Spec
}

// GetStatus returns the status of this resource
func (database *FlexibleServersDatabase) GetStatus() genruntime.ConvertibleStatus {
	return &database.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers/databases"
func (database *FlexibleServersDatabase) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers/databases"
}

// NewEmptyStatus returns a new empty (blank) status
func (database *FlexibleServersDatabase) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &FlexibleServers_Database_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (database *FlexibleServersDatabase) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(database.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  database.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (database *FlexibleServersDatabase) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*FlexibleServers_Database_STATUS); ok {
		database.Status = *st
		return nil
	}

	// Convert status to required version
	var st FlexibleServers_Database_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	database.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-dbforpostgresql-azure-com-v1api20220120preview-flexibleserversdatabase,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbforpostgresql.azure.com,resources=flexibleserversdatabases,verbs=create;update,versions=v1api20220120preview,name=validate.v1api20220120preview.flexibleserversdatabases.dbforpostgresql.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &FlexibleServersDatabase{}

// ValidateCreate validates the creation of the resource
func (database *FlexibleServersDatabase) ValidateCreate() error {
	validations := database.createValidations()
	var temp any = database
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateDelete validates the deletion of the resource
func (database *FlexibleServersDatabase) ValidateDelete() error {
	validations := database.deleteValidations()
	var temp any = database
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateUpdate validates an update of the resource
func (database *FlexibleServersDatabase) ValidateUpdate(old runtime.Object) error {
	validations := database.updateValidations()
	var temp any = database
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation(old)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// createValidations validates the creation of the resource
func (database *FlexibleServersDatabase) createValidations() []func() error {
	return []func() error{database.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (database *FlexibleServersDatabase) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (database *FlexibleServersDatabase) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return database.validateResourceReferences()
		},
		database.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (database *FlexibleServersDatabase) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&database.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (database *FlexibleServersDatabase) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*FlexibleServersDatabase)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, database)
}

// AssignProperties_From_FlexibleServersDatabase populates our FlexibleServersDatabase from the provided source FlexibleServersDatabase
func (database *FlexibleServersDatabase) AssignProperties_From_FlexibleServersDatabase(source *v1api20220120ps.FlexibleServersDatabase) error {

	// ObjectMeta
	database.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec FlexibleServers_Database_Spec
	err := spec.AssignProperties_From_FlexibleServers_Database_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_FlexibleServers_Database_Spec() to populate field Spec")
	}
	database.Spec = spec

	// Status
	var status FlexibleServers_Database_STATUS
	err = status.AssignProperties_From_FlexibleServers_Database_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_FlexibleServers_Database_STATUS() to populate field Status")
	}
	database.Status = status

	// No error
	return nil
}

// AssignProperties_To_FlexibleServersDatabase populates the provided destination FlexibleServersDatabase from our FlexibleServersDatabase
func (database *FlexibleServersDatabase) AssignProperties_To_FlexibleServersDatabase(destination *v1api20220120ps.FlexibleServersDatabase) error {

	// ObjectMeta
	destination.ObjectMeta = *database.ObjectMeta.DeepCopy()

	// Spec
	var spec v1api20220120ps.FlexibleServers_Database_Spec
	err := database.Spec.AssignProperties_To_FlexibleServers_Database_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_FlexibleServers_Database_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1api20220120ps.FlexibleServers_Database_STATUS
	err = database.Status.AssignProperties_To_FlexibleServers_Database_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_FlexibleServers_Database_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (database *FlexibleServersDatabase) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: database.Spec.OriginalVersion(),
		Kind:    "FlexibleServersDatabase",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2022-01-20-preview/Databases.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/databases/{databaseName}
type FlexibleServersDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServersDatabase `json:"items"`
}

type FlexibleServers_Database_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// Charset: The charset of the database.
	Charset *string `json:"charset,omitempty"`

	// Collation: The collation of the database.
	Collation *string `json:"collation,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a dbforpostgresql.azure.com/FlexibleServer resource
	Owner *genruntime.KnownResourceReference `group:"dbforpostgresql.azure.com" json:"owner,omitempty" kind:"FlexibleServer"`
}

var _ genruntime.ARMTransformer = &FlexibleServers_Database_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (database *FlexibleServers_Database_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if database == nil {
		return nil, nil
	}
	result := &FlexibleServers_Database_Spec_ARM{}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if database.Charset != nil || database.Collation != nil {
		result.Properties = &DatabaseProperties_ARM{}
	}
	if database.Charset != nil {
		charset := *database.Charset
		result.Properties.Charset = &charset
	}
	if database.Collation != nil {
		collation := *database.Collation
		result.Properties.Collation = &collation
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (database *FlexibleServers_Database_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &FlexibleServers_Database_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (database *FlexibleServers_Database_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(FlexibleServers_Database_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected FlexibleServers_Database_Spec_ARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	database.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Charset’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Charset != nil {
			charset := *typedInput.Properties.Charset
			database.Charset = &charset
		}
	}

	// Set property ‘Collation’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Collation != nil {
			collation := *typedInput.Properties.Collation
			database.Collation = &collation
		}
	}

	// Set property ‘Owner’:
	database.Owner = &genruntime.KnownResourceReference{Name: owner.Name}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &FlexibleServers_Database_Spec{}

// ConvertSpecFrom populates our FlexibleServers_Database_Spec from the provided source
func (database *FlexibleServers_Database_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1api20220120ps.FlexibleServers_Database_Spec)
	if ok {
		// Populate our instance from source
		return database.AssignProperties_From_FlexibleServers_Database_Spec(src)
	}

	// Convert to an intermediate form
	src = &v1api20220120ps.FlexibleServers_Database_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = database.AssignProperties_From_FlexibleServers_Database_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our FlexibleServers_Database_Spec
func (database *FlexibleServers_Database_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1api20220120ps.FlexibleServers_Database_Spec)
	if ok {
		// Populate destination from our instance
		return database.AssignProperties_To_FlexibleServers_Database_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v1api20220120ps.FlexibleServers_Database_Spec{}
	err := database.AssignProperties_To_FlexibleServers_Database_Spec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_FlexibleServers_Database_Spec populates our FlexibleServers_Database_Spec from the provided source FlexibleServers_Database_Spec
func (database *FlexibleServers_Database_Spec) AssignProperties_From_FlexibleServers_Database_Spec(source *v1api20220120ps.FlexibleServers_Database_Spec) error {

	// AzureName
	database.AzureName = source.AzureName

	// Charset
	database.Charset = genruntime.ClonePointerToString(source.Charset)

	// Collation
	database.Collation = genruntime.ClonePointerToString(source.Collation)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		database.Owner = &owner
	} else {
		database.Owner = nil
	}

	// No error
	return nil
}

// AssignProperties_To_FlexibleServers_Database_Spec populates the provided destination FlexibleServers_Database_Spec from our FlexibleServers_Database_Spec
func (database *FlexibleServers_Database_Spec) AssignProperties_To_FlexibleServers_Database_Spec(destination *v1api20220120ps.FlexibleServers_Database_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = database.AzureName

	// Charset
	destination.Charset = genruntime.ClonePointerToString(database.Charset)

	// Collation
	destination.Collation = genruntime.ClonePointerToString(database.Collation)

	// OriginalVersion
	destination.OriginalVersion = database.OriginalVersion()

	// Owner
	if database.Owner != nil {
		owner := database.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (database *FlexibleServers_Database_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (database *FlexibleServers_Database_Spec) SetAzureName(azureName string) {
	database.AzureName = azureName
}

type FlexibleServers_Database_STATUS struct {
	// Charset: The charset of the database.
	Charset *string `json:"charset,omitempty"`

	// Collation: The collation of the database.
	Collation *string `json:"collation,omitempty"`

	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// SystemData: The system metadata relating to this resource.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &FlexibleServers_Database_STATUS{}

// ConvertStatusFrom populates our FlexibleServers_Database_STATUS from the provided source
func (database *FlexibleServers_Database_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1api20220120ps.FlexibleServers_Database_STATUS)
	if ok {
		// Populate our instance from source
		return database.AssignProperties_From_FlexibleServers_Database_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v1api20220120ps.FlexibleServers_Database_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = database.AssignProperties_From_FlexibleServers_Database_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our FlexibleServers_Database_STATUS
func (database *FlexibleServers_Database_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1api20220120ps.FlexibleServers_Database_STATUS)
	if ok {
		// Populate destination from our instance
		return database.AssignProperties_To_FlexibleServers_Database_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v1api20220120ps.FlexibleServers_Database_STATUS{}
	err := database.AssignProperties_To_FlexibleServers_Database_STATUS(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

var _ genruntime.FromARMConverter = &FlexibleServers_Database_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (database *FlexibleServers_Database_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &FlexibleServers_Database_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (database *FlexibleServers_Database_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(FlexibleServers_Database_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected FlexibleServers_Database_STATUS_ARM, got %T", armInput)
	}

	// Set property ‘Charset’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Charset != nil {
			charset := *typedInput.Properties.Charset
			database.Charset = &charset
		}
	}

	// Set property ‘Collation’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Collation != nil {
			collation := *typedInput.Properties.Collation
			database.Collation = &collation
		}
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		database.Id = &id
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		database.Name = &name
	}

	// Set property ‘SystemData’:
	if typedInput.SystemData != nil {
		var systemData1 SystemData_STATUS
		err := systemData1.PopulateFromARM(owner, *typedInput.SystemData)
		if err != nil {
			return err
		}
		systemData := systemData1
		database.SystemData = &systemData
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		database.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_FlexibleServers_Database_STATUS populates our FlexibleServers_Database_STATUS from the provided source FlexibleServers_Database_STATUS
func (database *FlexibleServers_Database_STATUS) AssignProperties_From_FlexibleServers_Database_STATUS(source *v1api20220120ps.FlexibleServers_Database_STATUS) error {

	// Charset
	database.Charset = genruntime.ClonePointerToString(source.Charset)

	// Collation
	database.Collation = genruntime.ClonePointerToString(source.Collation)

	// Conditions
	database.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	database.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	database.Name = genruntime.ClonePointerToString(source.Name)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
		}
		database.SystemData = &systemDatum
	} else {
		database.SystemData = nil
	}

	// Type
	database.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_FlexibleServers_Database_STATUS populates the provided destination FlexibleServers_Database_STATUS from our FlexibleServers_Database_STATUS
func (database *FlexibleServers_Database_STATUS) AssignProperties_To_FlexibleServers_Database_STATUS(destination *v1api20220120ps.FlexibleServers_Database_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Charset
	destination.Charset = genruntime.ClonePointerToString(database.Charset)

	// Collation
	destination.Collation = genruntime.ClonePointerToString(database.Collation)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(database.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(database.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(database.Name)

	// SystemData
	if database.SystemData != nil {
		var systemDatum v1api20220120ps.SystemData_STATUS
		err := database.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(database.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&FlexibleServersDatabase{}, &FlexibleServersDatabaseList{})
}
