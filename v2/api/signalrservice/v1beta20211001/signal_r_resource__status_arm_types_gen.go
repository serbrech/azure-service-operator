// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211001

type SignalRResource_StatusARM struct {
	// Id: Fully qualified resource Id for the resource.
	Id       *string                    `json:"id,omitempty"`
	Identity *ManagedIdentity_StatusARM `json:"identity,omitempty"`
	Kind     *ServiceKind_Status        `json:"kind,omitempty"`

	// Location: The GEO location of the resource. e.g. West US | East US | North Central US | South Central US.
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource.
	Name       *string                      `json:"name,omitempty"`
	Properties *SignalRProperties_StatusARM `json:"properties,omitempty"`
	Sku        *ResourceSku_StatusARM       `json:"sku,omitempty"`
	SystemData *SystemData_StatusARM        `json:"systemData,omitempty"`

	// Tags: Tags of the service which is a list of key value pairs that describe the resource.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
	Type *string `json:"type,omitempty"`
}

type ManagedIdentity_StatusARM struct {
	// PrincipalId: Get the principal id for the system assigned identity.
	// Only be used in response.
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: Get the tenant id for the system assigned identity.
	// Only be used in response
	TenantId *string                     `json:"tenantId,omitempty"`
	Type     *ManagedIdentityType_Status `json:"type,omitempty"`

	// UserAssignedIdentities: Get or set the user assigned identities
	UserAssignedIdentities map[string]UserAssignedIdentityProperty_StatusARM `json:"userAssignedIdentities,omitempty"`
}

type ResourceSku_StatusARM struct {
	// Capacity: Optional, integer. The unit count of the resource. 1 by default.
	// If present, following values are allowed:
	// Free: 1
	// Standard: 1,2,5,10,20,50,100
	Capacity *int `json:"capacity,omitempty"`

	// Family: Not used. Retained for future use.
	Family *string `json:"family,omitempty"`

	// Name: The name of the SKU. Required.
	// Allowed values: Standard_S1, Free_F1
	Name *string `json:"name,omitempty"`

	// Size: Not used. Retained for future use.
	Size *string                `json:"size,omitempty"`
	Tier *SignalRSkuTier_Status `json:"tier,omitempty"`
}

type ServiceKind_Status string

const (
	ServiceKind_StatusRawWebSockets = ServiceKind_Status("RawWebSockets")
	ServiceKind_StatusSignalR       = ServiceKind_Status("SignalR")
)

type SignalRProperties_StatusARM struct {
	Cors *SignalRCorsSettings_StatusARM `json:"cors,omitempty"`

	// DisableAadAuth: DisableLocalAuth
	// Enable or disable aad auth
	// When set as true, connection with AuthType=aad won't work.
	DisableAadAuth *bool `json:"disableAadAuth,omitempty"`

	// DisableLocalAuth: DisableLocalAuth
	// Enable or disable local auth with AccessKey
	// When set as true, connection with AccessKey=xxx won't work.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`

	// ExternalIP: The publicly accessible IP of the resource.
	ExternalIP *string `json:"externalIP,omitempty"`

	// Features: List of the featureFlags.
	// FeatureFlags that are not included in the parameters for the update operation will not be modified.
	// And the response will only include featureFlags that are explicitly set.
	// When a featureFlag is not explicitly set, its globally default value will be used
	// But keep in mind, the default value doesn't mean "false". It varies in terms of different FeatureFlags.
	Features []SignalRFeature_StatusARM `json:"features,omitempty"`

	// HostName: FQDN of the service instance.
	HostName *string `json:"hostName,omitempty"`

	// HostNamePrefix: Deprecated.
	HostNamePrefix *string                       `json:"hostNamePrefix,omitempty"`
	NetworkACLs    *SignalRNetworkACLs_StatusARM `json:"networkACLs,omitempty"`

	// PrivateEndpointConnections: Private endpoint connections to the resource.
	PrivateEndpointConnections []PrivateEndpointConnection_Status_SignalR_SubResourceEmbeddedARM `json:"privateEndpointConnections,omitempty"`
	ProvisioningState          *ProvisioningState_Status                                         `json:"provisioningState,omitempty"`

	// PublicNetworkAccess: Enable or disable public network access. Default to "Enabled".
	// When it's Enabled, network ACLs still apply.
	// When it's Disabled, public network access is always disabled no matter what you set in network ACLs.
	PublicNetworkAccess *string `json:"publicNetworkAccess,omitempty"`

	// PublicPort: The publicly accessible port of the resource which is designed for browser/client side usage.
	PublicPort               *int                                `json:"publicPort,omitempty"`
	ResourceLogConfiguration *ResourceLogConfiguration_StatusARM `json:"resourceLogConfiguration,omitempty"`

	// ServerPort: The publicly accessible port of the resource which is designed for customer server side usage.
	ServerPort *int `json:"serverPort,omitempty"`

	// SharedPrivateLinkResources: The list of shared private link resources.
	SharedPrivateLinkResources []SharedPrivateLinkResource_Status_SignalR_SubResourceEmbeddedARM `json:"sharedPrivateLinkResources,omitempty"`
	Tls                        *SignalRTlsSettings_StatusARM                                     `json:"tls,omitempty"`
	Upstream                   *ServerlessUpstreamSettings_StatusARM                             `json:"upstream,omitempty"`

	// Version: Version of the resource. Probably you need the same or higher version of client SDKs.
	Version *string `json:"version,omitempty"`
}

