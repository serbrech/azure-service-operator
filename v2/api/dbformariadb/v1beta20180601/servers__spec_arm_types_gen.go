// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180601

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type Servers_SpecARM struct {
	// Location: The location the resource resides in.
	Location *string `json:"location,omitempty"`

	// Name: The name of the server.
	Name string `json:"name,omitempty"`

	// Properties: The properties used to create a new server.
	Properties *ServerPropertiesForCreateARM `json:"properties,omitempty"`

	// Sku: Billing information related properties of a server.
	Sku *SkuARM `json:"sku,omitempty"`

	// Tags: Application-specific metadata in the form of key-value pairs.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Servers_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-06-01"
func (servers Servers_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (servers *Servers_SpecARM) GetName() string {
	return servers.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMariaDB/servers"
func (servers *Servers_SpecARM) GetType() string {
	return "Microsoft.DBforMariaDB/servers"
}

// Generated from: https://schema.management.azure.com/schemas/2018-06-01/Microsoft.DBforMariaDB.json#/definitions/ServerPropertiesForCreate
type ServerPropertiesForCreateARM struct {
	// ServerPropertiesForDefaultCreate: Mutually exclusive with all other properties
	ServerPropertiesForDefaultCreate *ServerPropertiesForDefaultCreateARM `json:"serverPropertiesForDefaultCreate,omitempty"`

	// ServerPropertiesForGeoRestore: Mutually exclusive with all other properties
	ServerPropertiesForGeoRestore *ServerPropertiesForGeoRestoreARM `json:"serverPropertiesForGeoRestore,omitempty"`

	// ServerPropertiesForReplica: Mutually exclusive with all other properties
	ServerPropertiesForReplica *ServerPropertiesForReplicaARM `json:"serverPropertiesForReplica,omitempty"`

	// ServerPropertiesForRestore: Mutually exclusive with all other properties
	ServerPropertiesForRestore *ServerPropertiesForRestoreARM `json:"serverPropertiesForRestore,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because ServerPropertiesForCreateARM represents a discriminated union (JSON OneOf)
func (create ServerPropertiesForCreateARM) MarshalJSON() ([]byte, error) {
	if create.ServerPropertiesForDefaultCreate != nil {
		return json.Marshal(create.ServerPropertiesForDefaultCreate)
	}
	if create.ServerPropertiesForGeoRestore != nil {
		return json.Marshal(create.ServerPropertiesForGeoRestore)
	}
	if create.ServerPropertiesForReplica != nil {
		return json.Marshal(create.ServerPropertiesForReplica)
	}
	if create.ServerPropertiesForRestore != nil {
		return json.Marshal(create.ServerPropertiesForRestore)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the ServerPropertiesForCreateARM
func (create *ServerPropertiesForCreateARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["createMode"]
	if discriminator == "Default" {
		create.ServerPropertiesForDefaultCreate = &ServerPropertiesForDefaultCreateARM{}
		return json.Unmarshal(data, create.ServerPropertiesForDefaultCreate)
	}
	if discriminator == "GeoRestore" {
		create.ServerPropertiesForGeoRestore = &ServerPropertiesForGeoRestoreARM{}
		return json.Unmarshal(data, create.ServerPropertiesForGeoRestore)
	}
	if discriminator == "PointInTimeRestore" {
		create.ServerPropertiesForRestore = &ServerPropertiesForRestoreARM{}
		return json.Unmarshal(data, create.ServerPropertiesForRestore)
	}
	if discriminator == "Replica" {
		create.ServerPropertiesForReplica = &ServerPropertiesForReplicaARM{}
		return json.Unmarshal(data, create.ServerPropertiesForReplica)
	}

	// No error
	return nil
}

// Generated from: https://schema.management.azure.com/schemas/2018-06-01/Microsoft.DBforMariaDB.json#/definitions/Sku
type SkuARM struct {
	// Capacity: The scale up/out capacity, representing server's compute units.
	Capacity *int `json:"capacity,omitempty"`

	// Family: The family of hardware.
	Family *string `json:"family,omitempty"`

	// Name: The name of the sku, typically, tier + family + cores, e.g. B_Gen4_1, GP_Gen5_8.
	Name *string `json:"name,omitempty"`

	// Size: The size code, to be interpreted by resource as appropriate.
	Size *string `json:"size,omitempty"`

	// Tier: The tier of the particular SKU, e.g. Basic.
	Tier *SkuTier `json:"tier,omitempty"`
}

type ServerPropertiesForDefaultCreateARM struct {
	// AdministratorLogin: The administrator's login name of a server. Can only be specified when the server is being created
	// (and is required for creation).
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	// AdministratorLoginPassword: The password of the administrator login.
	AdministratorLoginPassword string                                                              `json:"administratorLoginPassword,omitempty"`
	CreateMode                 ServerPropertiesForCreateServerPropertiesForDefaultCreateCreateMode `json:"createMode,omitempty"`

	// MinimalTlsVersion: Enforce a minimal Tls version for the server.
	MinimalTlsVersion *ServerPropertiesForCreateServerPropertiesForDefaultCreateMinimalTlsVersion `json:"minimalTlsVersion,omitempty"`

	// PublicNetworkAccess: Whether or not public network access is allowed for this server. Value is optional but if passed
	// in, must be 'Enabled' or 'Disabled'.
	PublicNetworkAccess *ServerPropertiesForCreateServerPropertiesForDefaultCreatePublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// SslEnforcement: Enable ssl enforcement or not when connect to server.
	SslEnforcement *ServerPropertiesForCreateServerPropertiesForDefaultCreateSslEnforcement `json:"sslEnforcement,omitempty"`

	// StorageProfile: Storage Profile properties of a server
	StorageProfile *StorageProfileARM `json:"storageProfile,omitempty"`

	// Version: Server version.
	Version *ServerPropertiesForCreateServerPropertiesForDefaultCreateVersion `json:"version,omitempty"`
}

type ServerPropertiesForGeoRestoreARM struct {
	CreateMode ServerPropertiesForCreateServerPropertiesForGeoRestoreCreateMode `json:"createMode,omitempty"`

	// MinimalTlsVersion: Enforce a minimal Tls version for the server.
	MinimalTlsVersion *ServerPropertiesForCreateServerPropertiesForGeoRestoreMinimalTlsVersion `json:"minimalTlsVersion,omitempty"`

	// PublicNetworkAccess: Whether or not public network access is allowed for this server. Value is optional but if passed
	// in, must be 'Enabled' or 'Disabled'.
	PublicNetworkAccess *ServerPropertiesForCreateServerPropertiesForGeoRestorePublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// SourceServerId: The source server id to restore from.
	SourceServerId *string `json:"sourceServerId,omitempty"`

	// SslEnforcement: Enable ssl enforcement or not when connect to server.
	SslEnforcement *ServerPropertiesForCreateServerPropertiesForGeoRestoreSslEnforcement `json:"sslEnforcement,omitempty"`

	// StorageProfile: Storage Profile properties of a server
	StorageProfile *StorageProfileARM `json:"storageProfile,omitempty"`

	// Version: Server version.
	Version *ServerPropertiesForCreateServerPropertiesForGeoRestoreVersion `json:"version,omitempty"`
}

type ServerPropertiesForReplicaARM struct {
	CreateMode ServerPropertiesForCreateServerPropertiesForReplicaCreateMode `json:"createMode,omitempty"`

	// MinimalTlsVersion: Enforce a minimal Tls version for the server.
	MinimalTlsVersion *ServerPropertiesForCreateServerPropertiesForReplicaMinimalTlsVersion `json:"minimalTlsVersion,omitempty"`

	// PublicNetworkAccess: Whether or not public network access is allowed for this server. Value is optional but if passed
	// in, must be 'Enabled' or 'Disabled'.
	PublicNetworkAccess *ServerPropertiesForCreateServerPropertiesForReplicaPublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// SourceServerId: The master server id to create replica from.
	SourceServerId *string `json:"sourceServerId,omitempty"`

	// SslEnforcement: Enable ssl enforcement or not when connect to server.
	SslEnforcement *ServerPropertiesForCreateServerPropertiesForReplicaSslEnforcement `json:"sslEnforcement,omitempty"`

	// StorageProfile: Storage Profile properties of a server
	StorageProfile *StorageProfileARM `json:"storageProfile,omitempty"`

	// Version: Server version.
	Version *ServerPropertiesForCreateServerPropertiesForReplicaVersion `json:"version,omitempty"`
}

type ServerPropertiesForRestoreARM struct {
	CreateMode ServerPropertiesForCreateServerPropertiesForRestoreCreateMode `json:"createMode,omitempty"`

	// MinimalTlsVersion: Enforce a minimal Tls version for the server.
	MinimalTlsVersion *ServerPropertiesForCreateServerPropertiesForRestoreMinimalTlsVersion `json:"minimalTlsVersion,omitempty"`

	// PublicNetworkAccess: Whether or not public network access is allowed for this server. Value is optional but if passed
	// in, must be 'Enabled' or 'Disabled'.
	PublicNetworkAccess *ServerPropertiesForCreateServerPropertiesForRestorePublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// RestorePointInTime: Restore point creation time (ISO8601 format), specifying the time to restore from.
	RestorePointInTime *string `json:"restorePointInTime,omitempty"`

	// SourceServerId: The source server id to restore from.
	SourceServerId *string `json:"sourceServerId,omitempty"`

	// SslEnforcement: Enable ssl enforcement or not when connect to server.
	SslEnforcement *ServerPropertiesForCreateServerPropertiesForRestoreSslEnforcement `json:"sslEnforcement,omitempty"`

	// StorageProfile: Storage Profile properties of a server
	StorageProfile *StorageProfileARM `json:"storageProfile,omitempty"`

	// Version: Server version.
	Version *ServerPropertiesForCreateServerPropertiesForRestoreVersion `json:"version,omitempty"`
}

// +kubebuilder:validation:Enum={"Basic","GeneralPurpose","MemoryOptimized"}
type SkuTier string

const (
	SkuTierBasic           = SkuTier("Basic")
	SkuTierGeneralPurpose  = SkuTier("GeneralPurpose")
	SkuTierMemoryOptimized = SkuTier("MemoryOptimized")
)

// +kubebuilder:validation:Enum={"Default"}
type ServerPropertiesForCreateServerPropertiesForDefaultCreateCreateMode string

const ServerPropertiesForCreateServerPropertiesForDefaultCreateCreateModeDefault = ServerPropertiesForCreateServerPropertiesForDefaultCreateCreateMode("Default")

// +kubebuilder:validation:Enum={"TLS1_0","TLS1_1","TLS1_2","TLSEnforcementDisabled"}
type ServerPropertiesForCreateServerPropertiesForDefaultCreateMinimalTlsVersion string

const (
	ServerPropertiesForCreateServerPropertiesForDefaultCreateMinimalTlsVersionTLS10                  = ServerPropertiesForCreateServerPropertiesForDefaultCreateMinimalTlsVersion("TLS1_0")
	ServerPropertiesForCreateServerPropertiesForDefaultCreateMinimalTlsVersionTLS11                  = ServerPropertiesForCreateServerPropertiesForDefaultCreateMinimalTlsVersion("TLS1_1")
	ServerPropertiesForCreateServerPropertiesForDefaultCreateMinimalTlsVersionTLS12                  = ServerPropertiesForCreateServerPropertiesForDefaultCreateMinimalTlsVersion("TLS1_2")
	ServerPropertiesForCreateServerPropertiesForDefaultCreateMinimalTlsVersionTLSEnforcementDisabled = ServerPropertiesForCreateServerPropertiesForDefaultCreateMinimalTlsVersion("TLSEnforcementDisabled")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ServerPropertiesForCreateServerPropertiesForDefaultCreatePublicNetworkAccess string

const (
	ServerPropertiesForCreateServerPropertiesForDefaultCreatePublicNetworkAccessDisabled = ServerPropertiesForCreateServerPropertiesForDefaultCreatePublicNetworkAccess("Disabled")
	ServerPropertiesForCreateServerPropertiesForDefaultCreatePublicNetworkAccessEnabled  = ServerPropertiesForCreateServerPropertiesForDefaultCreatePublicNetworkAccess("Enabled")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ServerPropertiesForCreateServerPropertiesForDefaultCreateSslEnforcement string

const (
	ServerPropertiesForCreateServerPropertiesForDefaultCreateSslEnforcementDisabled = ServerPropertiesForCreateServerPropertiesForDefaultCreateSslEnforcement("Disabled")
	ServerPropertiesForCreateServerPropertiesForDefaultCreateSslEnforcementEnabled  = ServerPropertiesForCreateServerPropertiesForDefaultCreateSslEnforcement("Enabled")
)

// +kubebuilder:validation:Enum={"10.2","10.3"}
type ServerPropertiesForCreateServerPropertiesForDefaultCreateVersion string

const (
	ServerPropertiesForCreateServerPropertiesForDefaultCreateVersion102 = ServerPropertiesForCreateServerPropertiesForDefaultCreateVersion("10.2")
	ServerPropertiesForCreateServerPropertiesForDefaultCreateVersion103 = ServerPropertiesForCreateServerPropertiesForDefaultCreateVersion("10.3")
)

// +kubebuilder:validation:Enum={"GeoRestore"}
type ServerPropertiesForCreateServerPropertiesForGeoRestoreCreateMode string

const ServerPropertiesForCreateServerPropertiesForGeoRestoreCreateModeGeoRestore = ServerPropertiesForCreateServerPropertiesForGeoRestoreCreateMode("GeoRestore")

// +kubebuilder:validation:Enum={"TLS1_0","TLS1_1","TLS1_2","TLSEnforcementDisabled"}
type ServerPropertiesForCreateServerPropertiesForGeoRestoreMinimalTlsVersion string

const (
	ServerPropertiesForCreateServerPropertiesForGeoRestoreMinimalTlsVersionTLS10                  = ServerPropertiesForCreateServerPropertiesForGeoRestoreMinimalTlsVersion("TLS1_0")
	ServerPropertiesForCreateServerPropertiesForGeoRestoreMinimalTlsVersionTLS11                  = ServerPropertiesForCreateServerPropertiesForGeoRestoreMinimalTlsVersion("TLS1_1")
	ServerPropertiesForCreateServerPropertiesForGeoRestoreMinimalTlsVersionTLS12                  = ServerPropertiesForCreateServerPropertiesForGeoRestoreMinimalTlsVersion("TLS1_2")
	ServerPropertiesForCreateServerPropertiesForGeoRestoreMinimalTlsVersionTLSEnforcementDisabled = ServerPropertiesForCreateServerPropertiesForGeoRestoreMinimalTlsVersion("TLSEnforcementDisabled")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ServerPropertiesForCreateServerPropertiesForGeoRestorePublicNetworkAccess string

const (
	ServerPropertiesForCreateServerPropertiesForGeoRestorePublicNetworkAccessDisabled = ServerPropertiesForCreateServerPropertiesForGeoRestorePublicNetworkAccess("Disabled")
	ServerPropertiesForCreateServerPropertiesForGeoRestorePublicNetworkAccessEnabled  = ServerPropertiesForCreateServerPropertiesForGeoRestorePublicNetworkAccess("Enabled")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ServerPropertiesForCreateServerPropertiesForGeoRestoreSslEnforcement string

const (
	ServerPropertiesForCreateServerPropertiesForGeoRestoreSslEnforcementDisabled = ServerPropertiesForCreateServerPropertiesForGeoRestoreSslEnforcement("Disabled")
	ServerPropertiesForCreateServerPropertiesForGeoRestoreSslEnforcementEnabled  = ServerPropertiesForCreateServerPropertiesForGeoRestoreSslEnforcement("Enabled")
)

// +kubebuilder:validation:Enum={"10.2","10.3"}
type ServerPropertiesForCreateServerPropertiesForGeoRestoreVersion string

const (
	ServerPropertiesForCreateServerPropertiesForGeoRestoreVersion102 = ServerPropertiesForCreateServerPropertiesForGeoRestoreVersion("10.2")
	ServerPropertiesForCreateServerPropertiesForGeoRestoreVersion103 = ServerPropertiesForCreateServerPropertiesForGeoRestoreVersion("10.3")
)

// +kubebuilder:validation:Enum={"Replica"}
type ServerPropertiesForCreateServerPropertiesForReplicaCreateMode string

const ServerPropertiesForCreateServerPropertiesForReplicaCreateModeReplica = ServerPropertiesForCreateServerPropertiesForReplicaCreateMode("Replica")

// +kubebuilder:validation:Enum={"TLS1_0","TLS1_1","TLS1_2","TLSEnforcementDisabled"}
type ServerPropertiesForCreateServerPropertiesForReplicaMinimalTlsVersion string

const (
	ServerPropertiesForCreateServerPropertiesForReplicaMinimalTlsVersionTLS10                  = ServerPropertiesForCreateServerPropertiesForReplicaMinimalTlsVersion("TLS1_0")
	ServerPropertiesForCreateServerPropertiesForReplicaMinimalTlsVersionTLS11                  = ServerPropertiesForCreateServerPropertiesForReplicaMinimalTlsVersion("TLS1_1")
	ServerPropertiesForCreateServerPropertiesForReplicaMinimalTlsVersionTLS12                  = ServerPropertiesForCreateServerPropertiesForReplicaMinimalTlsVersion("TLS1_2")
	ServerPropertiesForCreateServerPropertiesForReplicaMinimalTlsVersionTLSEnforcementDisabled = ServerPropertiesForCreateServerPropertiesForReplicaMinimalTlsVersion("TLSEnforcementDisabled")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ServerPropertiesForCreateServerPropertiesForReplicaPublicNetworkAccess string

const (
	ServerPropertiesForCreateServerPropertiesForReplicaPublicNetworkAccessDisabled = ServerPropertiesForCreateServerPropertiesForReplicaPublicNetworkAccess("Disabled")
	ServerPropertiesForCreateServerPropertiesForReplicaPublicNetworkAccessEnabled  = ServerPropertiesForCreateServerPropertiesForReplicaPublicNetworkAccess("Enabled")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ServerPropertiesForCreateServerPropertiesForReplicaSslEnforcement string

const (
	ServerPropertiesForCreateServerPropertiesForReplicaSslEnforcementDisabled = ServerPropertiesForCreateServerPropertiesForReplicaSslEnforcement("Disabled")
	ServerPropertiesForCreateServerPropertiesForReplicaSslEnforcementEnabled  = ServerPropertiesForCreateServerPropertiesForReplicaSslEnforcement("Enabled")
)

// +kubebuilder:validation:Enum={"10.2","10.3"}
type ServerPropertiesForCreateServerPropertiesForReplicaVersion string

const (
	ServerPropertiesForCreateServerPropertiesForReplicaVersion102 = ServerPropertiesForCreateServerPropertiesForReplicaVersion("10.2")
	ServerPropertiesForCreateServerPropertiesForReplicaVersion103 = ServerPropertiesForCreateServerPropertiesForReplicaVersion("10.3")
)

// +kubebuilder:validation:Enum={"PointInTimeRestore"}
type ServerPropertiesForCreateServerPropertiesForRestoreCreateMode string

const ServerPropertiesForCreateServerPropertiesForRestoreCreateModePointInTimeRestore = ServerPropertiesForCreateServerPropertiesForRestoreCreateMode("PointInTimeRestore")

// +kubebuilder:validation:Enum={"TLS1_0","TLS1_1","TLS1_2","TLSEnforcementDisabled"}
type ServerPropertiesForCreateServerPropertiesForRestoreMinimalTlsVersion string

const (
	ServerPropertiesForCreateServerPropertiesForRestoreMinimalTlsVersionTLS10                  = ServerPropertiesForCreateServerPropertiesForRestoreMinimalTlsVersion("TLS1_0")
	ServerPropertiesForCreateServerPropertiesForRestoreMinimalTlsVersionTLS11                  = ServerPropertiesForCreateServerPropertiesForRestoreMinimalTlsVersion("TLS1_1")
	ServerPropertiesForCreateServerPropertiesForRestoreMinimalTlsVersionTLS12                  = ServerPropertiesForCreateServerPropertiesForRestoreMinimalTlsVersion("TLS1_2")
	ServerPropertiesForCreateServerPropertiesForRestoreMinimalTlsVersionTLSEnforcementDisabled = ServerPropertiesForCreateServerPropertiesForRestoreMinimalTlsVersion("TLSEnforcementDisabled")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ServerPropertiesForCreateServerPropertiesForRestorePublicNetworkAccess string

const (
	ServerPropertiesForCreateServerPropertiesForRestorePublicNetworkAccessDisabled = ServerPropertiesForCreateServerPropertiesForRestorePublicNetworkAccess("Disabled")
	ServerPropertiesForCreateServerPropertiesForRestorePublicNetworkAccessEnabled  = ServerPropertiesForCreateServerPropertiesForRestorePublicNetworkAccess("Enabled")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ServerPropertiesForCreateServerPropertiesForRestoreSslEnforcement string

const (
	ServerPropertiesForCreateServerPropertiesForRestoreSslEnforcementDisabled = ServerPropertiesForCreateServerPropertiesForRestoreSslEnforcement("Disabled")
	ServerPropertiesForCreateServerPropertiesForRestoreSslEnforcementEnabled  = ServerPropertiesForCreateServerPropertiesForRestoreSslEnforcement("Enabled")
)

// +kubebuilder:validation:Enum={"10.2","10.3"}
type ServerPropertiesForCreateServerPropertiesForRestoreVersion string

const (
	ServerPropertiesForCreateServerPropertiesForRestoreVersion102 = ServerPropertiesForCreateServerPropertiesForRestoreVersion("10.2")
	ServerPropertiesForCreateServerPropertiesForRestoreVersion103 = ServerPropertiesForCreateServerPropertiesForRestoreVersion("10.3")
)

// Generated from: https://schema.management.azure.com/schemas/2018-06-01/Microsoft.DBforMariaDB.json#/definitions/StorageProfile
type StorageProfileARM struct {
	// BackupRetentionDays: Backup retention days for the server.
	BackupRetentionDays *int `json:"backupRetentionDays,omitempty"`

	// GeoRedundantBackup: Enable Geo-redundant or not for server backup.
	GeoRedundantBackup *StorageProfileGeoRedundantBackup `json:"geoRedundantBackup,omitempty"`

	// StorageAutogrow: Enable Storage Auto Grow.
	StorageAutogrow *StorageProfileStorageAutogrow `json:"storageAutogrow,omitempty"`

	// StorageMB: Max storage allowed for a server.
	StorageMB *int `json:"storageMB,omitempty"`
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type StorageProfileGeoRedundantBackup string

const (
	StorageProfileGeoRedundantBackupDisabled = StorageProfileGeoRedundantBackup("Disabled")
	StorageProfileGeoRedundantBackupEnabled  = StorageProfileGeoRedundantBackup("Enabled")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type StorageProfileStorageAutogrow string

const (
	StorageProfileStorageAutogrowDisabled = StorageProfileStorageAutogrow("Disabled")
	StorageProfileStorageAutogrowEnabled  = StorageProfileStorageAutogrow("Enabled")
)
