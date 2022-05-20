// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210901

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of Registries_Spec. Use v1beta20210901.Registries_Spec instead
type Registries_SpecARM struct {
	Identity   *IdentityPropertiesARM `json:"identity,omitempty"`
	Location   *string                `json:"location,omitempty"`
	Name       string                 `json:"name,omitempty"`
	Properties *RegistryPropertiesARM `json:"properties,omitempty"`
	Sku        *SkuARM                `json:"sku,omitempty"`
	Tags       map[string]string      `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Registries_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-09-01"
func (registries Registries_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (registries *Registries_SpecARM) GetName() string {
	return registries.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ContainerRegistry/registries"
func (registries *Registries_SpecARM) GetType() string {
	return "Microsoft.ContainerRegistry/registries"
}

// Deprecated version of IdentityProperties. Use v1beta20210901.IdentityProperties instead
type IdentityPropertiesARM struct {
	PrincipalId            *string                              `json:"principalId,omitempty"`
	TenantId               *string                              `json:"tenantId,omitempty"`
	Type                   *IdentityPropertiesType              `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserIdentityPropertiesARM `json:"userAssignedIdentities,omitempty"`
}

// Deprecated version of RegistryProperties. Use v1beta20210901.RegistryProperties instead
type RegistryPropertiesARM struct {
	AdminUserEnabled         *bool                                       `json:"adminUserEnabled,omitempty"`
	DataEndpointEnabled      *bool                                       `json:"dataEndpointEnabled,omitempty"`
	Encryption               *EncryptionPropertyARM                      `json:"encryption,omitempty"`
	NetworkRuleBypassOptions *RegistryPropertiesNetworkRuleBypassOptions `json:"networkRuleBypassOptions,omitempty"`
	NetworkRuleSet           *NetworkRuleSetARM                          `json:"networkRuleSet,omitempty"`
	Policies                 *PoliciesARM                                `json:"policies,omitempty"`
	PublicNetworkAccess      *RegistryPropertiesPublicNetworkAccess      `json:"publicNetworkAccess,omitempty"`
	ZoneRedundancy           *RegistryPropertiesZoneRedundancy           `json:"zoneRedundancy,omitempty"`
}

// Deprecated version of Sku. Use v1beta20210901.Sku instead
type SkuARM struct {
	Name *SkuName `json:"name,omitempty"`
}

// Deprecated version of EncryptionProperty. Use v1beta20210901.EncryptionProperty instead
type EncryptionPropertyARM struct {
	KeyVaultProperties *KeyVaultPropertiesARM    `json:"keyVaultProperties,omitempty"`
	Status             *EncryptionPropertyStatus `json:"status,omitempty"`
}

// Deprecated version of IdentityPropertiesType. Use v1beta20210901.IdentityPropertiesType instead
// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned, UserAssigned","UserAssigned"}
type IdentityPropertiesType string

const (
	IdentityPropertiesTypeNone                       = IdentityPropertiesType("None")
	IdentityPropertiesTypeSystemAssigned             = IdentityPropertiesType("SystemAssigned")
	IdentityPropertiesTypeSystemAssignedUserAssigned = IdentityPropertiesType("SystemAssigned, UserAssigned")
	IdentityPropertiesTypeUserAssigned               = IdentityPropertiesType("UserAssigned")
)

// Deprecated version of NetworkRuleSet. Use v1beta20210901.NetworkRuleSet instead
type NetworkRuleSetARM struct {
	DefaultAction *NetworkRuleSetDefaultAction `json:"defaultAction,omitempty"`
	IpRules       []IPRuleARM                  `json:"ipRules,omitempty"`
}

// Deprecated version of Policies. Use v1beta20210901.Policies instead
type PoliciesARM struct {
	ExportPolicy     *ExportPolicyARM     `json:"exportPolicy,omitempty"`
	QuarantinePolicy *QuarantinePolicyARM `json:"quarantinePolicy,omitempty"`
	RetentionPolicy  *RetentionPolicyARM  `json:"retentionPolicy,omitempty"`
	TrustPolicy      *TrustPolicyARM      `json:"trustPolicy,omitempty"`
}

// Deprecated version of RegistryPropertiesNetworkRuleBypassOptions. Use
// v1beta20210901.RegistryPropertiesNetworkRuleBypassOptions instead
// +kubebuilder:validation:Enum={"AzureServices","None"}
type RegistryPropertiesNetworkRuleBypassOptions string

const (
	RegistryPropertiesNetworkRuleBypassOptionsAzureServices = RegistryPropertiesNetworkRuleBypassOptions("AzureServices")
	RegistryPropertiesNetworkRuleBypassOptionsNone          = RegistryPropertiesNetworkRuleBypassOptions("None")
)

// Deprecated version of RegistryPropertiesPublicNetworkAccess. Use v1beta20210901.RegistryPropertiesPublicNetworkAccess
// instead
// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type RegistryPropertiesPublicNetworkAccess string

const (
	RegistryPropertiesPublicNetworkAccessDisabled = RegistryPropertiesPublicNetworkAccess("Disabled")
	RegistryPropertiesPublicNetworkAccessEnabled  = RegistryPropertiesPublicNetworkAccess("Enabled")
)

// Deprecated version of RegistryPropertiesZoneRedundancy. Use v1beta20210901.RegistryPropertiesZoneRedundancy instead
// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type RegistryPropertiesZoneRedundancy string

