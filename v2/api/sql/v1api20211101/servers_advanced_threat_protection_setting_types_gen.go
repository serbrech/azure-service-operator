// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import (
	"fmt"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
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
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/ServerAdvancedThreatProtectionSettings.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/advancedThreatProtectionSettings/Default
type ServersAdvancedThreatProtectionSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Servers_AdvancedThreatProtectionSetting_Spec   `json:"spec,omitempty"`
	Status            Servers_AdvancedThreatProtectionSetting_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ServersAdvancedThreatProtectionSetting{}

// GetConditions returns the conditions of the resource
func (setting *ServersAdvancedThreatProtectionSetting) GetConditions() conditions.Conditions {
	return setting.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (setting *ServersAdvancedThreatProtectionSetting) SetConditions(conditions conditions.Conditions) {
	setting.Status.Conditions = conditions
}

var _ conversion.Convertible = &ServersAdvancedThreatProtectionSetting{}

// ConvertFrom populates our ServersAdvancedThreatProtectionSetting from the provided hub ServersAdvancedThreatProtectionSetting
func (setting *ServersAdvancedThreatProtectionSetting) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20211101s.ServersAdvancedThreatProtectionSetting)
	if !ok {
		return fmt.Errorf("expected sql/v1api20211101/storage/ServersAdvancedThreatProtectionSetting but received %T instead", hub)
	}

	return setting.AssignProperties_From_ServersAdvancedThreatProtectionSetting(source)
}

// ConvertTo populates the provided hub ServersAdvancedThreatProtectionSetting from our ServersAdvancedThreatProtectionSetting
func (setting *ServersAdvancedThreatProtectionSetting) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20211101s.ServersAdvancedThreatProtectionSetting)
	if !ok {
		return fmt.Errorf("expected sql/v1api20211101/storage/ServersAdvancedThreatProtectionSetting but received %T instead", hub)
	}

	return setting.AssignProperties_To_ServersAdvancedThreatProtectionSetting(destination)
}

// +kubebuilder:webhook:path=/mutate-sql-azure-com-v1api20211101-serversadvancedthreatprotectionsetting,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversadvancedthreatprotectionsettings,verbs=create;update,versions=v1api20211101,name=default.v1api20211101.serversadvancedthreatprotectionsettings.sql.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &ServersAdvancedThreatProtectionSetting{}

// Default applies defaults to the ServersAdvancedThreatProtectionSetting resource
func (setting *ServersAdvancedThreatProtectionSetting) Default() {
	setting.defaultImpl()
	var temp any = setting
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the ServersAdvancedThreatProtectionSetting resource
func (setting *ServersAdvancedThreatProtectionSetting) defaultImpl() {}

var _ genruntime.ImportableResource = &ServersAdvancedThreatProtectionSetting{}

// InitializeSpec initializes the spec for this resource from the given status
func (setting *ServersAdvancedThreatProtectionSetting) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*Servers_AdvancedThreatProtectionSetting_STATUS); ok {
		return setting.Spec.Initialize_From_Servers_AdvancedThreatProtectionSetting_STATUS(s)
	}

	return fmt.Errorf("expected Status of type Servers_AdvancedThreatProtectionSetting_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &ServersAdvancedThreatProtectionSetting{}

// AzureName returns the Azure name of the resource (always "Default")
func (setting *ServersAdvancedThreatProtectionSetting) AzureName() string {
	return "Default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (setting ServersAdvancedThreatProtectionSetting) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (setting *ServersAdvancedThreatProtectionSetting) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (setting *ServersAdvancedThreatProtectionSetting) GetSpec() genruntime.ConvertibleSpec {
	return &setting.Spec
}

// GetStatus returns the status of this resource
func (setting *ServersAdvancedThreatProtectionSetting) GetStatus() genruntime.ConvertibleStatus {
	return &setting.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/advancedThreatProtectionSettings"
func (setting *ServersAdvancedThreatProtectionSetting) GetType() string {
	return "Microsoft.Sql/servers/advancedThreatProtectionSettings"
}

// NewEmptyStatus returns a new empty (blank) status
func (setting *ServersAdvancedThreatProtectionSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Servers_AdvancedThreatProtectionSetting_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (setting *ServersAdvancedThreatProtectionSetting) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(setting.Spec)
	return setting.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (setting *ServersAdvancedThreatProtectionSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Servers_AdvancedThreatProtectionSetting_STATUS); ok {
		setting.Status = *st
		return nil
	}

	// Convert status to required version
	var st Servers_AdvancedThreatProtectionSetting_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	setting.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-sql-azure-com-v1api20211101-serversadvancedthreatprotectionsetting,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversadvancedthreatprotectionsettings,verbs=create;update,versions=v1api20211101,name=validate.v1api20211101.serversadvancedthreatprotectionsettings.sql.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &ServersAdvancedThreatProtectionSetting{}

// ValidateCreate validates the creation of the resource
func (setting *ServersAdvancedThreatProtectionSetting) ValidateCreate() (admission.Warnings, error) {
	validations := setting.createValidations()
	var temp any = setting
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (setting *ServersAdvancedThreatProtectionSetting) ValidateDelete() (admission.Warnings, error) {
	validations := setting.deleteValidations()
	var temp any = setting
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (setting *ServersAdvancedThreatProtectionSetting) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := setting.updateValidations()
	var temp any = setting
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (setting *ServersAdvancedThreatProtectionSetting) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){setting.validateResourceReferences, setting.validateOwnerReference}
}

// deleteValidations validates the deletion of the resource
func (setting *ServersAdvancedThreatProtectionSetting) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (setting *ServersAdvancedThreatProtectionSetting) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return setting.validateResourceReferences()
		},
		setting.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return setting.validateOwnerReference()
		},
	}
}

