// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package api

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	Pat_basicScopes  = "pat_basic.Scopes"
	Pat_bearerScopes = "pat_bearer.Scopes"
)

// Defines values for ClientType.
const (
	ClientTypeCustom     ClientType = "custom"
	ClientTypeGo         ClientType = "go"
	ClientTypeJavascript ClientType = "javascript"
	ClientTypePython     ClientType = "python"
	ClientTypeRuby       ClientType = "ruby"
)

// Defines values for CredentialStatus.
const (
	Expired CredentialStatus = "expired"
	Online  CredentialStatus = "online"
	Revoked CredentialStatus = "revoked"
)

// Defines values for ServiceServerType.
const (
	ServiceServerTypeCaddy      ServiceServerType = "caddy"
	ServiceServerTypeCustom     ServiceServerType = "custom"
	ServiceServerTypeDiagnostic ServiceServerType = "diagnostic"
	ServiceServerTypeGo         ServiceServerType = "go"
	ServiceServerTypeJavascript ServiceServerType = "javascript"
	ServiceServerTypePython     ServiceServerType = "python"
	ServiceServerTypeRuby       ServiceServerType = "ruby"
)

// Attachment defines model for attachment.
type Attachment struct {
	// Domains A list of domains for this attachment.
	Domains []string `json:"domains"`

	// Relationships Values used as parameters when referencing related resources.
	Relationships struct {
		Chain        RelationshipsChainApid         `json:"chain"`
		Organization RelationshipsOrganizationApid  `json:"organization"`
		Realm        RelationshipsRealmApid         `json:"realm"`
		Service      RelationshipsServiceApid       `json:"service"`
		SubCa        RelationshipsSubCaApidNullable `json:"sub_ca"`
	} `json:"relationships"`
}

// Attachments defines model for attachments.
type Attachments struct {
	Items []Attachment `json:"items"`
}

// AuthCliCodesResponse defines model for auth_cli_codes_response.
type AuthCliCodesResponse struct {
	// DeviceCode Unique code associated with origin device for CLI auth flow.
	DeviceCode string `json:"device_code"`

	// ExpiresIn Number of seconds until device and user codes expire, defaults to 900s (15m).
	ExpiresIn int32 `json:"expires_in"`

	// Interval Minimum number of seconds to wait before re-polling during CLI auth flow, defaults to 5s.
	Interval int32 `json:"interval"`

	// UserCode User verification code to be presented for the user to enter into a browser, formed by 8 characters with a hyphen in the middle.
	UserCode string `json:"user_code"`

	// VerificationUri The URL where the user will need to enter their code to complete the CLI auth flow.
	VerificationUri string `json:"verification_uri"`
}

// AuthCliPatTokensResponse defines model for auth_cli_pat_tokens_response.
type AuthCliPatTokensResponse struct {
	// PatToken Authenticated token for user API interactions.
	PatToken string `json:"pat_token"`
}

// Chain defines model for chain.
type Chain struct {
	// Name The name of the chain.
	Name string `json:"name"`

	// Relationships Values used as parameters when referencing related resources.
	Relationships struct {
		Org   RelationshipsOrganizationSlug `json:"org"`
		Realm RelationshipsRealmSlug        `json:"realm"`
	} `json:"relationships"`

	// Slug A value used as a parameter when referencing this chain.
	Slug string `json:"slug"`
}

// Client defines model for client.
type Client struct {
	// Name A name for the client.
	Name string `json:"name"`

	// Relationships Values used as parameters when referencing related resources.
	Relationships struct {
		Organization *RelationshipsOrganizationSlug `json:"organization,omitempty"`
		Service      *RelationshipsServiceSlug      `json:"service,omitempty"`
	} `json:"relationships"`

	// Slug A value used as a parameter when referencing this client.
	Slug string `json:"slug"`

	// Type A type for the client.
	Type ClientType `json:"type"`
}

// ClientType A type for the client.
type ClientType string

// Credential defines model for credential.
type Credential struct {
	// CreatedAt UTC time when credential was created.
	CreatedAt time.Time `json:"created_at"`

	// Name name of credential
	Name string `json:"name"`

	// Revision Current revision of credential
	Revision int64 `json:"revision"`

	// RevokedAt UTC time after which credential will be revoked
	RevokedAt *time.Time `json:"revoked_at"`

	// Serial serial id for credential
	Serial string `json:"serial"`

	// SignatureAlgorithm Algorithm used to sign credential
	SignatureAlgorithm interface{} `json:"signature_algorithm"`

	// Status current status of credential
	Status CredentialStatus `json:"status"`

	// TextualEncoding base64 textual encoding of credential based on RFC7468
	TextualEncoding string `json:"textual_encoding"`

	// Uuid uuid for credential
	Uuid *openapi_types.UUID `json:"uuid,omitempty"`

	// ValidAfter UTC time after which credential will be valid.
	ValidAfter time.Time `json:"valid_after"`

	// ValidBefore UTC time after which credential will no longer be valid.
	ValidBefore time.Time `json:"valid_before"`
}

// CredentialStatus current status of credential
type CredentialStatus string