const (
	RegistryPropertiesZoneRedundancyDisabled = RegistryPropertiesZoneRedundancy("Disabled")
	RegistryPropertiesZoneRedundancyEnabled  = RegistryPropertiesZoneRedundancy("Enabled")
)

// Deprecated version of SkuName. Use v1beta20210901.SkuName instead
// +kubebuilder:validation:Enum={"Basic","Classic","Premium","Standard"}
type SkuName string

const (
	SkuNameBasic    = SkuName("Basic")
	SkuNameClassic  = SkuName("Classic")
	SkuNamePremium  = SkuName("Premium")
	SkuNameStandard = SkuName("Standard")
)

// Deprecated version of UserIdentityProperties. Use v1beta20210901.UserIdentityProperties instead
type UserIdentityPropertiesARM struct {
	ClientId    *string `json:"clientId,omitempty"`
	PrincipalId *string `json:"principalId,omitempty"`
}

// Deprecated version of EncryptionPropertyStatus. Use v1beta20210901.EncryptionPropertyStatus instead
// +kubebuilder:validation:Enum={"disabled","enabled"}
type EncryptionPropertyStatus string

const (
	EncryptionPropertyStatusDisabled = EncryptionPropertyStatus("disabled")
	EncryptionPropertyStatusEnabled  = EncryptionPropertyStatus("enabled")
)

// Deprecated version of ExportPolicy. Use v1beta20210901.ExportPolicy instead
type ExportPolicyARM struct {
	Status *ExportPolicyStatus `json:"status,omitempty"`
}

// Deprecated version of IPRule. Use v1beta20210901.IPRule instead
type IPRuleARM struct {
	Action *IPRuleAction `json:"action,omitempty"`
	Value  *string       `json:"value,omitempty"`
}

// Deprecated version of KeyVaultProperties. Use v1beta20210901.KeyVaultProperties instead
type KeyVaultPropertiesARM struct {
	Identity      *string `json:"identity,omitempty"`
	KeyIdentifier *string `json:"keyIdentifier,omitempty"`
}

// Deprecated version of NetworkRuleSetDefaultAction. Use v1beta20210901.NetworkRuleSetDefaultAction instead
// +kubebuilder:validation:Enum={"Allow","Deny"}
type NetworkRuleSetDefaultAction string

const (
	NetworkRuleSetDefaultActionAllow = NetworkRuleSetDefaultAction("Allow")
	NetworkRuleSetDefaultActionDeny  = NetworkRuleSetDefaultAction("Deny")
)

// Deprecated version of QuarantinePolicy. Use v1beta20210901.QuarantinePolicy instead
type QuarantinePolicyARM struct {
	Status *QuarantinePolicyStatus `json:"status,omitempty"`
}

// Deprecated version of RetentionPolicy. Use v1beta20210901.RetentionPolicy instead
type RetentionPolicyARM struct {
	Days   *int                   `json:"days,omitempty"`
	Status *RetentionPolicyStatus `json:"status,omitempty"`
}

// Deprecated version of TrustPolicy. Use v1beta20210901.TrustPolicy instead
type TrustPolicyARM struct {
	Status *TrustPolicyStatus `json:"status,omitempty"`
	Type   *TrustPolicyType   `json:"type,omitempty"`
}

// Deprecated version of ExportPolicyStatus. Use v1beta20210901.ExportPolicyStatus instead
// +kubebuilder:validation:Enum={"disabled","enabled"}
type ExportPolicyStatus string

const (
	ExportPolicyStatusDisabled = ExportPolicyStatus("disabled")
	ExportPolicyStatusEnabled  = ExportPolicyStatus("enabled")
)

// Deprecated version of IPRuleAction. Use v1beta20210901.IPRuleAction instead
// +kubebuilder:validation:Enum={"Allow"}
type IPRuleAction string

const IPRuleActionAllow = IPRuleAction("Allow")

// Deprecated version of QuarantinePolicyStatus. Use v1beta20210901.QuarantinePolicyStatus instead
// +kubebuilder:validation:Enum={"disabled","enabled"}
type QuarantinePolicyStatus string

const (
	QuarantinePolicyStatusDisabled = QuarantinePolicyStatus("disabled")
	QuarantinePolicyStatusEnabled  = QuarantinePolicyStatus("enabled")
)

// Deprecated version of RetentionPolicyStatus. Use v1beta20210901.RetentionPolicyStatus instead
// +kubebuilder:validation:Enum={"disabled","enabled"}
type RetentionPolicyStatus string

const (
	RetentionPolicyStatusDisabled = RetentionPolicyStatus("disabled")
	RetentionPolicyStatusEnabled  = RetentionPolicyStatus("enabled")
)

// Deprecated version of TrustPolicyStatus. Use v1beta20210901.TrustPolicyStatus instead
// +kubebuilder:validation:Enum={"disabled","enabled"}
type TrustPolicyStatus string

const (
	TrustPolicyStatusDisabled = TrustPolicyStatus("disabled")
	TrustPolicyStatusEnabled  = TrustPolicyStatus("enabled")
)

// Deprecated version of TrustPolicyType. Use v1beta20210901.TrustPolicyType instead
// +kubebuilder:validation:Enum={"Notary"}
type TrustPolicyType string

const TrustPolicyTypeNotary = TrustPolicyType("Notary")