type SystemData_StatusARM struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemDataStatusCreatedByType `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemDataStatusLastModifiedByType `json:"lastModifiedByType,omitempty"`
}

type ManagedIdentityType_Status string

const (
	ManagedIdentityType_StatusNone           = ManagedIdentityType_Status("None")
	ManagedIdentityType_StatusSystemAssigned = ManagedIdentityType_Status("SystemAssigned")
	ManagedIdentityType_StatusUserAssigned   = ManagedIdentityType_Status("UserAssigned")
)

type PrivateEndpointConnection_Status_SignalR_SubResourceEmbeddedARM struct {
	// Id: Fully qualified resource Id for the resource.
	Id         *string               `json:"id,omitempty"`
	SystemData *SystemData_StatusARM `json:"systemData,omitempty"`
}

type ResourceLogConfiguration_StatusARM struct {
	// Categories: Gets or sets the list of category configurations.
	Categories []ResourceLogCategory_StatusARM `json:"categories,omitempty"`
}

type ServerlessUpstreamSettings_StatusARM struct {
	// Templates: Gets or sets the list of Upstream URL templates. Order matters, and the first matching template takes effects.
	Templates []UpstreamTemplate_StatusARM `json:"templates,omitempty"`
}

type SharedPrivateLinkResource_Status_SignalR_SubResourceEmbeddedARM struct {
	// Id: Fully qualified resource Id for the resource.
	Id         *string               `json:"id,omitempty"`
	SystemData *SystemData_StatusARM `json:"systemData,omitempty"`
}

type SignalRCorsSettings_StatusARM struct {
	// AllowedOrigins: Gets or sets the list of origins that should be allowed to make cross-origin calls (for example:
	// http://example.com:12345). Use "*" to allow all. If omitted, allow all by default.
	AllowedOrigins []string `json:"allowedOrigins,omitempty"`
}

type SignalRFeature_StatusARM struct {
	Flag *FeatureFlags_Status `json:"flag,omitempty"`

	// Properties: Optional properties related to this feature.
	Properties map[string]string `json:"properties,omitempty"`

	// Value: Value of the feature flag. See Azure SignalR service document https://docs.microsoft.com/azure/azure-signalr/ for
	// allowed values.
	Value *string `json:"value,omitempty"`
}

type SignalRNetworkACLs_StatusARM struct {
	DefaultAction *ACLAction_Status `json:"defaultAction,omitempty"`

	// PrivateEndpoints: ACLs for requests from private endpoints
	PrivateEndpoints []PrivateEndpointACL_StatusARM `json:"privateEndpoints,omitempty"`
	PublicNetwork    *NetworkACL_StatusARM          `json:"publicNetwork,omitempty"`
}

type SignalRSkuTier_Status string

const (
	SignalRSkuTier_StatusBasic    = SignalRSkuTier_Status("Basic")
	SignalRSkuTier_StatusFree     = SignalRSkuTier_Status("Free")
	SignalRSkuTier_StatusPremium  = SignalRSkuTier_Status("Premium")
	SignalRSkuTier_StatusStandard = SignalRSkuTier_Status("Standard")
)

type SignalRTlsSettings_StatusARM struct {
	// ClientCertEnabled: Request client certificate during TLS handshake if enabled
	ClientCertEnabled *bool `json:"clientCertEnabled,omitempty"`
}

type SystemDataStatusCreatedByType string

const (
	SystemDataStatusCreatedByTypeApplication     = SystemDataStatusCreatedByType("Application")
	SystemDataStatusCreatedByTypeKey             = SystemDataStatusCreatedByType("Key")
	SystemDataStatusCreatedByTypeManagedIdentity = SystemDataStatusCreatedByType("ManagedIdentity")
	SystemDataStatusCreatedByTypeUser            = SystemDataStatusCreatedByType("User")
)

type SystemDataStatusLastModifiedByType string

const (
	SystemDataStatusLastModifiedByTypeApplication     = SystemDataStatusLastModifiedByType("Application")
	SystemDataStatusLastModifiedByTypeKey             = SystemDataStatusLastModifiedByType("Key")
	SystemDataStatusLastModifiedByTypeManagedIdentity = SystemDataStatusLastModifiedByType("ManagedIdentity")
	SystemDataStatusLastModifiedByTypeUser            = SystemDataStatusLastModifiedByType("User")
)

type UserAssignedIdentityProperty_StatusARM struct {
	// ClientId: Get the client id for the user assigned identity
	ClientId *string `json:"clientId,omitempty"`

	// PrincipalId: Get the principal id for the user assigned identity
	PrincipalId *string `json:"principalId,omitempty"`
}

type NetworkACL_StatusARM struct {
	// Allow: Allowed request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.
	Allow []SignalRRequestType_Status `json:"allow,omitempty"`

	// Deny: Denied request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.
	Deny []SignalRRequestType_Status `json:"deny,omitempty"`
}

type PrivateEndpointACL_StatusARM struct {
	// Allow: Allowed request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.
	Allow []SignalRRequestType_Status `json:"allow,omitempty"`

	// Deny: Denied request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.
	Deny []SignalRRequestType_Status `json:"deny,omitempty"`

	// Name: Name of the private endpoint connection
	Name *string `json:"name,omitempty"`
}

type ResourceLogCategory_StatusARM struct {
	// Enabled: Indicates whether or the resource log category is enabled.
	// Available values: true, false.
	// Case insensitive.
	Enabled *string `json:"enabled,omitempty"`

	// Name: Gets or sets the resource log category's name.
	// Available values: ConnectivityLogs, MessagingLogs.
	// Case insensitive.
	Name *string `json:"name,omitempty"`
}

type UpstreamTemplate_StatusARM struct {
	Auth *UpstreamAuthSettings_StatusARM `json:"auth,omitempty"`

	// CategoryPattern: Gets or sets the matching pattern for category names. If not set, it matches any category.
	// There are 3 kind of patterns supported:
	// 1. "*", it to matches any category name
	// 2. Combine multiple categories with ",", for example "connections,messages", it matches category "connections" and
	// "messages"
	// 3. The single category name, for example, "connections", it matches the category "connections"
	CategoryPattern *string `json:"categoryPattern,omitempty"`

	// EventPattern: Gets or sets the matching pattern for event names. If not set, it matches any event.
	// There are 3 kind of patterns supported:
	// 1. "*", it to matches any event name
	// 2. Combine multiple events with ",", for example "connect,disconnect", it matches event "connect" and "disconnect"
	// 3. The single event name, for example, "connect", it matches "connect"
	EventPattern *string `json:"eventPattern,omitempty"`

	// HubPattern: Gets or sets the matching pattern for hub names. If not set, it matches any hub.
	// There are 3 kind of patterns supported:
	// 1. "*", it to matches any hub name
	// 2. Combine multiple hubs with ",", for example "hub1,hub2", it matches "hub1" and "hub2"
	// 3. The single hub name, for example, "hub1", it matches "hub1"
	HubPattern *string `json:"hubPattern,omitempty"`

	// UrlTemplate: Gets or sets the Upstream URL template. You can use 3 predefined parameters {hub}, {category} {event}
	// inside the template, the value of the Upstream URL is dynamically calculated when the client request comes in.
	// For example, if the urlTemplate is `http://example.com/{hub}/api/{event}`, with a client request from hub `chat`
	// connects, it will first POST to this URL: `http://example.com/chat/api/connect`.
	UrlTemplate *string `json:"urlTemplate,omitempty"`
}

type UpstreamAuthSettings_StatusARM struct {
	ManagedIdentity *ManagedIdentitySettings_StatusARM `json:"managedIdentity,omitempty"`
	Type            *UpstreamAuthType_Status           `json:"type,omitempty"`
}

type ManagedIdentitySettings_StatusARM struct {
	// Resource: The Resource indicating the App ID URI of the target resource.
	// It also appears in the aud (audience) claim of the issued token.
	Resource *string `json:"resource,omitempty"`
}