// Credentials defines model for credentials.
type Credentials struct {
	Items []Credential `json:"items"`
}

// Eab defines model for eab.
type Eab struct {
	// HmacKey EAB HMAC key
	HmacKey string `json:"hmac_key"`

	// Kid EAB key identifier
	Kid string `json:"kid"`

	// Relationships Values used as parameters when referencing related resources.
	Relationships struct {
		Chain        RelationshipsChainSlug        `json:"chain"`
		Organization RelationshipsOrganizationSlug `json:"organization"`
		Realm        RelationshipsRealmSlug        `json:"realm"`
		Service      RelationshipsServiceSlug      `json:"service"`
		SubCa        RelationshipsSubCaSlug        `json:"sub_ca"`
	} `json:"relationships"`
}

// Error defines model for error.
type Error struct {
	// Detail A human-readable description of this occurrence of the problem.
	Detail string `json:"detail"`

	// Status HTTP status code of this occurrence of the problem.
	Status int32 `json:"status"`

	// Title A human-readable description of this problem type.
	Title string `json:"title"`

	// Type URI identifying problem.
	Type string `json:"type"`
}

// Organization defines model for organization.
type Organization struct {
	// Apid A value used as a parameter when referencing this organization.
	Apid string `json:"apid"`

	// Name The name for this organization.
	Name string `json:"name"`

	// Slug The slugified name for this organization.
	Slug string `json:"slug"`
}

// Organizations defines model for organizations.
type Organizations struct {
	Items []Organization `json:"items"`
}

// Realm defines model for realm.
type Realm struct {
	// Apid A value used as a parameter when referencing this realm.
	Apid string `json:"apid"`

	// Name The name for this realm.
	Name string `json:"name"`

	// Relationships Values used as parameters when referencing related resources.
	Relationships struct {
		Organization RelationshipsOrganizationApid `json:"organization"`
	} `json:"relationships"`

	// Slug The slugified name for this realm.
	Slug string `json:"slug"`
}

// Realms defines model for realms.
type Realms struct {
	Items []Realm `json:"items"`
}

// RelationshipsChainApid defines model for relationships_chain_apid.
type RelationshipsChainApid struct {
	// Apid A value used as a parameter when referencing this chain.
	Apid string `json:"apid"`
}

// RelationshipsChainSlug defines model for relationships_chain_slug.
type RelationshipsChainSlug struct {
	// Slug A value used as a parameter when referencing this chain.
	Slug string `json:"slug"`
}

// RelationshipsOrganizationApid defines model for relationships_organization_apid.
type RelationshipsOrganizationApid struct {
	// Apid A value used as a parameter when referencing this organization.
	Apid string `json:"apid"`
}

// RelationshipsOrganizationSlug defines model for relationships_organization_slug.
type RelationshipsOrganizationSlug struct {
	// Slug A value used as a parameter when referencing this organization.
	Slug string `json:"slug"`
}

// RelationshipsRealmApid defines model for relationships_realm_apid.
type RelationshipsRealmApid struct {
	// Apid A value used as a parameter when referencing this realm.
	Apid string `json:"apid"`
}

// RelationshipsRealmSlug defines model for relationships_realm_slug.
type RelationshipsRealmSlug struct {
	// Slug A value used as a parameter when referencing this realm.
	Slug string `json:"slug"`
}

// RelationshipsServiceApid defines model for relationships_service_apid.
type RelationshipsServiceApid struct {
	// Apid A value used as a parameter when referencing this service.
	Apid string `json:"apid"`
}

// RelationshipsServiceSlug defines model for relationships_service_slug.
type RelationshipsServiceSlug struct {
	// Slug A value used as a parameter when referencing this service.
	Slug string `json:"slug"`
}

// RelationshipsSubCaApidNullable defines model for relationships_sub_ca_apid_nullable.
type RelationshipsSubCaApidNullable struct {
	// Apid A value used as a parameter when referencing this sub_ca.
	Apid *string `json:"apid"`
}

// RelationshipsSubCaSlug defines model for relationships_sub_ca_slug.
type RelationshipsSubCaSlug struct {
	// Slug A value used as a parameter when referencing this sub_ca.
	Slug string `json:"slug"`
}

// Root defines model for root.
type Root struct {
	MinimumCliVersion string `json:"minimum_cli_version"`
	PersonalOrg       struct {
		Slug string `json:"slug"`
	} `json:"personal_org"`
	Whoami string `json:"whoami"`
}

// Service defines model for service.
type Service struct {
	// LocalhostPort A port number for use on localhost or null.
	LocalhostPort *int `json:"localhost_port"`

	// Name A name for the service.
	Name string `json:"name"`

	// Relationships Values used as parameters when referencing related resources.
	Relationships struct {
		Organization RelationshipsOrganizationSlug `json:"organization"`
	} `json:"relationships"`

	// ServerType A server type for the service.
	ServerType ServiceServerType `json:"server_type"`

	// Slug A value used as a parameter when referencing this service.
	Slug string `json:"slug"`
}

// ServiceServerType A server type for the service.
type ServiceServerType string