// validateOwnerReference validates the owner field
func (setting *ServersAdvancedThreatProtectionSetting) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(setting)
}

// validateResourceReferences validates all resource references
func (setting *ServersAdvancedThreatProtectionSetting) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&setting.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (setting *ServersAdvancedThreatProtectionSetting) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*ServersAdvancedThreatProtectionSetting)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, setting)
}

// AssignProperties_From_ServersAdvancedThreatProtectionSetting populates our ServersAdvancedThreatProtectionSetting from the provided source ServersAdvancedThreatProtectionSetting
func (setting *ServersAdvancedThreatProtectionSetting) AssignProperties_From_ServersAdvancedThreatProtectionSetting(source *v20211101s.ServersAdvancedThreatProtectionSetting) error {

	// ObjectMeta
	setting.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Servers_AdvancedThreatProtectionSetting_Spec
	err := spec.AssignProperties_From_Servers_AdvancedThreatProtectionSetting_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_AdvancedThreatProtectionSetting_Spec() to populate field Spec")
	}
	setting.Spec = spec

	// Status
	var status Servers_AdvancedThreatProtectionSetting_STATUS
	err = status.AssignProperties_From_Servers_AdvancedThreatProtectionSetting_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_AdvancedThreatProtectionSetting_STATUS() to populate field Status")
	}
	setting.Status = status

	// No error
	return nil
}

// AssignProperties_To_ServersAdvancedThreatProtectionSetting populates the provided destination ServersAdvancedThreatProtectionSetting from our ServersAdvancedThreatProtectionSetting
func (setting *ServersAdvancedThreatProtectionSetting) AssignProperties_To_ServersAdvancedThreatProtectionSetting(destination *v20211101s.ServersAdvancedThreatProtectionSetting) error {

	// ObjectMeta
	destination.ObjectMeta = *setting.ObjectMeta.DeepCopy()

	// Spec
	var spec v20211101s.Servers_AdvancedThreatProtectionSetting_Spec
	err := setting.Spec.AssignProperties_To_Servers_AdvancedThreatProtectionSetting_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_AdvancedThreatProtectionSetting_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20211101s.Servers_AdvancedThreatProtectionSetting_STATUS
	err = setting.Status.AssignProperties_To_Servers_AdvancedThreatProtectionSetting_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_AdvancedThreatProtectionSetting_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (setting *ServersAdvancedThreatProtectionSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: setting.Spec.OriginalVersion(),
		Kind:    "ServersAdvancedThreatProtectionSetting",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/ServerAdvancedThreatProtectionSettings.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/advancedThreatProtectionSettings/Default
type ServersAdvancedThreatProtectionSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServersAdvancedThreatProtectionSetting `json:"items"`
}

type Servers_AdvancedThreatProtectionSetting_Spec struct {
	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a sql.azure.com/Server resource
	Owner *genruntime.KnownResourceReference `group:"sql.azure.com" json:"owner,omitempty" kind:"Server"`

	// +kubebuilder:validation:Required
	// State: Specifies the state of the Advanced Threat Protection, whether it is enabled or disabled or a state has not been
	// applied yet on the specific database or server.
	State *AdvancedThreatProtectionProperties_State `json:"state,omitempty"`
}

var _ genruntime.ARMTransformer = &Servers_AdvancedThreatProtectionSetting_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (setting *Servers_AdvancedThreatProtectionSetting_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if setting == nil {
		return nil, nil
	}
	result := &Servers_AdvancedThreatProtectionSetting_Spec_ARM{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if setting.State != nil {
		result.Properties = &AdvancedThreatProtectionProperties_ARM{}
	}
	if setting.State != nil {
		state := *setting.State
		result.Properties.State = &state
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (setting *Servers_AdvancedThreatProtectionSetting_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_AdvancedThreatProtectionSetting_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (setting *Servers_AdvancedThreatProtectionSetting_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_AdvancedThreatProtectionSetting_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_AdvancedThreatProtectionSetting_Spec_ARM, got %T", armInput)
	}

	// Set property "Owner":
	setting.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "State":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.State != nil {
			state := *typedInput.Properties.State
			setting.State = &state
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Servers_AdvancedThreatProtectionSetting_Spec{}

// ConvertSpecFrom populates our Servers_AdvancedThreatProtectionSetting_Spec from the provided source
func (setting *Servers_AdvancedThreatProtectionSetting_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20211101s.Servers_AdvancedThreatProtectionSetting_Spec)
	if ok {
		// Populate our instance from source
		return setting.AssignProperties_From_Servers_AdvancedThreatProtectionSetting_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.Servers_AdvancedThreatProtectionSetting_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = setting.AssignProperties_From_Servers_AdvancedThreatProtectionSetting_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Servers_AdvancedThreatProtectionSetting_Spec
func (setting *Servers_AdvancedThreatProtectionSetting_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20211101s.Servers_AdvancedThreatProtectionSetting_Spec)
	if ok {
		// Populate destination from our instance
		return setting.AssignProperties_To_Servers_AdvancedThreatProtectionSetting_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.Servers_AdvancedThreatProtectionSetting_Spec{}
	err := setting.AssignProperties_To_Servers_AdvancedThreatProtectionSetting_Spec(dst)
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

// AssignProperties_From_Servers_AdvancedThreatProtectionSetting_Spec populates our Servers_AdvancedThreatProtectionSetting_Spec from the provided source Servers_AdvancedThreatProtectionSetting_Spec
func (setting *Servers_AdvancedThreatProtectionSetting_Spec) AssignProperties_From_Servers_AdvancedThreatProtectionSetting_Spec(source *v20211101s.Servers_AdvancedThreatProtectionSetting_Spec) error {

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		setting.Owner = &owner
	} else {
		setting.Owner = nil
	}

	// State
	if source.State != nil {
		state := AdvancedThreatProtectionProperties_State(*source.State)
		setting.State = &state
	} else {
		setting.State = nil
	}

	// No error
	return nil
}

// AssignProperties_To_Servers_AdvancedThreatProtectionSetting_Spec populates the provided destination Servers_AdvancedThreatProtectionSetting_Spec from our Servers_AdvancedThreatProtectionSetting_Spec
func (setting *Servers_AdvancedThreatProtectionSetting_Spec) AssignProperties_To_Servers_AdvancedThreatProtectionSetting_Spec(destination *v20211101s.Servers_AdvancedThreatProtectionSetting_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// OriginalVersion
	destination.OriginalVersion = setting.OriginalVersion()

	// Owner
	if setting.Owner != nil {
		owner := setting.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// State
	if setting.State != nil {
		state := string(*setting.State)
		destination.State = &state
	} else {
		destination.State = nil
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

// Initialize_From_Servers_AdvancedThreatProtectionSetting_STATUS populates our Servers_AdvancedThreatProtectionSetting_Spec from the provided source Servers_AdvancedThreatProtectionSetting_STATUS
func (setting *Servers_AdvancedThreatProtectionSetting_Spec) Initialize_From_Servers_AdvancedThreatProtectionSetting_STATUS(source *Servers_AdvancedThreatProtectionSetting_STATUS) error {

	// State
	if source.State != nil {
		state := AdvancedThreatProtectionProperties_State(*source.State)
		setting.State = &state
	} else {
		setting.State = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (setting *Servers_AdvancedThreatProtectionSetting_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

type Servers_AdvancedThreatProtectionSetting_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// CreationTime: Specifies the UTC creation time of the policy.
	CreationTime *string `json:"creationTime,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// State: Specifies the state of the Advanced Threat Protection, whether it is enabled or disabled or a state has not been
	// applied yet on the specific database or server.
	State *AdvancedThreatProtectionProperties_State_STATUS `json:"state,omitempty"`

	// SystemData: SystemData of AdvancedThreatProtectionResource.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Servers_AdvancedThreatProtectionSetting_STATUS{}

// ConvertStatusFrom populates our Servers_AdvancedThreatProtectionSetting_STATUS from the provided source
func (setting *Servers_AdvancedThreatProtectionSetting_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20211101s.Servers_AdvancedThreatProtectionSetting_STATUS)
	if ok {
		// Populate our instance from source
		return setting.AssignProperties_From_Servers_AdvancedThreatProtectionSetting_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.Servers_AdvancedThreatProtectionSetting_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = setting.AssignProperties_From_Servers_AdvancedThreatProtectionSetting_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Servers_AdvancedThreatProtectionSetting_STATUS
func (setting *Servers_AdvancedThreatProtectionSetting_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20211101s.Servers_AdvancedThreatProtectionSetting_STATUS)
	if ok {
		// Populate destination from our instance
		return setting.AssignProperties_To_Servers_AdvancedThreatProtectionSetting_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.Servers_AdvancedThreatProtectionSetting_STATUS{}
	err := setting.AssignProperties_To_Servers_AdvancedThreatProtectionSetting_STATUS(dst)
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

var _ genruntime.FromARMConverter = &Servers_AdvancedThreatProtectionSetting_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (setting *Servers_AdvancedThreatProtectionSetting_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_AdvancedThreatProtectionSetting_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (setting *Servers_AdvancedThreatProtectionSetting_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_AdvancedThreatProtectionSetting_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_AdvancedThreatProtectionSetting_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "CreationTime":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.CreationTime != nil {
			creationTime := *typedInput.Properties.CreationTime
			setting.CreationTime = &creationTime
		}
	}

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		setting.Id = &id
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		setting.Name = &name
	}

	// Set property "State":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.State != nil {
			state := *typedInput.Properties.State
			setting.State = &state
		}
	}

	// Set property "SystemData":
	if typedInput.SystemData != nil {
		var systemData1 SystemData_STATUS
		err := systemData1.PopulateFromARM(owner, *typedInput.SystemData)
		if err != nil {
			return err
		}
		systemData := systemData1
		setting.SystemData = &systemData
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		setting.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_Servers_AdvancedThreatProtectionSetting_STATUS populates our Servers_AdvancedThreatProtectionSetting_STATUS from the provided source Servers_AdvancedThreatProtectionSetting_STATUS
func (setting *Servers_AdvancedThreatProtectionSetting_STATUS) AssignProperties_From_Servers_AdvancedThreatProtectionSetting_STATUS(source *v20211101s.Servers_AdvancedThreatProtectionSetting_STATUS) error {

	// Conditions
	setting.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// CreationTime
	setting.CreationTime = genruntime.ClonePointerToString(source.CreationTime)

	// Id
	setting.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	setting.Name = genruntime.ClonePointerToString(source.Name)

	// State
	if source.State != nil {
		state := AdvancedThreatProtectionProperties_State_STATUS(*source.State)
		setting.State = &state
	} else {
		setting.State = nil
	}

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
		}
		setting.SystemData = &systemDatum
	} else {
		setting.SystemData = nil
	}

	// Type
	setting.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_Servers_AdvancedThreatProtectionSetting_STATUS populates the provided destination Servers_AdvancedThreatProtectionSetting_STATUS from our Servers_AdvancedThreatProtectionSetting_STATUS
func (setting *Servers_AdvancedThreatProtectionSetting_STATUS) AssignProperties_To_Servers_AdvancedThreatProtectionSetting_STATUS(destination *v20211101s.Servers_AdvancedThreatProtectionSetting_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(setting.Conditions)

	// CreationTime
	destination.CreationTime = genruntime.ClonePointerToString(setting.CreationTime)

	// Id
	destination.Id = genruntime.ClonePointerToString(setting.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(setting.Name)

	// State
	if setting.State != nil {
		state := string(*setting.State)
		destination.State = &state
	} else {
		destination.State = nil
	}

	// SystemData
	if setting.SystemData != nil {
		var systemDatum v20211101s.SystemData_STATUS
		err := setting.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(setting.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"Disabled","Enabled","New"}
type AdvancedThreatProtectionProperties_State string

const (
	AdvancedThreatProtectionProperties_State_Disabled = AdvancedThreatProtectionProperties_State("Disabled")
	AdvancedThreatProtectionProperties_State_Enabled  = AdvancedThreatProtectionProperties_State("Enabled")
	AdvancedThreatProtectionProperties_State_New      = AdvancedThreatProtectionProperties_State("New")
)

type AdvancedThreatProtectionProperties_State_STATUS string

const (
	AdvancedThreatProtectionProperties_State_STATUS_Disabled = AdvancedThreatProtectionProperties_State_STATUS("Disabled")
	AdvancedThreatProtectionProperties_State_STATUS_Enabled  = AdvancedThreatProtectionProperties_State_STATUS("Enabled")
	AdvancedThreatProtectionProperties_State_STATUS_New      = AdvancedThreatProtectionProperties_State_STATUS("New")
)

// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_CreatedByType_STATUS `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_LastModifiedByType_STATUS `json:"lastModifiedByType,omitempty"`
}

var _ genruntime.FromARMConverter = &SystemData_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (data *SystemData_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &SystemData_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (data *SystemData_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SystemData_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SystemData_STATUS_ARM, got %T", armInput)
	}

	// Set property "CreatedAt":
	if typedInput.CreatedAt != nil {
		createdAt := *typedInput.CreatedAt
		data.CreatedAt = &createdAt
	}

	// Set property "CreatedBy":
	if typedInput.CreatedBy != nil {
		createdBy := *typedInput.CreatedBy
		data.CreatedBy = &createdBy
	}

	// Set property "CreatedByType":
	if typedInput.CreatedByType != nil {
		createdByType := *typedInput.CreatedByType
		data.CreatedByType = &createdByType
	}

	// Set property "LastModifiedAt":
	if typedInput.LastModifiedAt != nil {
		lastModifiedAt := *typedInput.LastModifiedAt
		data.LastModifiedAt = &lastModifiedAt
	}

	// Set property "LastModifiedBy":
	if typedInput.LastModifiedBy != nil {
		lastModifiedBy := *typedInput.LastModifiedBy
		data.LastModifiedBy = &lastModifiedBy
	}

	// Set property "LastModifiedByType":
	if typedInput.LastModifiedByType != nil {
		lastModifiedByType := *typedInput.LastModifiedByType
		data.LastModifiedByType = &lastModifiedByType
	}

	// No error
	return nil
}

// AssignProperties_From_SystemData_STATUS populates our SystemData_STATUS from the provided source SystemData_STATUS
func (data *SystemData_STATUS) AssignProperties_From_SystemData_STATUS(source *v20211101s.SystemData_STATUS) error {

	// CreatedAt
	data.CreatedAt = genruntime.ClonePointerToString(source.CreatedAt)

	// CreatedBy
	data.CreatedBy = genruntime.ClonePointerToString(source.CreatedBy)

	// CreatedByType
	if source.CreatedByType != nil {
		createdByType := SystemData_CreatedByType_STATUS(*source.CreatedByType)
		data.CreatedByType = &createdByType
	} else {
		data.CreatedByType = nil
	}

	// LastModifiedAt
	data.LastModifiedAt = genruntime.ClonePointerToString(source.LastModifiedAt)

	// LastModifiedBy
	data.LastModifiedBy = genruntime.ClonePointerToString(source.LastModifiedBy)

	// LastModifiedByType
	if source.LastModifiedByType != nil {
		lastModifiedByType := SystemData_LastModifiedByType_STATUS(*source.LastModifiedByType)
		data.LastModifiedByType = &lastModifiedByType
	} else {
		data.LastModifiedByType = nil
	}

	// No error
	return nil
}

// AssignProperties_To_SystemData_STATUS populates the provided destination SystemData_STATUS from our SystemData_STATUS
func (data *SystemData_STATUS) AssignProperties_To_SystemData_STATUS(destination *v20211101s.SystemData_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// CreatedAt
	destination.CreatedAt = genruntime.ClonePointerToString(data.CreatedAt)

	// CreatedBy
	destination.CreatedBy = genruntime.ClonePointerToString(data.CreatedBy)

	// CreatedByType
	if data.CreatedByType != nil {
		createdByType := string(*data.CreatedByType)
		destination.CreatedByType = &createdByType
	} else {
		destination.CreatedByType = nil
	}

	// LastModifiedAt
	destination.LastModifiedAt = genruntime.ClonePointerToString(data.LastModifiedAt)

	// LastModifiedBy
	destination.LastModifiedBy = genruntime.ClonePointerToString(data.LastModifiedBy)

	// LastModifiedByType
	if data.LastModifiedByType != nil {
		lastModifiedByType := string(*data.LastModifiedByType)
		destination.LastModifiedByType = &lastModifiedByType
	} else {
		destination.LastModifiedByType = nil
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

func init() {
	SchemeBuilder.Register(&ServersAdvancedThreatProtectionSetting{}, &ServersAdvancedThreatProtectionSettingList{})
}