// Services defines model for services.
type Services struct {
	Items []Service `json:"items"`
}

// PathOrgParam defines model for path_org_param.
type PathOrgParam = string

// PathRealmParam defines model for path_realm_param.
type PathRealmParam = string

// PathServiceParam defines model for path_service_param.
type PathServiceParam = string

// QueryCaParam defines model for query_ca_param.
type QueryCaParam = string

// ServicesXtach200 defines model for services_xtach_200.
type ServicesXtach200 struct {
	// Domains A list of domains for this attachment.
	Domains []string `json:"domains"`

	// Relationships Values used as parameters when referencing related resources.
	Relationships struct {
		Chain        RelationshipsChainSlug        `json:"chain"`
		Organization RelationshipsOrganizationSlug `json:"organization"`
		Realm        *RelationshipsRealmSlug       `json:"realm,omitempty"`
		Service      RelationshipsServiceSlug      `json:"service"`
		SubCa        RelationshipsSubCaSlug        `json:"sub_ca"`
	} `json:"relationships"`
}

// ServicesXtach defines model for services_xtach.
type ServicesXtach struct {
	// Domains A list of domains for this attachment.
	Domains []string `json:"domains"`

	// Relationships Values used as parameters when referencing related resources.
	Relationships struct {
		Chain RelationshipsChainSlug `json:"chain"`
		Realm RelationshipsRealmSlug `json:"realm"`
	} `json:"relationships"`
}

// CreateEabTokenJSONBody defines parameters for CreateEabToken.
type CreateEabTokenJSONBody struct {
	Relationships struct {
		Chain        RelationshipsChainSlug        `json:"chain"`
		Organization RelationshipsOrganizationSlug `json:"organization"`
		Realm        RelationshipsRealmSlug        `json:"realm"`
		Service      *RelationshipsServiceSlug     `json:"service,omitempty"`
		SubCa        RelationshipsSubCaSlug        `json:"sub_ca"`
	} `json:"relationships"`
}

// CreateCliTokenJSONBody defines parameters for CreateCliToken.
type CreateCliTokenJSONBody struct {
	// DeviceCode Unique code associated with origin device for CLI auth flow.
	DeviceCode string `json:"device_code"`
}

// CreateClientJSONBody defines parameters for CreateClient.
type CreateClientJSONBody struct {
	Relationships *struct {
		Organization *RelationshipsOrganizationSlug `json:"organization,omitempty"`
		Service      *RelationshipsServiceSlug      `json:"service,omitempty"`
	} `json:"relationships,omitempty"`
	ServerType *string `json:"server_type,omitempty"`
	Type       *string `json:"type,omitempty"`
}

// GetCredentialsParams defines parameters for GetCredentials.
type GetCredentialsParams struct {
	// CaParam ca for operation
	CaParam *QueryCaParam `form:"ca_param,omitempty" json:"ca_param,omitempty"`
}

// AttachOrgServiceJSONBody defines parameters for AttachOrgService.
type AttachOrgServiceJSONBody struct {
	// Domains A list of domains for this attachment.
	Domains []string `json:"domains"`

	// Relationships Values used as parameters when referencing related resources.
	Relationships struct {
		Chain RelationshipsChainSlug `json:"chain"`
		Realm RelationshipsRealmSlug `json:"realm"`
	} `json:"relationships"`
}

// DetachOrgServiceJSONBody defines parameters for DetachOrgService.
type DetachOrgServiceJSONBody struct {
	// Domains A list of domains for this attachment.
	Domains []string `json:"domains"`

	// Relationships Values used as parameters when referencing related resources.
	Relationships struct {
		Chain RelationshipsChainSlug `json:"chain"`
		Realm RelationshipsRealmSlug `json:"realm"`
	} `json:"relationships"`
}

// CreateServiceJSONBody defines parameters for CreateService.
type CreateServiceJSONBody struct {
	LocalhostPort *int   `json:"localhost_port,omitempty"`
	Name          string `json:"name"`
	Relationships struct {
		Organization RelationshipsOrganizationSlug `json:"organization"`
	} `json:"relationships"`
	ServerType string `json:"server_type"`
}

// CreateEabTokenJSONRequestBody defines body for CreateEabToken for application/json ContentType.
type CreateEabTokenJSONRequestBody CreateEabTokenJSONBody

// CreateCliTokenJSONRequestBody defines body for CreateCliToken for application/json ContentType.
type CreateCliTokenJSONRequestBody CreateCliTokenJSONBody

// CreateClientJSONRequestBody defines body for CreateClient for application/json ContentType.
type CreateClientJSONRequestBody CreateClientJSONBody

// AttachOrgServiceJSONRequestBody defines body for AttachOrgService for application/json ContentType.
type AttachOrgServiceJSONRequestBody AttachOrgServiceJSONBody

// DetachOrgServiceJSONRequestBody defines body for DetachOrgService for application/json ContentType.
type DetachOrgServiceJSONRequestBody DetachOrgServiceJSONBody

// CreateServiceJSONRequestBody defines body for CreateService for application/json ContentType.
type CreateServiceJSONRequestBody CreateServiceJSONBody
